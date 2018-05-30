package loganalyzer

import (
	"encoding/json"

	log "github.com/ColdOrange/loganalyzer/loganalyzer/logging"
)

type PageViews struct {
	Time string `json:"time"`
	PV   int64  `json:"pv"`
}

func pageViewsDaily(table string) []byte {
	return pageViewsQuery("SELECT DATE(time) as date, count(*) FROM " + table + " GROUP BY date ORDER BY date ASC")
}

func pageViewsHourly(table, date string) []byte {
	return pageViewsQuery("SELECT HOUR(TIME(time)) as hour, count(*) FROM "+table+" WHERE DATE(time)=? GROUP BY hour ORDER BY hour ASC", date)
}

func pageViewsMonthly(table string) []byte {
	return pageViewsQuery("SELECT DATE_FORMAT(time,'%Y-%m') AS ym, count(*) FROM " + table + " GROUP BY ym ORDER BY ym ASC")
}

func pageViewsQuery(query string, args ...interface{}) []byte {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	var (
		pageViews []PageViews
		t         string
		pv        int64
	)
	for rows.Next() {
		err := rows.Scan(&t, &pv)
		if err != nil {
			log.Errorln("DB query error:", err)
			return jsonError("DB query error", err)
		}
		pageViews = append(pageViews, PageViews{Time: t, PV: pv})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return jsonError("DB query error", err)
	}

	data, err := json.Marshal(pageViews)
	if err != nil {
		log.Errorln("Json marshal error:", err)
		return jsonError("Json marshal error", err)
	}
	return data
}
