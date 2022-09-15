package logger

import (
	"context"
	gormLogger "gorm.io/gorm/logger"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	x := New()
	ctx := context.Background()
	x.Error(ctx, "1")
	x.Info(ctx, "1")
	x.LogMode(gormLogger.Error)
	x.Trace(ctx, time.Now(), func() (string, int64) { return "", 0 }, nil)
	x.Warn(ctx, "1")
}
