package loganalyzer

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	log "loganalyzer/loganalyzer/logging"
)

type Bandwidth struct {
	Time      string `json:"time"`
	Bandwidth int64  `json:"bandwidth"`
}

func bandwidthDaily() []byte {
	return bandwidthQuery("SELECT DATE(time) as date, sum(content_size) FROM log GROUP BY date ORDER BY date ASC")
}

func bandwidthHourly(date string) []byte {
	return bandwidthQuery("SELECT HOUR(TIME(time)) as hour, sum(content_size) FROM log WHERE DATE(time)=? GROUP BY hour ORDER BY hour ASC", date)
}

func bandwidthMonthly() []byte {
	return bandwidthQuery("SELECT DATE_FORMAT(time,'%Y-%m') AS ym, sum(content_size) FROM log GROUP BY ym ORDER BY ym ASC")
}

func bandwidthQuery(query string, args ...interface{}) []byte {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
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
			return []byte(`{"status": "failed"}`)
		}
		bandwidth = append(bandwidth, Bandwidth{Time: t, Bandwidth: b})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(bandwidth)
	if err != nil {
		log.Errorln("Bandwidth json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
