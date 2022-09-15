package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type color uint8

const (
	red     = color(iota + 91)
	green   //	绿
	yellow  //	黄
	blue    // 	蓝
	magenta //	洋红
	cyan    // 青色
	white   //   白色
)

func setColor(s string, color color) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, s)
}

var ColorMap = map[log.Level]string{
	log.TraceLevel: setColor("TRACE", white),
	log.DebugLevel: setColor("DEBUG", blue),
	log.InfoLevel:  setColor("INFO ", green),
	log.WarnLevel:  setColor("WARN ", yellow),
	log.ErrorLevel: setColor("ERROR", red),
	log.FatalLevel: setColor("FATAL", cyan),
	log.PanicLevel: setColor("PANIC", magenta),
}

func SetLevelColor(level log.Level) string {
	if ans, ok := ColorMap[level]; ok {
		return ans
	}
	return setColor("INFO ", green)
}

var logLevel = map[string]log.Level{
	"debug":   log.DebugLevel,
	"info":    log.InfoLevel,
	"warn":    log.WarnLevel,
	"warning": log.WarnLevel,
	"error":   log.ErrorLevel,
}

func LogLevel(level string) log.Level {
	if level, ok := logLevel[strings.ToLower(level)]; ok {
		return level
	}
	return log.InfoLevel
}
