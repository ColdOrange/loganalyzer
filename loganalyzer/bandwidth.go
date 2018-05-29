package loganalyzer

import (
	"encoding/json"

	log "loganalyzer/loganalyzer/logging"
)

type Bandwidth struct {
	Time      string `json:"time"`
	Bandwidth int64  `json:"bandwidth"`
}

func bandwidthDaily(table string) []byte {
	return bandwidthQuery("SELECT DATE(time) as date, sum(content_size) FROM " + table + " GROUP BY date ORDER BY date ASC")
}

func bandwidthHourly(table, date string) []byte {
	return bandwidthQuery("SELECT HOUR(TIME(time)) as hour, sum(content_size) FROM "+table+" WHERE DATE(time)=? GROUP BY hour ORDER BY hour ASC", date)
}

func bandwidthMonthly(table string) []byte {
	return bandwidthQuery("SELECT DATE_FORMAT(time,'%Y-%m') AS ym, sum(content_size) FROM " + table + " GROUP BY ym ORDER BY ym ASC")
}

func bandwidthQuery(query string, args ...interface{}) []byte {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	var (
		bandwidth []Bandwidth
		t         string
		b         int64
	)
	for rows.Next() {
		err := rows.Scan(&t, &b)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		bandwidth = append(bandwidth, Bandwidth{Time: t, Bandwidth: b})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(bandwidth)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}
