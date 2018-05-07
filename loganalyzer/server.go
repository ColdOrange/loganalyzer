package loganalyzer

import (
	"io/ioutil"
	"net/http"

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

// Summary
func handlerSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := summary()
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
		value := pageViewsDaily()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerPageViewsHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := pageViewsHourly(r.URL.Query().Get("date"))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerPageViewsMonthly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := pageViewsMonthly()
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
		value := userViewsDaily()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerUserViewsHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := userViewsHourly(r.URL.Query().Get("date"))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerUserViewsMonthly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := userViewsMonthly()
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
		value := bandwidthDaily()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerBandwidthHourly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := bandwidthHourly(r.URL.Query().Get("date"))
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerBandwidthMonthly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := bandwidthMonthly()
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
		value := requestMethod()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerHTTPVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := httpVersion()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerRequestURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := requestURL()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerStaticFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := staticFile()
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
		value := statusCode()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerResponseTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := responseTime()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerResponseURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := responseURL()
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
		value := operatingSystem()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := device()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerBrowser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := browser()
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
		value := referringSite()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func handlerReferringURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if cache.Exist(r.RequestURI) {
		w.Write(cache.Get(r.RequestURI))
	} else {
		value := referringURL()
		w.Write(value)
		cache.Set(r.RequestURI, value)
	}
}

func NewServer(addr string) *http.Server {
	handler := NewHandler()
	handler.Bind("/api/config/database", handlerDatabaseConfig)
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
	handler.Bind("/api/http-version", handlerHTTPVersion)
	handler.Bind("/api/request-url", handlerRequestURL)
	handler.Bind("/api/static-file", handlerStaticFile)
	handler.Bind("/api/status-code", handlerStatusCode)
	handler.Bind("/api/response-time", handlerResponseTime)
	handler.Bind("/api/response-url", handlerResponseURL)
	handler.Bind("/api/user-agent/os", handlerOperatingSystem)
	handler.Bind("/api/user-agent/device", handlerDevice)
	handler.Bind("/api/user-agent/browser", handlerBrowser)
	handler.Bind("/api/referer/site", handlerReferringSite)
	handler.Bind("/api/referer/url", handlerReferringURL)

	cache = NewCache()

	log.Infof("Server started listening on [%v]", addr)
	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
