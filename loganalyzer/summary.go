package loganalyzer

import (
	"encoding/json"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "loganalyzer/loganalyzer/logging"
)

type Summary struct {
	FileName  string    `json:"file_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	PageViews int64     `json:"page_views"`
	UserViews int64     `json:"user_views"`
	Bandwidth float64   `json:"bandwidth"`
}

func summary() []byte {
	var summary Summary
	_, summary.FileName = filepath.Split(config.LogFile)

	row := db.QueryRow("SELECT time FROM log ORDER BY id LIMIT 1")
	err := row.Scan(&summary.StartTime)
	if err != nil { // TODO: why so many err checks...
		log.Errorf("DB query error: %v", err)
		return []byte(`{"status": "failed"}`)
	}

	row = db.QueryRow("SELECT time FROM log ORDER BY id DESC LIMIT 1")
	err = row.Scan(&summary.EndTime)
	if err != nil {
		log.Errorf("DB query error: %v", err)
		return []byte(`{"status": "failed"}`)
	}

	row = db.QueryRow("SELECT count(*) FROM log")
	err = row.Scan(&summary.PageViews)
	if err != nil {
		log.Errorf("DB query error: %v", err)
		return []byte(`{"status": "failed"}`)
	}

	row = db.QueryRow("SELECT count(distinct(ip)) FROM log")
	err = row.Scan(&summary.UserViews)
	if err != nil {
		log.Errorf("DB query error: %v", err)
		return []byte(`{"status": "failed"}`)
	}

	row = db.QueryRow("SELECT sum(content_size) FROM log")
	err = row.Scan(&summary.Bandwidth)
	if err != nil {
		log.Errorf("DB query error: %v", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(summary)
	if err != nil {
		log.Errorf("Summary json marshal error: %v", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
