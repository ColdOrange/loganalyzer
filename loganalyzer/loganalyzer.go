package loganalyzer

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "loganalyzer/loganalyzer/logging"
)

var (
	config     Config
	db         *sql.DB
	fieldFlags = make(map[string]bool)
)

func Analyze() {
	log.Infoln("Starting analyzing log file")
	zero := time.Now()

	// Load config
	config = loadConfig()
	for _, field := range config.LogFormat {
		fieldFlags[field] = true
	}
	pattern, err := regexp.Compile(config.LogPattern)
	if err != nil {
		log.Fatalln("Log pattern compile error:", err)
	}

	// Open DB and clear table
	log.Debugln("Open and clear DB")
	db, err = sql.Open(config.Driver, fmt.Sprintf("%s:%s@/%s?parseTime=true", config.Username, config.Password, config.Database))
	if err != nil {
		log.Fatalln("DB open error:", err)
	}
	//defer db.Close()
	_, err = db.Exec("TRUNCATE TABLE log")
	if err != nil {
		log.Fatalln("DB truncate table error:", err)
	}

	// Prepare insert stmt
	stmt, err := db.Prepare(prepareInsertStmt())
	if err != nil {
		log.Fatalln("DB insert stmt prepare error:", err)
	}
	defer stmt.Close()

	// Open log file
	file, err := os.Open(config.LogFile)
	if err != nil {
		log.Fatalln("Log file open error:", err)
	}
	defer file.Close()

	// Read and process log file by line
	reader := bufio.NewReader(file)
ReadLog:
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

		// Get log fields by regex
		fields := pattern.FindStringSubmatch(line)
		if len(fields) != len(config.LogFormat)+1 {
			log.Warnf("Log format wrong at line %d", i)
			continue
		}

		// Get field values
		var values []interface{}
		for j := 1; j <= len(config.LogFormat); j++ {
			switch config.LogFormat[j-1] {
			case "IP", "RequestMethod", "RequestURL", "HTTPVersion", "ResponseCode", "UserAgent", "Referer":
				if len(fields[j]) > 255 {
					log.Debugf("Log [%s] exceed max length at line %d, ignored", config.LogFormat[j-1], i)
					continue ReadLog
				}
				values = append(values, fields[j])
			case "Time":
				timestamp, err := time.Parse(config.TimeFormat, fields[j])
				if err != nil {
					log.Warnf("Log [Time] format wrong at line %d", i)
					continue ReadLog
				}
				values = append(values, timestamp)
			case "ResponseTime":
			case "ContentSize":
				if fields[j] == "-" {
					values = append(values, 0)
					continue
				}
				size, err := strconv.Atoi(fields[j])
				if err != nil {
					log.Warnf("Log [ContentSize] format wrong at line %d", i)
					continue ReadLog
				}
				values = append(values, size)
			}
		}

		// Insert into DB
		_, err = stmt.Exec(values...) // TODO: use Batch to improve performance
		if err != nil {
			log.Fatalln("DB insert stmt execute error:", err)
		}
	}

	log.Debugln("Finished inserting into DB")
	log.Infof("Finished analyzing log file in %.3fs", time.Since(zero).Seconds())
}

func prepareInsertStmt() string {
	var fields, placeholders []string
	for i, field := range LogFields {
		if fieldFlags[field] {
			fields = append(fields, DBFields[i])
			placeholders = append(placeholders, "?")
		}
	}

	result := fmt.Sprintf("INSERT INTO log (%s) VALUES (%s)", strings.Join(fields, ","), strings.Join(placeholders, ","))
	log.Debugln("Prepare Insert statement:", result)
	return result
}
