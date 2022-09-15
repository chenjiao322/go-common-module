package tracer

import (
	"context"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/sirupsen/logrus"
)

var tracer *go2sky.Tracer

func HasTraceId(ctx context.Context) bool {
	return go2sky.TraceID(ctx) != go2sky.EmptyTraceID && go2sky.TraceID(ctx) != ""
}

func InitTracer(serverName, addr string) {
	r, err := reporter.NewGRPCReporter(addr)
	if err != nil {
		logrus.Fatal("InitTracer Error")
	}
	tracer, err = go2sky.NewTracer(serverName, go2sky.WithReporter(r))
	if err != nil {
		logrus.Fatal("InitTracer Error")
	}
}

func SetTracer(t *go2sky.Tracer) {
	tracer = t
}

func GetTracer() *go2sky.Tracer {
	return tracer
}
