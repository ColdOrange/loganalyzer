package loganalyzer

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
	log "loganalyzer/loganalyzer/logging"
)

var (
	config Config
	db     *sql.DB
)

func Analyze() {
	log.Infoln("Starting analyzing log file")
	zero := time.Now()

	// Load config
	config = loadConfig()
	pattern, err := regexp.Compile(config.LogPattern)
	if err != nil {
		log.Fatalln("Log pattern compile error:", err)
	}

	// Open DB and clear table
	log.Debugln("Open and clear DB")
	db, err = sql.Open(config.Driver, fmt.Sprintf("%s:%s@/%s?parseTime=false", config.Username, config.Password, config.Database))
	if err != nil {
		log.Fatalln("DB open error:", err)
	}
	//defer db.Close()
	_, err = db.Exec("TRUNCATE TABLE log")
	if err != nil {
		log.Fatalln("DB truncate table error:", err)
	}

	// Prepare `batch` insert stmt
	const batchSize = 100
	var batch = 0
	var batchValues []interface{}
	stmt, err := db.Prepare(prepareBatchInsertStmt(batchSize))
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
	for i := 1; ; i++ {
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
			case "IP":
				if len(fields[j]) > 46 { // Max IPv6 length
					log.Warnf("[IP] exceed max length (46, got %d) at line %d", len(fields[j]), i)
					continue ReadLog
				}
				values = append(values, fields[j])
			case "ResponseCode":
				if len(fields[j]) > 3 { // HTTP Status Code length is always 3 as known
					log.Warnf("[ResponseCode] exceed max length (5, got %d) at line %d", len(fields[j]), i)
					continue ReadLog
				}
				values = append(values, fields[j])
			case "RequestMethod", "HTTPVersion":
				if len(fields[j]) > 10 { // Normal length shouldn't be larger
					log.Debugf("[%s] exceed max length (10, got %d) at line %d", config.LogFormat[j-1], len(fields[j]), i)
					continue ReadLog
				}
				values = append(values, fields[j])
			case "UserAgent", "Referer":
				if len(fields[j]) > 255 { // We set this threshold in purpose, these values can be madly long in fact
					log.Debugf("[%s] exceed max length (255, got %d) at line %d", config.LogFormat[j-1], len(fields[j]), i)
					continue ReadLog
				}
				values = append(values, fields[j])
			case "RequestURL":
				u, err := url.Parse(fields[j])
				if err != nil {
					log.Warnf("[RequestURL] format wrong at line %d", i)
					continue ReadLog
				}
				if len(u.Path) > 255 { // We set this threshold in purpose, these values can be madly long in fact
					log.Debugf("[RequestURL.Path] exceed max length (255, got %d) at line %d", len(u.Path), i)
					continue ReadLog
				}
				if !utf8.ValidString(u.Path) {
					log.Debugf("[RequestURL.Path] is not a valid utf8 string at line %d", i)
					continue ReadLog
				}
				values = append(values, u.Path)
				if len(u.RawQuery) > 255 { // We set this threshold in purpose, these values can be madly long in fact
					log.Debugf("[RequestURL.Query] exceed max length (255, got %d) at line %d", len(u.RawQuery), i)
					continue ReadLog
				}
				values = append(values, u.RawQuery)
				values = append(values, isStatic(u.Path))
			case "Time":
				timestamp, err := time.Parse(config.TimeFormat, fields[j])
				if err != nil {
					log.Warnf("[Time] format wrong at line %d", i)
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
					log.Warnf("[ContentSize] format wrong at line %d", i)
					continue ReadLog
				}
				values = append(values, size)
			}
		}

		// Batch insert into DB
		batch++
		batchValues = append(batchValues, values...)
		if batch == batchSize {
			_, err = stmt.Exec(batchValues...)
			if err != nil {
				log.Errorf("DB batch insert error at line %d: %v", i, err)
			}
			batch = 0
			batchValues = nil
		}
	}

	if batch != 0 { // Last batch insert
		stmt, err := db.Prepare(prepareBatchInsertStmt(batch))
		if err != nil {
			log.Fatalln("DB insert stmt prepare error:", err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(batchValues...)
		if err != nil {
			log.Errorln("DB last batch insert error:", err)
		}
	}

	log.Debugln("Finished inserting into DB")
	log.Infof("Finished analyzing log file in %.3fs", time.Since(zero).Seconds())
}

func prepareBatchInsertStmt(batchSize int) string {
	var fields []string
	for _, field := range config.LogFormat {
		fields = append(fields, LogTableFields[field]...)
	}

	placeholder := make([]string, len(fields))
	for i := 0; i < len(fields); i++ {
		placeholder[i] = "?"
	}
	placeholderString := "(" + strings.Join(placeholder, ",") + ")"
	placeholders := make([]string, batchSize)
	for i := 0; i < batchSize; i++ {
		placeholders[i] = placeholderString
	}

	result := fmt.Sprintf("INSERT INTO log (%s) VALUES %s", strings.Join(fields, ","), strings.Join(placeholders, ","))
	//log.Debugln("Prepare Insert statement:", result)
	return result
}

var nonStaticExt = []string{
	"html", "htm", "shtml", "shtm", "xml", "php", "jsp", "asp", "aspx", "cgi", "perl", "do",
}

func isStatic(urlPath string) int {
	if !strings.Contains(urlPath, ".") {
		return 0
	}
	s := strings.Split(urlPath, ".")
	ext := strings.ToLower(strings.Split(s[len(s)-1], " ")[0])
	for _, e := range nonStaticExt {
		if strings.HasPrefix(ext, e) {
			return 0
		}
	}
	return 1
}
