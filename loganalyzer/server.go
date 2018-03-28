package loganalyzer

import (
	"net/http"

	log "loganalyzer/loganalyzer/logging"
)

func handlerSummary(w http.ResponseWriter, _ *http.Request) {
	w.Write(summary())
}

func NewServer(addr string) *http.Server {
	handler := NewHandler()
	handler.Bind("/summary", handlerSummary)

	log.Infof("Server started listening on [%v]", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
