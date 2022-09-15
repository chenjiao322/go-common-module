package tracer

import (
	"context"
	"github.com/xormplus/xorm/contexts"
	"testing"
)

func TestNewSkyWalkingXormHook(t *testing.T) {
	InitTracer("mock", "127.0.0.1:80")
	hook := NewSkyWalkingXormHook(tracer)
	ctx := contexts.NewContextHook(context.Background(), "select 1", nil)
	_, _ = hook.BeforeProcess(ctx)
	_ = hook.AfterProcess(ctx)
}
