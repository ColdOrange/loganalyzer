package loganalyzer

import (
	"encoding/json"

	log "loganalyzer/loganalyzer/logging"
)

type UserViews struct {
	Time string `json:"time"`
	UV   int64  `json:"uv"`
}

func userViewsDaily(table string) []byte {
	return userViewsQuery("SELECT DATE(time) as date, count(distinct(ip)) FROM " + table + " GROUP BY date ORDER BY date ASC")
}

func userViewsHourly(table, date string) []byte {
	return userViewsQuery("SELECT HOUR(TIME(time)) as hour, count(distinct(ip)) FROM "+table+" WHERE DATE(time)=? GROUP BY hour ORDER BY hour ASC", date)
}

func userViewsMonthly(table string) []byte {
	return userViewsQuery("SELECT DATE_FORMAT(time,'%Y-%m') AS ym, count(distinct(ip)) FROM " + table + " GROUP BY ym ORDER BY ym ASC")
}

func userViewsQuery(query string, args ...interface{}) []byte {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	var (
		userViews []UserViews
		t         string
		uv        int64
	)
	for rows.Next() {
		err := rows.Scan(&t, &uv)
		if err != nil {
			log.Errorln("DB query error:", err)
			return []byte(`{"status": "failed"}`)
		}
		userViews = append(userViews, UserViews{Time: t, UV: uv})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(userViews)
	if err != nil {
		log.Errorln("UserViews json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
