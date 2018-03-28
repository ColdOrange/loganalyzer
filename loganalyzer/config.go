package loganalyzer

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"runtime"

	log "loganalyzer/loganalyzer/logging"
)

type Config struct {
	Format
}

type Format struct {
	LogPattern string
	LogFormat  []string
	TimeFormat string
}

func loadConfig() Config {
	_, filename, _, ok := runtime.Caller(1) // TODO: maybe a less tricky way?
	if !ok {
		log.Fatalf("Runtime caller error")
	}
	data, err := ioutil.ReadFile(path.Join(path.Dir(filename), "../conf/conf.yml"))
	if err != nil {
		log.Fatalf("Read config error: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Unmarshal config error: %v", err)
	}
	return config
}
