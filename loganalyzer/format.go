package loganalyzer

// Log format fields
const (
	IP            = "IP"            // Remote IP
	Time          = "Time"          // Time
	RequestMethod = "RequestMethod" // Request HTTP method
	RequestURL    = "RequestURL"    // Request URL (including params)
	HTTPVersion   = "HTTPVersion"   // HTTP Version
	ResponseCode  = "ResponseCode"  // Response status code
	ResponseTime  = "ResponseTime"  // Time cost to serve this request
	ContentSize   = "ContentSize"   // Byte sent in response
	UserAgent     = "UserAgent"     // User agent
	Referer       = "Referer"       // Referer
)

// Time format, copy from golang `time` package.
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
