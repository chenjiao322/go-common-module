package tracer

import (
	"context"
	"github.com/SkyAPM/go2sky"
	"github.com/xormplus/xorm/contexts"
	"time"
)

// XORM用hook
// 使用方法
// engine, err = xorm.NewEngine("mysql", "%s:%s@tcp(%s:%d)/%s?charset=%s")
// engine.AddHook(tracer.NewSkyWalkingXormHook(tracer.GetTracer()))

type SkyWalkingXormHook struct {
	tracer *go2sky.Tracer
}

func NewSkyWalkingXormHook(tracer *go2sky.Tracer) *SkyWalkingXormHook {
	return &SkyWalkingXormHook{tracer: tracer}
}

func (s SkyWalkingXormHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	if s.tracer == nil {
		return c.Ctx, nil
	}
	return c.Ctx, nil
}

func (s SkyWalkingXormHook) AfterProcess(c *contexts.ContextHook) error {
	if s.tracer == nil {
		return nil
	}
	if !HasTraceId(c.Ctx) {
		return nil
	}
	span, _, err := s.tracer.CreateLocalSpan(c.Ctx)
	if err != nil {
		return err
	}
	if c.ExecuteTime.Microseconds() > 1000 { // sql执行时间大于1s的视为慢sql
		span.Tag("slow", "slow")
	}
	span.SetOperationName("SQL")
	span.Log(time.Now(), c.SQL)
	span.SetComponent(componentIDGOHttpServer)
	span.End()
	return nil
}
