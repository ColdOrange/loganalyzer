package loganalyzer

import (
	"io/ioutil"
	"net/http"
	"path"

	log "loganalyzer/loganalyzer/logging"
)

func handlerIndex(w http.ResponseWriter, _ *http.Request) {
	index, _ := ioutil.ReadFile(path.Join(ProjectPath, "assets/static/index.html")) // TODO: in memory?
	w.Write(index)
}

func handlerSummary(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Set("Content-Type", "application/json")
	w.Write(summary())
}

func NewServer(addr string) *http.Server {
	handler := NewHandler()
	handler.Bind("/", handlerIndex)
	handler.Bind("/api/summary", handlerSummary)

	log.Infof("Server started listening on [%v]", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
