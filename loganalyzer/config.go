package loganalyzer

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	log "github.com/ColdOrange/loganalyzer/loganalyzer/logging"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Initialized bool   `json:"initialized"`
	Driver      string `json:"driver"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Database    string `json:"database"`
}

type LogConfig struct {
	Initialized bool     `json:"initialized"`
	LogFile     string   `json:"logFile"`
	LogPattern  string   `json:"logPattern"`
	LogFormat   []string `json:"logFormat"`
	TimeFormat  string   `json:"timeFormat"`
}

var (
	ProjectPath   string
	db            *sql.DB
	dbConfig      DBConfig
	dbConfigFile  string
	logConfig     LogConfig
	logConfigFile string
)

func init() {
	// Resolve project root path, it's a little bit tricky but I can't figure out a better way...
	_, filename, _, ok := runtime.Caller(0) // TODO: maybe a less tricky way?
	if !ok {
		log.Fatalln("Runtime caller error")
	}
	ProjectPath, err := filepath.Abs(path.Join(path.Dir(filename), "/.."))
	if err != nil {
		log.Fatalln("Resolve project root path error:", err)
	}
	log.Debugln("Project root path:", ProjectPath)

	// Create config directory if not exists
	configDir := path.Join(ProjectPath, "config")
	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		if os.Mkdir(configDir, 0755) != nil {
			log.Fatalln("Create config directory error")
		}
		log.Debugln("Create config directory")
	}

	// Load Config
	dbConfigFile = path.Join(ProjectPath, "config/db_config.json")
	loadDBConfig()
	logConfigFile = path.Join(ProjectPath, "config/log_config.json")
	loadLogConfig()
}

func loadDBConfig() {
	data, err := ioutil.ReadFile(dbConfigFile)
	if os.IsNotExist(err) {
		log.Debugln("DB config file not found, uninitialized")
		return
	}
	if err != nil {
		log.Errorf("Read DBConfig file error: %v, uninitialized", err)
		return
	}
	err = json.Unmarshal(data, &dbConfig)
	if err != nil {
		log.Errorf("Unmarshal DBConfig file error: %v, uninitialized", err)
		return
	}

	// Open DB
	log.Debugln("DBConfig loaded, open DB")
	err = openDB(&dbConfig)
	if err != nil {
		log.Errorln(err)
		return
	}
	log.Debugln("Open DB successfully")
	dbConfig.Initialized = true
}

func openDB(config *DBConfig) error {
	var err error // Reuse global db variable
	db, err = sql.Open(config.Driver, fmt.Sprintf("%s:%s@/%s?multiStatements=true", config.Username, config.Password, config.Database))
	if err != nil {
		return errors.New("Open DB error: " + err.Error())
	}
	// Note: sql.Open doesn't really open a connection, so we need ping to check error.
	// See https://github.com/go-sql-driver/mysql/wiki/Examples
	err = db.Ping()
	if err != nil {
		return errors.New("Connect to DB error: " + err.Error())
	}
	return nil
}

func loadLogConfig() {
	data, err := ioutil.ReadFile(logConfigFile)
	if os.IsNotExist(err) {
		log.Debugln("Log config file not found, use default config")
		useDefaultLogConfig()
		return
	}
	if err != nil {
		log.Errorf("Read LogConfig file error: %v, use default config", err)
		useDefaultLogConfig()
		return
	}
	err = json.Unmarshal(data, &logConfig)
	if err != nil {
		log.Errorf("Unmarshal LogConfig file error: %v, use default config", err)
		useDefaultLogConfig()
	}
}

func useDefaultLogConfig() {
	sampleLogFile, err := filepath.Abs(path.Join(ProjectPath, "sample/sample.log"))
	if err != nil {
		log.Errorf("Resolve sample log file path error: %v, LogConfig uninitialized", err)
		return
	}
	logConfig.Initialized = true
	logConfig.LogFile = sampleLogFile
	logConfig.LogPattern = "(.*) - - \\[(.*)\\] \"(.*) (.*) (.*)\" (.*) (.*) \"(.*)\" \"(.*)\" (.*)"
	logConfig.LogFormat = []string{"IP", "Time", "RequestMethod", "RequestURL", "HTTPVersion", "ResponseCode", "ContentSize", "Referrer", "UserAgent", "ResponseTime"}
	logConfig.TimeFormat = "02/Jan/2006:15:04:05 -0700"
}

func getDBConfig() []byte {
	data, err := json.Marshal(dbConfig)
	if err != nil {
		log.Errorln("Marshal DBConfig error:", err)
		return jsonError("Marshal DBConfig error", err)
	}
	return data
}

func setDBConfig(data []byte) []byte {
	var config DBConfig
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Errorln("Unmarshal DBConfig error:", err)
		return jsonError("Unmarshal DBConfig error", err)
	}

	//if dbConfigUnchanged(&config) { // TODO: not stable
	//	log.Debugln("DBConfig unchanged")
	//	return jsonSuccess() // TODO: maybe put config unchanged in success message?
	//}

	// Write to config file
	file, err := os.Create(dbConfigFile)
	if err != nil {
		log.Errorln("Create DBConfig file error:", err)
		return jsonError("Create DBConfig file error", err)
	}
	defer file.Close() // TODO: should we check close error?
	_, err = file.Write(data)
	if err != nil {
		log.Errorln("Write DBConfig file error:", err)
		os.Remove(dbConfigFile) // TODO: should we check remove error?
		return jsonError("Write DBConfig file error", err)
	}

	// Create and (re)open DB with new config
	if dbConfig.Initialized == false {
		log.Debugln("DBConfig was uninitialized, create and open DB")
	} else {
		log.Debugln("DBConfig was initialized, clear and reopen DB")
		_, err = db.Exec("DROP DATABASE IF EXISTS " + dbConfig.Database) // Drop old DB
		if err != nil {
			log.Errorln("Drop DB error:", err)
			return jsonError("Drop DB error", err)
		}
	}
	err = createNewDB(&config)
	if err != nil {
		log.Errorln(err)
		return jsonError(err)
	}
	// Note: db (returned by sql.Open) is not a connection, but a connection pool.
	//       db.exec("USE ...") is not helpful here, because it only affect one connection.
	//       To change all connections, we have to reopen db.
	// See: http://go-database-sql.org/surprises.html - Connection State Mismatch
	if db != nil {
		err = db.Close()
		if err != nil {
			log.Errorln("Close DB error:", err)
			return jsonError("Close DB error", err)
		}
	}
	err = openDB(&config)
	if err != nil {
		log.Errorln(err)
		return jsonError(err)
	}

	// Create reports table
	_, err = db.Exec(createReportsTable)
	if err != nil {
		log.Errorln("Create reports table error:", err)
		return jsonError("Create reports table error", err)
	}

	dbConfig = config
	dbConfig.Initialized = true
	return jsonSuccess()
}

func dbConfigUnchanged(config *DBConfig) bool {
	return config.Driver == dbConfig.Driver &&
		config.Username == dbConfig.Username &&
		config.Password == dbConfig.Password &&
		config.Database == dbConfig.Database
}

func createNewDB(config *DBConfig) error {
	var db *sql.DB // Open a local db to do the creating job, the global db will reopen anyway
	db, err := sql.Open(config.Driver, fmt.Sprintf("%s:%s@/?multiStatements=true", config.Username, config.Password))
	if err != nil {
		return errors.New("Open DB error: " + err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return errors.New("Connect to DB error: " + err.Error())
	}
	_, err = db.Exec("CREATE DATABASE " + config.Database) // Throw on purpose if exists (in case we drop some valuable data)
	if err != nil {
		return errors.New("Create to DB error: " + err.Error())
	}
	return nil
}

func getLogConfig() []byte {
	data, err := json.Marshal(logConfig)
	if err != nil {
		log.Errorln("Marshal LogConfig error:", err)
		return jsonError("Marshal LogConfig error", err)
	}
	return data
}

func setLogConfig(data []byte) []byte {
	var config LogConfig
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Errorln("Unmarshal LogConfig error:", err)
		return jsonError("Unmarshal LogConfig error", err)
	}

	//if logConfigUnchanged(&config) {
	//	log.Debugln("LogConfig unchanged")
	//	return jsonSuccess()
	//}

	// Write to config file
	file, err := os.Create(logConfigFile)
	if err != nil {
		log.Errorln("Create LogConfig file error:", err)
		return jsonError("Create LogConfig file error", err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		log.Errorln("Write LogConfig file error:", err)
		os.Remove(logConfigFile)
		return jsonError("Write LogConfig file error", err)
	}

	// Insert into reports table
	if db == nil {
		log.Errorln("Database uninitialized")
		return jsonError("Database uninitialized")
	}
	res, err := db.Exec("INSERT INTO reports (file) VALUES (?)", config.LogFile)
	if err != nil {
		log.Errorln("Insert into reports table error:", err)
		return jsonError("Insert into reports table error,", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Errorln("Get last insert id error:", err)
		return jsonError("Get last insert id error", err)
	}

	// Start analyze
	logConfig = config
	logConfig.Initialized = true
	err = Analyze(int(id))
	if err != nil {
		return jsonError("Analyze log file error", err)
	}
	return jsonSuccess()
}

func logConfigUnchanged(config *LogConfig) bool {
	return config.LogFile == logConfig.LogFile &&
		config.LogPattern == logConfig.LogPattern &&
		stringSliceEqual(config.LogFormat, logConfig.LogFormat) &&
		config.TimeFormat == logConfig.TimeFormat
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
