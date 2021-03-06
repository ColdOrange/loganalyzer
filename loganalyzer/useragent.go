package loganalyzer

import (
	"encoding/json"

	log "github.com/ColdOrange/loganalyzer/loganalyzer/logging"
)

type OperatingSystem struct {
	OS    string `json:"os"`
	Count int64  `json:"count"`
}

func operatingSystem(table string) []byte {
	rows, err := db.Query("SELECT ua_os, count(*) as count FROM " + table + " WHERE ua_os != 'Unknown' GROUP BY ua_os ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}
	defer rows.Close()

	var (
		operatingSystem []OperatingSystem
		os              string
		count           int64
	)
	for rows.Next() {
		err := rows.Scan(&os, &count)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		operatingSystem = append(operatingSystem, OperatingSystem{OS: os, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(operatingSystem)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}

type Device struct {
	Device string `json:"device"`
	Count  int64  `json:"count"`
}

func device(table string) []byte {
	rows, err := db.Query("SELECT ua_device, count(*) as count FROM " + table + " WHERE ua_device != 'Unknown' GROUP BY ua_device ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}
	defer rows.Close()

	var (
		devices []Device
		device  string
		count   int64
	)
	for rows.Next() {
		err := rows.Scan(&device, &count)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		devices = append(devices, Device{Device: device, Count: count})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(devices)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}

type Browser struct {
	Browser string `json:"browser"`
	PV      int64  `json:"pv"`
	UV      int64  `json:"uv"`
}

func browser(table string) []byte {
	rows, err := db.Query("SELECT ua_browser, count(*) as count, count(distinct(ip)) FROM " + table + " WHERE ua_browser != 'Unknown' GROUP BY ua_browser ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}
	defer rows.Close()

	var (
		browsers []Browser
		browser  string
		pv       int64
		uv       int64
	)
	for rows.Next() {
		err := rows.Scan(&browser, &pv, &uv)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		browsers = append(browsers, Browser{Browser: browser, PV: pv, UV: uv})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(browsers)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}
