package loganalyzer

import (
	"io/ioutil"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"

	log "loganalyzer/loganalyzer/logging"
)

type Handler struct {
	mux map[string]func(http.ResponseWriter, *http.Request)
	reg map[string]*regexp.Regexp
}

func NewHandler() *Handler {
	Handler := &Handler{}
	Handler.mux = make(map[string]func(http.ResponseWriter, *http.Request))
	Handler.reg = make(map[string]*regexp.Regexp)
	return Handler
}

func (Handler *Handler) Bind(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	Handler.mux[pattern] = handler
	Handler.reg[pattern] = regexp.MustCompile(pattern)
}

// Static file server
var staticFilesHandler = http.StripPrefix("/static/", http.FileServer(http.Dir(path.Join(ProjectPath, "assets/static"))))

func (Handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zero := time.Now()
	defer func() {
		duration := time.Since(zero)
		log.Infof("%s %s [%.3fms]", r.Method, r.URL, duration.Seconds()*1000)
	}()

	if strings.HasPrefix(r.URL.Path, "/static/") { // static files
		staticFilesHandler.ServeHTTP(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/api/") { // api service
		for pattern, regex := range Handler.reg {
			if regex.MatchString(r.URL.Path) {
				Handler.mux[pattern](w, r)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	} else { // otherwise, just send index page and let front end do the route
		index, _ := ioutil.ReadFile(path.Join(ProjectPath, "assets/static/index.html")) // TODO: in memory?
		w.Write(index)
	}
}
