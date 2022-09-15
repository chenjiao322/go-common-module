package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm/log"
	"testing"
)

func TestNewXormLogger(t *testing.T) {
	x := NewXormLogger(logrus.StandardLogger())

	x.Info()
	x.Infof("")

	x.Warn()
	x.Warnf("")

	x.Error()
	x.Errorf("")

	x.Level()
	x.SetLevel(log.LOG_DEBUG)

}
