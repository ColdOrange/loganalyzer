package loganalyzer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	log "loganalyzer/loganalyzer/logging"
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
		log.Fatalln("Config project root path error:", err)
	}
	log.Debugln("Project root path:", ProjectPath)

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
		dbConfig.Initialized = false
		return
	}
	if err != nil {
		log.Fatalln("Load DBConfig error:", err)
	}

	err = json.Unmarshal(data, &dbConfig)
	if err != nil {
		log.Fatalln("Load DBConfig error:", err)
	}
	dbConfig.Initialized = true

	// Open DB
	log.Debugln("Open DB")
	db, err = sql.Open(dbConfig.Driver, fmt.Sprintf("%s:%s@/%s", dbConfig.Username, dbConfig.Password, dbConfig.Database))
	if err != nil {
		log.Fatalln("Open DB error:", err)
	}
}

func loadLogConfig() {
	data, err := ioutil.ReadFile(logConfigFile)
	if os.IsNotExist(err) {
		log.Debugln("Log config file not found, uninitialized")
		logConfig.Initialized = false
		return
	}
	if err != nil {
		log.Errorln("Load LogConfig error:", err)
		return
	}

	err = json.Unmarshal(data, &logConfig)
	if err != nil {
		log.Errorln("Load LogConfig error:", err)
		return
	}
	logConfig.Initialized = true
}

func getDBConfig() []byte {
	data, err := json.Marshal(dbConfig)
	if err != nil {
		log.Errorln("Marshal DBConfig error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}

func setDBConfig(data []byte) []byte { // TODO: maybe this function is too long
	var config DBConfig
	// Unmarshal
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Errorln("Unmarshal DBConfig error:", err)
		return []byte(`{"status": "failed"}`) // TODO: provide detailed error message
	}
	if config.Driver == dbConfig.Driver &&
		config.Username == dbConfig.Username &&
		config.Password == dbConfig.Password &&
		config.Database == dbConfig.Database { // TODO: check config values
		return []byte(`{"status": "successful"}`)
	}

	// Write to config file
	file, err := os.Create(dbConfigFile)
	if err != nil {
		log.Errorln("Create DBConfig file error:", err)
		return []byte(`{"status": "failed"}`)
	}
	_, err = file.Write(data)
	if err != nil {
		log.Errorln("Write DBConfig file error:", err)
		file.Close()
		return []byte(`{"status": "failed"}`)
	}
	err = file.Close()
	if err != nil {
		log.Errorln("Close DBConfig file error:", err)
		return []byte(`{"status": "failed"}`)
	}

	// Uninitialized, just open DB
	if dbConfig.Initialized == false {
		log.Debugln("Open DB")
		db, err = sql.Open(config.Driver, fmt.Sprintf("%s:%s@/%s", config.Username, config.Password, config.Database))
		if err != nil {
			log.Fatalln("Open DB error:", err)
		}
		dbConfig = config
		dbConfig.Initialized = true
		return []byte(`{"status": "successful"}`)
	}
	// Initialized, reopen and clear DB
	log.Debugln("Reopen and clear DB")
	err = db.Close()
	if err != nil {
		log.Errorln("Close old DB error:", err)
		//return []byte(`{"status": "failed"}`) // TODO: is it OK not to return error?
	}
	db, err = sql.Open(config.Driver, fmt.Sprintf("%s:%s@/", config.Username, config.Password))
	if err != nil {
		log.Errorln("Open DB error:", err)
		return []byte(`{"status": "failed"}`)
	}
	// Drop old DB
	_, err = db.Exec("DROP DATABASE IF EXISTS " + dbConfig.Database)
	if err != nil {
		log.Errorln("Drop old DB error:", err)
		return []byte(`{"status": "failed"}`)
	}
	// Create new DB (may throw error if exists)
	_, err = db.Exec("CREATE DATABASE " + config.Database)
	if err != nil {
		log.Errorln("Create new DB error:", err)
		return []byte(`{"status": "failed"}`)
	}
	_, err = db.Exec("USE " + config.Database)
	if err != nil {
		log.Fatalln("USE new DB error:", err)
		return []byte(`{"status": "failed"}`)
	}
	dbConfig = config
	dbConfig.Initialized = true
	return []byte(`{"status": "successful"}`)
}
