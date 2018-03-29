package loganalyzer

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"

	log "loganalyzer/loganalyzer/logging"
)

type Config struct {
	Format
	DataSource
}

type Format struct {
	LogFile    string
	LogPattern string
	LogFormat  []string
	TimeFormat string
}

type DataSource struct {
	Driver   string
	Username string
	Password string
	Database string
	Table    string
}

var ProjectPath string

func init() {
	// Get project root path, it's a little bit tricky but I can't figure out a better way...
	_, filename, _, ok := runtime.Caller(1) // TODO: maybe a less tricky way?
	if !ok {
		log.Fatalln("Runtime caller error")
	}
	ProjectPath, err := filepath.Abs(path.Join(path.Dir(filename), "/.."))
	if err != nil {
		log.Fatalln("Config project root path error:", err)
	}
	log.Debugln("Project root path:", ProjectPath)
}

func loadConfig() Config {
	data, err := ioutil.ReadFile(path.Join(ProjectPath, "config/config.yml"))
	if err != nil {
		log.Fatalln("Read config error:", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln("Unmarshal config error:", err)
	}
	return config
}
