package loganalyzer

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"

	log "loganalyzer/loganalyzer/logging"
)

var (
	config  Config
	summary Summary
)

func Analyze(logFile string) {
	log.Infoln("Starting analyzing log file")
	start := time.Now()

	config = loadConfig()
	pattern, err := regexp.Compile(config.LogPattern)
	if err != nil {
		log.Fatalf("Log pattern compile error: %v", err)
	}

	_, summary.FileName = filepath.Split(logFile)
	file, err := os.Open(logFile)
	if err != nil {
		log.Fatalf("Log file open error: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for i := 0; ; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Errorf("Read log file error at line %d: %v", i, err)
				continue
			} else if len(line) == 0 {
				break // if len(line) != 0, we still have a last line to parse which is not ended with '\n'
			}
		}

		fields := pattern.FindStringSubmatch(line)
		if len(fields) != len(config.LogFormat)+1 {
			log.Warnf("Log file format wrong at line %d", i)
			continue
		}

		// Do analysis for each report item
		summary.analyze(i, fields)
	}

	log.Infof("Finished analyzing log file in %.3fs", time.Since(start).Seconds())
}
