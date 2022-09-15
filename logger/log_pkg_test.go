package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLog(t *testing.T) {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(NewEsLogFormatter())
	logrus.Infof("123")
	logrus.SetFormatter(NewLogFormatter())
	logrus.Infof("123")

	logrus.SetReportCaller(false)
	logrus.SetFormatter(NewEsLogFormatter())
	logrus.WithFields(map[string]interface{}{"123": "312"}).Infof("123")
	logrus.SetFormatter(NewLogFormatter())
	logrus.WithFields(map[string]interface{}{"123": "312"}).Infof("321")

	_, err := NewRotateFile("/tmp/123")
	if err != nil {
		panic(err)
	}
}
