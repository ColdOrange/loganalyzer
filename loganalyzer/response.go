package loganalyzer

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "loganalyzer/loganalyzer/logging"
)

type StatusCode struct {
	StatusCode string `json:"statusCode"`
	Count      int64  `json:"count"`
}

func statusCode(table string) []byte {
	rows, err := db.Query("SELECT response_code, count(*) as count FROM " + table + " GROUP BY response_code ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
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
			return jsonError("DB query error", err)
		}
		statusText := strconv.Itoa(code) + " " + http.StatusText(code)
		statusCode = append(statusCode, StatusCode{StatusCode: statusText, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(statusCode)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}

type ResponseTime struct {
	TimeRange string `json:"timeRange"`
	Count     int64  `json:"count"`
}

func responseTime(table string) []byte {
	rows, err := db.Query("SELECT CASE " +
		"WHEN response_time < 50 THEN '<50ms' " +
		"WHEN response_time >= 50 AND response_time < 100 THEN '50~100ms' " +
		"WHEN response_time >= 100 AND response_time < 200 THEN '100~200ms' " +
		"WHEN response_time >= 200 AND response_time < 300 THEN '200~300ms' " +
		"WHEN response_time >= 300 AND response_time < 400 THEN '300~400ms'" +
		"WHEN response_time >= 400 AND response_time < 500 THEN '400~500ms' " +
		"ELSE '>500ms' " +
		"END AS time_range, count(*) AS count " +
		"FROM " + table + " " +
		"GROUP BY time_range " +
		"ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	var (
		responseTime []ResponseTime
		timeRange    string
		count        int64
	)
	for rows.Next() {
		err := rows.Scan(&timeRange, &count)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		responseTime = append(responseTime, ResponseTime{TimeRange: timeRange, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(responseTime)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}

type ResponseURL struct {
	URL    string `json:"url"`
	PV     int64  `json:"pv"`
	Avg    int64  `json:"avg"`
	StdDev int64  `json:"stdDev"`
}

func responseURL(table string) []byte {
	rows, err := db.Query("SELECT url_path, COUNT(*) AS pv, AVG(response_time), STD(response_time) FROM " + table + " WHERE url_is_static='0' GROUP BY url_path ORDER BY pv DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	var (
		responseURL []ResponseURL
		url         string
		pv          int64
		avg         float64
		stdDev      float64
	)
	for rows.Next() {
		err := rows.Scan(&url, &pv, &avg, &stdDev)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		responseURL = append(responseURL, ResponseURL{URL: url, PV: pv, Avg: int64(avg), StdDev: int64(stdDev)})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(responseURL)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}
