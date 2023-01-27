package log_test

import (
	"errors"
	"go-benchmark-demo/log"
	l "log"
	"testing"
)

func assertion(t *testing.T, assert any, result any) {
	if assert != result {
		t.Errorf("Error, %s, %s", result, assert)
	}
}

func TestLog(t *testing.T) {
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error", errors.New("Error demo"))

	n := struct {
		name string
	}{
		name: "name",
	}
	log.Debug(n)
	log.Info(n)
	log.Warn(n)
	log.Error(n, errors.New("Error demo"))

	tt := struct {
		name string
	}{
		name: "name",
	}
	log.Debug(tt, "debug test", "category flag", "anchor", "and so on...")
	log.Info(tt, "info test", "category flag", "anchor", "and so on...")
	log.Warn(tt, "info test", "category flag", "anchor", "and so on...")
	log.Error(tt, errors.New("Error demo"), "error test", "category flag", "anchor", "and so on...")
	l.Fatal("err")
}
