package loganalyzer

import (
	"net/http"

	log "loganalyzer/loganalyzer/logging"
)

func handlerSummary(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

func handlerUserViewsDaily(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(userViewsDaily())
}

func handlerUserViewsHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(userViewsHourly(r.URL.Query().Get("date")))
}

func handlerUserViewsMonthly(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(userViewsMonthly())
}

func NewServer(addr string) *http.Server {
	handler := NewHandler()
	handler.Bind("/api/summary", handlerSummary)
	handler.Bind("/api/page-views/daily", handlerPageViewsDaily)
	handler.Bind("/api/page-views/hourly", handlerPageViewsHourly)
	handler.Bind("/api/page-views/monthly", handlerPageViewsMonthly)
	handler.Bind("/api/user-views/daily", handlerUserViewsDaily)
	handler.Bind("/api/user-views/hourly", handlerUserViewsHourly)
	handler.Bind("/api/user-views/monthly", handlerUserViewsMonthly)

	log.Infof("Server started listening on [%v]", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
