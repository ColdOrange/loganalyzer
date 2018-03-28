package loganalyzer

import (
	"net/http"
	"time"

	log "loganalyzer/loganalyzer/logging"
)

type Handler struct {
	mux map[string]func(http.ResponseWriter, *http.Request)
}

func NewHandler() *Handler {
	Handler := &Handler{}
	Handler.mux = make(map[string]func(http.ResponseWriter, *http.Request))
	return Handler
}

func (Handler *Handler) Bind(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	Handler.mux[pattern] = handler
}

func (Handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zero := time.Now()
	defer func() {
		duration := time.Since(zero)
		log.Infof("%s [%.3fms] %s", r.Method, duration.Seconds()*1000, r.URL)
	}()

	if handler, ok := Handler.mux[r.URL.String()]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(404)
	}
}
