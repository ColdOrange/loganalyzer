package loganalyzer

import (
	"encoding/json"
	"os"
	"path/filepath"

	log "loganalyzer/loganalyzer/logging"
)

type Summary struct {
	FileName  string  `json:"fileName"`
	FileSize  int64   `json:"fileSize"`
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
	PageViews int64   `json:"pageViews"`
	UserViews int64   `json:"userViews"`
	Bandwidth float64 `json:"bandwidth"`
}

func summary(table string) []byte {
	var summary Summary
	_, summary.FileName = filepath.Split(logConfig.LogFile)

	file, err := os.Open(logConfig.LogFile)
	if err != nil {
		log.Errorln("Log file open error:", err)
		return []byte(`{"status": "failed"}`)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Errorln("Get file stat error:", err)
		return []byte(`{"status": "failed"}`)
	}
	summary.FileSize = stat.Size()

	row := db.QueryRow("SELECT time FROM " + table + " ORDER BY id LIMIT 1")
	err = row.Scan(&summary.StartTime)
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	row = db.QueryRow("SELECT time FROM " + table + " ORDER BY id DESC LIMIT 1")
	err = row.Scan(&summary.EndTime)
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	row = db.QueryRow("SELECT count(*), count(distinct(ip)), sum(content_size) FROM " + table)
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
