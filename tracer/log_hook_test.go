package tracer

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewSkyWalkingHook(t *testing.T) {
	InitTracer("mock", "127.0.0.1:80")
	hook := NewSkyWalkingHook(tracer)
	hook.Levels()
	_ = hook.Fire(logrus.NewEntry(logrus.StandardLogger()))
}
