package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Info("hello")
	Warn("world")
	Error("!!!")
}
func TestLoggerf(t *testing.T) {
	Infof("hello %s", "yes")
	Warnf("world %s", "!")
	Errorf("!!! %s", "no")
}
