package loganalyzer

import (
	"io/ioutil"
	"net/http"
	"path"
	"strings"
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

var staticFilesHandler http.Handler

// Static file server
func init() {
	fileServer := http.FileServer(http.Dir(path.Join(ProjectPath, "assets/static")))
	staticFilesHandler = http.StripPrefix("/static/", fileServer)
}

func (Handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zero := time.Now()
	defer func() {
		duration := time.Since(zero)
		log.Infof("%s %s [%.3fms]", r.Method, r.URL, duration.Seconds()*1000)
	}()

	if strings.HasPrefix(r.URL.Path, "/static/") { // static files
		staticFilesHandler.ServeHTTP(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/api/") { // api service
		if handler, ok := Handler.mux[r.URL.Path]; ok {
			handler(w, r)
		} else {
			w.WriteHeader(404)
		}
	} else { // otherwise, just send index page and let front end do the route
		index, _ := ioutil.ReadFile(path.Join(ProjectPath, "assets/static/index.html")) // TODO: in memory?
		w.Write(index)
	}
}
