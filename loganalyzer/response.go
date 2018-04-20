package loganalyzer

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	log "loganalyzer/loganalyzer/logging"
)

type StatusCode struct {
	StatusCode string `json:"statusCode"`
	Count      int64  `json:"count"`
}

func statusCode() []byte {
	rows, err := db.Query("SELECT response_code, count(*) as count FROM log GROUP BY response_code ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	var (
		statusCode []StatusCode
		code       int
		count      int64
	)
	for rows.Next() {
		err := rows.Scan(&code, &count)
		if err != nil {
			log.Errorln("DB query error:", err)
			return []byte(`{"status": "failed"}`)
		}
		statusText := strconv.Itoa(code) + " " + http.StatusText(code)
		statusCode = append(statusCode, StatusCode{StatusCode: statusText, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(statusCode)
	if err != nil {
		log.Errorln("StatusCode json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
