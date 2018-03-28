package loganalyzer

import (
	"encoding/json"
	"time"

	log "loganalyzer/loganalyzer/logging"
)

type Summary struct {
	FileName  string    `json:"file_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	PageViews int64     `json:"page_views"`
	UserViews int64     `json:"user_views"`
	Bandwidth float64   `json:"bandwidth"`
}

func (s Summary) json() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		log.Errorf("Summary json marshal error: %v", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
