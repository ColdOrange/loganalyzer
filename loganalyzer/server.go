package loganalyzer

import (
	"net/http"

	log "loganalyzer/loganalyzer/logging"
)

func handlerSummary(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json") // TODO: do we really need this?
	w.Write(summary())
}

func handlerPageViewsDaily(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(pageViewsDaily())
}

func handlerPageViewsHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(pageViewsHourly(r.URL.Query().Get("date")))
}

func handlerPageViewsMonthly(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(pageViewsMonthly())
}

func NewServer(addr string) *http.Server {
	handler := NewHandler()
	handler.Bind("/api/summary", handlerSummary)
	handler.Bind("/api/page-views/daily", handlerPageViewsDaily)
	handler.Bind("/api/page-views/hourly", handlerPageViewsHourly)
	handler.Bind("/api/page-views/monthly", handlerPageViewsMonthly)

	log.Infof("Server started listening on [%v]", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
