package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm/log"
)

type XormLogger struct {
	logger  *logrus.Logger
	level   log.LogLevel
	showSql bool
}

func NewXormLogger(logger *logrus.Logger) *XormLogger {
	return &XormLogger{logger: logger}
}

func (x *XormLogger) Debug(v ...interface{}) {
	x.logger.Debug(v...)
}

func (x *XormLogger) Debugf(format string, v ...interface{}) {
	x.logger.Debugf(format, v...)
}

func (x *XormLogger) Error(v ...interface{}) {
	x.logger.Error(v...)
}

func (x *XormLogger) Errorf(format string, v ...interface{}) {
	x.logger.Errorf(format, v...)
}

func (x *XormLogger) Info(v ...interface{}) {
	x.logger.Info(v...)
}

func (x *XormLogger) Infof(format string, v ...interface{}) {
	x.logger.Infof(format, v...)
}

func (x *XormLogger) Warn(v ...interface{}) {
	x.logger.Warn(v...)
}

func (x *XormLogger) Warnf(format string, v ...interface{}) {
	x.logger.Warnf(format, v...)
}

func (x *XormLogger) Level() log.LogLevel {
	return x.level
}

func (x *XormLogger) SetLevel(l log.LogLevel) {
	x.level = l
}

func (x *XormLogger) ShowSQL(show ...bool) {
	for _, v := range show {
		x.showSql = v
	}
}
func (x *XormLogger) IsShowSQL() bool {
	return x.showSql
}
