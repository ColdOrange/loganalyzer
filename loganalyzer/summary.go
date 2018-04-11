package loganalyzer

import (
	"encoding/json"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	log "loganalyzer/loganalyzer/logging"
)

type Summary struct {
	FileName  string  `json:"file_name"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	PageViews int64   `json:"page_views"`
	UserViews int64   `json:"user_views"`
	Bandwidth float64 `json:"bandwidth"`
}

func summary() []byte {
	var summary Summary
	_, summary.FileName = filepath.Split(config.LogFile)

	row := db.QueryRow("SELECT time FROM log ORDER BY id LIMIT 1")
	err := row.Scan(&summary.StartTime)
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	row = db.QueryRow("SELECT time FROM log ORDER BY id DESC LIMIT 1")
	err = row.Scan(&summary.EndTime)
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	row = db.QueryRow("SELECT count(*), count(distinct(ip)), sum(content_size) FROM log")
	err = row.Scan(&summary.PageViews, &summary.UserViews, &summary.Bandwidth)
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(summary)
	if err != nil {
		log.Errorln("Summary json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
