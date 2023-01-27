package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Print(
	color string,
	level string,
	data interface{},
	err error,
	tags []string,
) {
	type json struct {
		level string
		time  string
		tags  string
		err   string
		data  interface{}
	}
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	s := json{
		level: level,
		time:  time.Now().Format(time.RFC3339),
		tags:  strings.Join(tags, ","),
		err:   errStr,
		data:  data,
	}
	fmt.Fprintf(os.Stdout, "\033[%sm%#v\033[0m\n", color, s)
}

func Debug(data interface{}, tags ...string) {
	Print(Color.blue, Level.debug, data, nil, tags)
}

func Info(data interface{}, tags ...string) {
	Print(Color.green, Level.info, data, nil, tags)
}

func Warn(data interface{}, tags ...string) {
	Print(Color.yellow, Level.warn, data, nil, tags)
}

func Error(data interface{}, err error, tags ...string) {
	Print(Color.red, Level.err, data, err, tags)
}
