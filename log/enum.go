package log

var Color = struct {
	Red    string
	Green  string
	Yellow string
	Blue   string
}{
	// color code reference: https://en.wikipedia.org/wiki/ANSI_escape_code
	Red:    "31",
	Green:  "32",
	Yellow: "33",
	Blue:   "34",
}

var Level = struct {
	Debug string
	Info  string
	Warn  string
	Error string
}{
	Debug: "debug",
	Info:  "info",
	Warn:  "warn",
	Error: "err",
}
