package loganalyzer

import (
	"encoding/json"

	log "github.com/ColdOrange/loganalyzer/loganalyzer/logging"
)

type RequestMethod struct {
	Method string `json:"requestMethod"`
	Count  int64  `json:"count"`
}

func requestMethod(table string) []byte {
	rows, err := db.Query("SELECT request_method, count(*) as count FROM " + table + " GROUP BY request_method ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
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
			return jsonError("DB query error", err)
		}
		requestMethod = append(requestMethod, RequestMethod{Method: method, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(requestMethod)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}

type HTTPVersion struct {
	Version string `json:"httpVersion"`
	Count   int64  `json:"count"`
}

func httpVersion(table string) []byte {
	rows, err := db.Query("SELECT http_version, count(*) as count FROM " + table + " GROUP BY http_version ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
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
			return jsonError("DB query error", err)
		}
		httpVersion = append(httpVersion, HTTPVersion{Version: version, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(httpVersion)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}

type RequestURL struct {
	URL       string `json:"url"`
	PV        int64  `json:"pv"`
	UV        int64  `json:"uv"`
	Bandwidth int64  `json:"bandwidth"`
}

func requestURL(table string) []byte {
	rows, err := db.Query("SELECT url_path, count(*) as pv, count(distinct(ip)), sum(content_size) FROM " + table + " WHERE url_is_static='0' GROUP BY url_path ORDER BY pv DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	var (
		requestURL []RequestURL
		url        string
		pv         int64
		uv         int64
		bandwidth  int64
	)
	for rows.Next() {
		err := rows.Scan(&url, &pv, &uv, &bandwidth)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		requestURL = append(requestURL, RequestURL{URL: url, PV: pv, UV: uv, Bandwidth: bandwidth})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(requestURL)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}

type StaticFile struct {
	File      string `json:"file"`
	Count     int64  `json:"count"`
	Size      int64  `json:"size"`
	Bandwidth int64  `json:"bandwidth"`
}

func staticFile(table string) []byte {
	rows, err := db.Query("SELECT url_path, count(*) as count, sum(content_size) FROM " + table + " WHERE url_is_static='1' and response_code='200' GROUP BY url_path ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	var (
		staticFile []StaticFile
		file       string
		count      int64
		bandwidth  int64
	)
	for rows.Next() {
		err := rows.Scan(&file, &count, &bandwidth)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		size := bandwidth / count // Average size is OK
		staticFile = append(staticFile, StaticFile{File: file, Count: count, Size: size, Bandwidth: bandwidth})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(staticFile)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}
