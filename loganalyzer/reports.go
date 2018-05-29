package loganalyzer

import (
	"encoding/json"

	log "loganalyzer/loganalyzer/logging"
)

type Report struct {
	Id   int64  `json:"id"`
	File string `json:"file"`
}

func reports() []byte {
	if db == nil {
		log.Errorln("DB uninitialized")
		return jsonError("DB uninitialized")
	}

	rows, err := db.Query("SELECT id, file FROM reports ORDER BY id ASC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	var (
		reports []Report
		id      int64
		file    string
	)
	for rows.Next() {
		err := rows.Scan(&id, &file)
		if err != nil {
			log.Errorln("DB query error:", err)
			return []byte(`{"status": "failed"}`)
		}
		reports = append(reports, Report{Id: id, File: file})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(reports)
	if err != nil {
		log.Errorln("Reports json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
