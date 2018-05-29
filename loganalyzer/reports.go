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
		log.Errorln("Database uninitialized")
		return jsonError("Database uninitialized")
	}

	rows, err := db.Query("SELECT id, file FROM reports ORDER BY id ASC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
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
			return jsonError("DB query error", err)
		}
		reports = append(reports, Report{Id: id, File: file})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(reports)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}
