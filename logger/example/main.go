package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/log"
	"gitlab.geinc.cn/services/go-common-module/logger"
	"gitlab.geinc.cn/services/go-common-module/tracer"
	"io"
	"os"
	"time"
)

type Config struct {
	LogPath        string // 日志的路径是目标是一个文件而不是文件夹,
	StdOut         bool   // 是否需要标准输出到console,开发环境推荐开,正式环境推荐关
	SkyWalkingAddr string // SkyWalking的ip+端口比如 172.16.0.127:11800
	ServerName     string // 服务名, 比如mpService
}

func setup(cfg Config) {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(logger.NewEsLogFormatter())
	//logrus.SetFormatter(logger.NewLogFormatter())
	buf, err := logger.NewRotateFile(cfg.LogPath)
	if err != nil {
		panic(err)
	}
	if cfg.StdOut {
		logrus.SetOutput(io.MultiWriter(buf, os.Stdout))
	} else {
		logrus.SetOutput(buf)
	}
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	setup(Config{
		LogPath:        "./",
		StdOut:         true,
		SkyWalkingAddr: "172.16.0.127:11800",
		ServerName:     "logger_example",
	})
	time.Sleep(time.Millisecond)
	ctx := context.WithValue(context.Background(), "traceId", "123")
	logrus.WithContext(ctx).Info("123")
	logrus.WithContext(ctx).
		WithField(tracer.Trace, tracer.TraceInfo{}).
		Infof("123 %v", "123")
	engine, _ := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	engine.SetLogger(logger.NewXormLogger(logrus.StandardLogger()))
	engine.SetLogLevel(log.LOG_DEBUG)
	engine.ShowSQL(true)
	engine.NewSession().Exec("select 1")
	engine.Exec("select 1")
}
