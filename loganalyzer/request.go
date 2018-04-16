package loganalyzer

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	log "loganalyzer/loganalyzer/logging"
)

type RequestMethod struct {
	Method string `json:"requestMethod"`
	Count  int64  `json:"count"`
}

type RequestURL struct {
	URL   string `json:"requestURL"`
	Count int64  `json:"count"`
}

type HTTPVersion struct {
	Version string `json:"httpVersion"`
	Count   int64  `json:"count"`
}

func requestMethod() []byte {
	rows, err := db.Query("SELECT request_method, count(*) as count FROM log GROUP BY request_method ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	var (
		requestMethod []RequestMethod
		method        string
		count         int64
	)
	for rows.Next() {
		err := rows.Scan(&method, &count)
		if err != nil {
			log.Errorln("DB query error:", err)
			return []byte(`{"status": "failed"}`)
		}
		requestMethod = append(requestMethod, RequestMethod{Method: method, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(requestMethod)
	if err != nil {
		log.Errorln("RequestMethod json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}

func requestURL() []byte {
	rows, err := db.Query("SELECT url_path, count(*) as count FROM log WHERE url_is_static='0' GROUP BY url_path ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	var (
		requestURL []RequestURL
		url        string
		count      int64
	)
	for rows.Next() {
		err := rows.Scan(&url, &count)
		if err != nil {
			log.Errorln("DB query error:", err)
			return []byte(`{"status": "failed"}`)
		}
		requestURL = append(requestURL, RequestURL{URL: url, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(requestURL)
	if err != nil {
		log.Errorln("RequestURL json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}

func httpVersion() []byte {
	rows, err := db.Query("SELECT http_version, count(*) as count FROM log GROUP BY http_version ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	var (
		httpVersion []HTTPVersion
		version     string
		count       int64
	)
	for rows.Next() {
		err := rows.Scan(&version, &count)
		if err != nil {
			log.Errorln("DB query error:", err)
			return []byte(`{"status": "failed"}`)
		}
		httpVersion = append(httpVersion, HTTPVersion{Version: version, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(httpVersion)
	if err != nil {
		log.Errorln("HTTPVersion json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
