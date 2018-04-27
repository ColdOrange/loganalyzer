package loganalyzer

import (
	"encoding/json"

	log "loganalyzer/loganalyzer/logging"
)

type PageViews struct {
	Time string `json:"time"`
	PV   int64  `json:"pv"`
}

func pageViewsDaily() []byte {
	return pageViewsQuery("SELECT DATE(time) as date, count(*) FROM log GROUP BY date ORDER BY date ASC")
}

func pageViewsHourly(date string) []byte {
	return pageViewsQuery("SELECT HOUR(TIME(time)) as hour, count(*) FROM log WHERE DATE(time)=? GROUP BY hour ORDER BY hour ASC", date)
}

func pageViewsMonthly() []byte {
	return pageViewsQuery("SELECT DATE_FORMAT(time,'%Y-%m') AS ym, count(*) FROM log GROUP BY ym ORDER BY ym ASC")
}

func pageViewsQuery(query string, args ...interface{}) []byte {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
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
			return []byte(`{"status": "failed"}`)
		}
		pageViews = append(pageViews, PageViews{Time: t, PV: pv})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(pageViews)
	if err != nil {
		log.Errorln("PageViews json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
