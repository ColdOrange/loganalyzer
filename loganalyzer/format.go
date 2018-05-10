package loganalyzer

// Log format fields
var LogFields = []string{
	"IP",
	"Time",
	"RequestMethod",
	"RequestURL",
	"HTTPVersion",
	"ResponseCode",
	"ResponseTime",
	"ContentSize",
	"UserAgent",
	"Referrer",
}

// Corresponding DB table fields
var LogTableFields = map[string][]string{
	"IP":            {"ip"},
	"Time":          {"time"},
	"RequestMethod": {"request_method"},
	"RequestURL":    {"url_path", "url_query", "url_is_static"},
	"HTTPVersion":   {"http_version"},
	"ResponseCode":  {"response_code"},
	"ResponseTime":  {"response_time"},
	"ContentSize":   {"content_size"},
	"UserAgent":     {"ua_browser", "ua_os", "ua_device"},
	"Referrer":       {"referrer_site", "referrer_path", "referrer_query"},
}

// Time format, copy from golang `time` package.
// You can also utilize the constants list in golang `time/format.go` to create new formats.
const (
	Apache      = "02/Jan/2006:15:04:05 -0700"
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700"
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	Stamp       = "Jan _2 15:04:05"
	StampMilli  = "Jan _2 15:04:05.000"
	StampMicro  = "Jan _2 15:04:05.000000"
	StampNano   = "Jan _2 15:04:05.000000000"
)
