package logger

// InitLogrusStandardLogger
// 性能指标: 100线程一共写1000000行长度为100的日志
// 开启GoId,不开启stdout, 13秒完成,qps : 75000/s
// 不开启GoId,不开启stdout, 5秒完成,qps : 200000/s

import (
	"context"
	"fmt"
	"github.com/SkyAPM/go2sky"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func NewRotateFile(path string) (*Writer, error) {
	var writer *rotatelogs.RotateLogs
	options := []rotatelogs.Option{
		rotatelogs.WithMaxAge(time.Duration(24*7) * time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24) * time.Hour),
	}
	if runtime.GOOS == "linux" {
		options = append(options, rotatelogs.WithLinkName(path))
	}
	writer, err := rotatelogs.New(path+".%Y%m%d", options...)
	if err != nil {
		return nil, err
	}
	buf := NewWriter(writer)
	go AutoFlush(buf)
	return buf, nil
}

func NewLogFormatter() *LogFormatter {
	return &LogFormatter{}
}

func NewEsLogFormatter() *EsLogFormatter {
	return &EsLogFormatter{}
}

func AutoFlush(buf *Writer) {
	for buf != nil {
		time.Sleep(time.Millisecond * 33) //
		err := buf.Flush()
		if err != nil {
			return
		}
	}
}

type LogFormatter struct {
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05.000")
	var file string
	var packageName string
	var functionName string
	var line int
	if entry.HasCaller() {
		packageName = filepath.Base(filepath.Dir(entry.Caller.File))
		file = filepath.Base(entry.Caller.File)
		tmp := strings.Split(entry.Caller.Func.Name(), ".")
		if len(tmp) > 0 {
			functionName = tmp[len(tmp)-1]
		}
		line = entry.Caller.Line
	}
	var traceId = GetTraceId(entry.Context)
	if len(entry.Data) == 0 {
		msg := fmt.Sprintf("%s [%s][Trace:%s][%s/%s:%d][FUNC:%s] %s\n",
			SetLevelColor(entry.Level),
			timestamp, traceId,

			packageName, file, line,
			functionName,
			entry.Message,
		)
		return []byte(msg), nil
	} else {
		msg := fmt.Sprintf("%s [%s][Trace:%s][%s/%s:%d][FUNC:%s] %s %+v\n",
			SetLevelColor(entry.Level),
			timestamp, traceId,
			packageName, file, line,
			functionName,
			entry.Message,
			entry.Data,
		)
		return []byte(msg), nil
	}
}

type EsLogFormatter struct {
}

func (f *EsLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("20060102.150405.000000")
	var file string
	var packageName string
	var functionName string
	var line int
	if entry.HasCaller() {
		packageName = filepath.Base(filepath.Dir(entry.Caller.File))
		file = filepath.Base(entry.Caller.File)
		tmp := strings.Split(entry.Caller.Func.Name(), ".")
		if len(tmp) > 0 {
			functionName = tmp[len(tmp)-1]
		}
		line = entry.Caller.Line
	}
	var traceId = GetTraceId(entry.Context)
	level := strings.ToUpper(fmt.Sprintf("%s", entry.Level))
	var logBody string
	if len(entry.Data) == 0 {
		logBody = fmt.Sprintf("%s%s traceId:%s [%s/%s/%s line:%d] %s\n",
			level,
			timestamp,
			traceId,
			packageName, file, functionName, line,
			entry.Message,
		)
	} else {
		logBody = fmt.Sprintf("%s%s traceId:%s [%s/%s/%s line:%d] %s %+v\n",
			level,
			timestamp,
			traceId,
			packageName, file, functionName, line,
			entry.Message,
			entry.Data,
		)
	}
	return []byte(logBody), nil
}

func GetTraceId(ctx context.Context) string {
	if ctx == nil {
		return go2sky.EmptyTraceID
	}
	trace := go2sky.TraceID(ctx)
	if trace != go2sky.EmptyTraceID {
		return trace
	}
	t, _ := ctx.Value("traceId").(string)
	if t == "" {
		return go2sky.EmptyTraceID
	}
	return t
}
