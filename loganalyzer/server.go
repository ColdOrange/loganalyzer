package loganalyzer

import (
	"io/ioutil"
	"net/http"
	"strings"

	log "loganalyzer/loganalyzer/logging"
)

var cache *Cache

// Database Config
func handlerDatabaseConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.Write(getDBConfig())
	} else if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		w.Write(setDBConfig(body))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Log Format Config
func handlerLogFormatConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.Write(getLogConfig())
	} else if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		w.Write(setLogConfig(body))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Reports
func handlerReports(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(reports())
}

// Summary
func handlerSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) { // Should use r.RequestURI instead of r.URL.Path, because r.URL.Path doesn't contain query params
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := summary(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

// Page Views
func handlerPageViewsDaily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := pageViewsDaily(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerPageViewsHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := pageViewsHourly(getLogTableFromURL(r.URL.Path), r.URL.Query().Get("date"))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerPageViewsMonthly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := pageViewsMonthly(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

// User Views
func handlerUserViewsDaily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := userViewsDaily(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerUserViewsHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := userViewsHourly(getLogTableFromURL(r.URL.Path), r.URL.Query().Get("date"))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerUserViewsMonthly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := userViewsMonthly(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

// Bandwidth
func handlerBandwidthDaily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := bandwidthDaily(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerBandwidthHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := bandwidthHourly(getLogTableFromURL(r.URL.Path), r.URL.Query().Get("date"))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerBandwidthMonthly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := bandwidthMonthly(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

// Request
func handlerRequestMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := requestMethod(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerHTTPVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := httpVersion(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerRequestURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := requestURL(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerStaticFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := staticFile(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

// Response
func handlerStatusCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := statusCode(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerResponseTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := responseTime(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerResponseURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := responseURL(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

// User Agent
func handlerOperatingSystem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := operatingSystem(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := device(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerBrowser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := browser(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

// Referer
func handlerReferringSite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := referringSite(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerReferringURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := referringURL(getLogTableFromURL(r.URL.Path))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func getLogTableFromURL(url string) string {
	id := url[13 : 13+strings.Index(url[13:], "/")]
	return "log_" + id
}

func NewServer(addr string) *http.Server {
	handler := NewHandler()
	handler.Bind("/api/config/database", handlerDatabaseConfig)
	handler.Bind("/api/config/log-format", handlerLogFormatConfig)
	handler.Bind("/api/reports$", handlerReports) // $ for exact match
	handler.Bind("/api/reports/[0-9]+/summary", handlerSummary)
	handler.Bind("/api/reports/[0-9]+/page-views/daily", handlerPageViewsDaily)
	handler.Bind("/api/reports/[0-9]+/page-views/hourly", handlerPageViewsHourly)
	handler.Bind("/api/reports/[0-9]+/page-views/monthly", handlerPageViewsMonthly)
	handler.Bind("/api/reports/[0-9]+/user-views/daily", handlerUserViewsDaily)
	handler.Bind("/api/reports/[0-9]+/user-views/hourly", handlerUserViewsHourly)
	handler.Bind("/api/reports/[0-9]+/user-views/monthly", handlerUserViewsMonthly)
	handler.Bind("/api/reports/[0-9]+/bandwidth/daily", handlerBandwidthDaily)
	handler.Bind("/api/reports/[0-9]+/bandwidth/hourly", handlerBandwidthHourly)
	handler.Bind("/api/reports/[0-9]+/bandwidth/monthly", handlerBandwidthMonthly)
	handler.Bind("/api/reports/[0-9]+/request-method", handlerRequestMethod)
	handler.Bind("/api/reports/[0-9]+/http-version", handlerHTTPVersion)
	handler.Bind("/api/reports/[0-9]+/request-url", handlerRequestURL)
	handler.Bind("/api/reports/[0-9]+/static-file", handlerStaticFile)
	handler.Bind("/api/reports/[0-9]+/status-code", handlerStatusCode)
	handler.Bind("/api/reports/[0-9]+/response-time", handlerResponseTime)
	handler.Bind("/api/reports/[0-9]+/response-url", handlerResponseURL)
	handler.Bind("/api/reports/[0-9]+/user-agent/os", handlerOperatingSystem)
	handler.Bind("/api/reports/[0-9]+/user-agent/device", handlerDevice)
	handler.Bind("/api/reports/[0-9]+/user-agent/browser", handlerBrowser)
	handler.Bind("/api/reports/[0-9]+/referer/site", handlerReferringSite)
	handler.Bind("/api/reports/[0-9]+/referer/url", handlerReferringURL)

	cache = NewCache()

	log.Infof("Server started listening on [%v]", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
