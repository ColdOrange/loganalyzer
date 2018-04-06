package loganalyzer

import (
	"net/http"

	log "loganalyzer/loganalyzer/logging"
)

func handlerSummary(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json") // TODO: do we really need this?
	w.Write(summary())
}

func NewServer(addr string) *http.Server {
	handler := NewHandler()
	handler.Bind("/api/summary", handlerSummary)

	log.Infof("Server started listening on [%v]", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
