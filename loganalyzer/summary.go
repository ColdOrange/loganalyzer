package loganalyzer

import (
	"encoding/json"
	"time"
	"strconv"

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

func (s Summary) analyze(line int, fields []string) {
	s.PageViews++
	for j := 1; j <= len(config.LogFormat); j++ {
		switch config.LogFormat[j-1] {
		case IP:
			s.UserViews++

		case Time:
			timestamp, err := time.Parse(config.TimeFormat, fields[j])
			if err != nil {
				log.Warnf("Log file [%s] format wrong at line %d", Time, line)
				continue
			}
			if line == 0 {
				s.StartTime = timestamp
			}
			s.EndTime = timestamp // TODO: performance? or a better way?

		case RequestMethod:
		case RequestURL:
		case HTTPVersion:
		case ResponseCode:
		case ResponseTime:
		case ContentSize:
			if fields[j] == "-" {
				fields[j] = "0" // TODO: do we need this, or can just ignore?
			}
			size, err := strconv.Atoi(fields[j])
			if err != nil {
				log.Warnf("Log file [%s] format wrong at line %d", ContentSize, line)
				continue
			}
			s.Bandwidth += float64(size)

		case UserAgent:
		case Referer:
		}
	}
}

func (s Summary) json() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		log.Errorf("Summary json marshal error: %v", err)
		return []byte(`{"status": "failed"}`)
	}
	return data
}
