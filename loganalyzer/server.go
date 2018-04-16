package loganalyzer

import (
	"net/http"

	log "loganalyzer/loganalyzer/logging"
)

// Summary
func handlerSummary(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(summary())
}

// Page Views
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

// User Views
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

// Bandwidth
func handlerBandwidthDaily(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(bandwidthDaily())
}

func handlerBandwidthHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(bandwidthHourly(r.URL.Query().Get("date")))
}

func handlerBandwidthMonthly(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(bandwidthMonthly())
}

// Request
func handlerRequestMethod(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(requestMethod())
}

func handlerRequestURL(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(requestURL())
}

func handlerHTTPVersion(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(httpVersion())
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
	handler.Bind("/api/bandwidth/daily", handlerBandwidthDaily)
	handler.Bind("/api/bandwidth/hourly", handlerBandwidthHourly)
	handler.Bind("/api/bandwidth/monthly", handlerBandwidthMonthly)
	handler.Bind("/api/request-method", handlerRequestMethod)
	handler.Bind("/api/request-url", handlerRequestURL)
	handler.Bind("/api/http-version", handlerHTTPVersion)

	log.Infof("Server started listening on [%v]", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
