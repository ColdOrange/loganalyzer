package loganalyzer

import (
	"encoding/json"

	log "loganalyzer/loganalyzer/logging"
)

type ReferringSite struct {
	Site string `json:"site"`
	PV   int64  `json:"pv"`
	UV   int64  `json:"uv"`
}

func referringSite(table string) []byte {
	rows, err := db.Query("SELECT referrer_site, count(*) as count, count(distinct(ip)) FROM " + table + " WHERE referrer_site != '' GROUP BY referrer_site ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	var (
		referringSites []ReferringSite
		site           string
		pv             int64
		uv             int64
	)
	for rows.Next() {
		err := rows.Scan(&site, &pv, &uv)
		if err != nil {
			log.Errorln("DB query error:", err)
			return []byte(`{"status": "failed"}`)
		}
		referringSites = append(referringSites, ReferringSite{Site: site, PV: pv, UV: uv})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(referringSites)
	if err != nil {
		log.Errorln("ReferringSite json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}

type ReferringURL struct {
	URL string `json:"url"`
	PV  int64  `json:"pv"`
	UV  int64  `json:"uv"`
}

func referringURL(table string) []byte {
	rows, err := db.Query("SELECT referrer_path, count(*) as count, count(distinct(ip)) FROM " + table + " WHERE referrer_path != '' GROUP BY referrer_path ORDER BY count DESC")
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	var (
		referringURLs []ReferringURL
		url           string
		pv            int64
		uv            int64
	)
	for rows.Next() {
		err := rows.Scan(&url, &pv, &uv)
		if err != nil {
			log.Errorln("DB query error:", err)
			return []byte(`{"status": "failed"}`)
		}
		referringURLs = append(referringURLs, ReferringURL{URL: url, PV: pv, UV: uv})
	}
	err = rows.Err()
	if err != nil {
		log.Errorln("DB query error:", err)
		return []byte(`{"status": "failed"}`)
	}

	data, err := json.Marshal(referringURLs)
	if err != nil {
		log.Errorln("ReferringURL json marshal error:", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
