package tracer

import (
	"github.com/SkyAPM/go2sky"
	"github.com/sirupsen/logrus"
	v3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	"time"
)

const (
	Trace                   = "__trace"
	componentIDGOHttpClient = 5005
	componentIDGOHttpServer = 5004
)

type TraceInfo struct {
	Operator  string
	Tags      map[go2sky.Tag]string
	Peer      string
	SpanLayer v3.SpanLayer
	Component int32
}

type SkyWalkingHook struct {
	tracer *go2sky.Tracer
}

func NewSkyWalkingHook(tracer *go2sky.Tracer) *SkyWalkingHook {
	return &SkyWalkingHook{tracer: tracer}
}

func (s SkyWalkingHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
	}
}

func (s SkyWalkingHook) Fire(entry *logrus.Entry) error {
	if s.tracer == nil {
		return nil
	}
	if entry.Context == nil {
		return nil
	}
	if entry.Data == nil {
		return nil
	}
	if !HasTraceId(entry.Context) {
		return nil
	}
	trace, ok := entry.Data[Trace]
	if !ok {
		return nil
	}
	traceInfo, ok := trace.(TraceInfo)
	if !ok {
		return nil
	}
	span, ctx, err := s.tracer.CreateLocalSpan(entry.Context)
	if err != nil {
		return err
	}
	entry.Context = ctx
	span.SetOperationName(traceInfo.Operator)
	span.SetSpanLayer(traceInfo.SpanLayer)
	span.SetPeer(traceInfo.Peer)
	span.SetComponent(componentIDGOHttpServer)

	for tag, value := range traceInfo.Tags {
		span.Tag(tag, value)
	}
	if traceInfo.Component != 0 {
		span.SetComponent(traceInfo.Component)
	}
	span.Log(time.Now(), entry.Message)
	span.End()
	return nil
}
