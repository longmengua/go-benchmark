package log

var Color = struct {
	red    string
	green  string
	yellow string
	blue   string
}{
	// color code reference: https://en.wikipedia.org/wiki/ANSI_escape_code
	red:    "31",
	green:  "32",
	yellow: "33",
	blue:   "34",
}

var Level = struct {
	debug string
	info  string
	warn  string
	err   string
}{
	debug: "debug",
	info:  "info",
	warn:  "warn",
	err:   "err",
}
