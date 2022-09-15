package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
		want logrus.Level
	}{
		{name: "info1", args: args{level: "info"}, want: logrus.InfoLevel},
		{name: "info2", args: args{level: "iNfO"}, want: logrus.InfoLevel},
		{name: "null", args: args{level: "null"}, want: logrus.InfoLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogLevel(tt.args.level); got != tt.want {
				t.Errorf("LogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
