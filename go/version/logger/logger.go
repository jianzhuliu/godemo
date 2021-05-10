package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"version/utils"
)

//接口定义
type ILogger interface {
	Output(calldepth int, s string) error
	SetPrefix(prefix string)
}

//默认日志接口
var stdLog ILogger = log.New(os.Stdout, "log", log.LstdFlags|log.Lmsgprefix|log.Lshortfile)

func Info(args ...interface{}) {
	stdLog.SetPrefix("[info]")
	stdLog.Output(2, fmt.Sprintln(args...))
}

func Warn(args ...interface{}) {
	stdLog.SetPrefix("[warn]")
	stdLog.Output(2, fmt.Sprintln(args...))
}

func Error(args ...interface{}) {
	stdLog.SetPrefix("[error]")
	stdLog.Output(2, fmt.Sprintln(args...))
}

func Exit(args ...interface{}) {
	stdLog.SetPrefix("[exit]")
	stdLog.Output(2, fmt.Sprintln(args...))
	os.Exit(1)
}

func Infof(format string, args ...interface{}) {
	stdLog.SetPrefix("[info]")
	format = addLine(format)
	stdLog.Output(2, fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	stdLog.SetPrefix("[warn]")
	format = addLine(format)
	stdLog.Output(2, fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	stdLog.SetPrefix("[error]")
	format = addLine(format)
	stdLog.Output(2, fmt.Sprintf(format, args...))
}

func Exitf(format string, args ...interface{}) {
	stdLog.SetPrefix("[exit]")
	format = addLine(format)
	stdLog.Output(2, fmt.Sprintf(format, args...))
	os.Exit(1)
}

//添加换行符
func addLine(format string) string {
	if !strings.HasSuffix(format, "\n") {
		format = utils.JoinStrings(format, "\n")
	}
	return format
}
