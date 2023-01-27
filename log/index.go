package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func PrintWithColor(
	color string,
	str string,
) {
	fmt.Fprintf(os.Stdout, fmt.Sprintf("\033[%sm%s\033[0m\n", color, str))
}

func print(
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
	print(Color.Blue, Level.Debug, data, nil, tags)
}

func Info(data interface{}, tags ...string) {
	print(Color.Green, Level.Info, data, nil, tags)
}

func Warn(data interface{}, tags ...string) {
	print(Color.Yellow, Level.Warn, data, nil, tags)
}

func Error(data interface{}, err error, tags ...string) {
	print(Color.Red, Level.Error, data, err, tags)
}
