package tracer

import (
	"net/http"
	"testing"
)

func handler(http.ResponseWriter, *http.Request) {
	return
}

type MockHttpResponseWriter struct {
}

func (m MockHttpResponseWriter) Header() http.Header {
	panic("implement me")
}

func (m MockHttpResponseWriter) Write(_ []byte) (int, error) {
	panic("implement me")
}

func (m MockHttpResponseWriter) WriteHeader(_ int) {
	panic("implement me")
}

func TestNewHttpTracer(t *testing.T) {
	InitTracer("mock", "127.0.0.1:80")
	warp := NewHttpTracer(tracer)
	warp.WrapHandle(handler)
	handler(&MockHttpResponseWriter{}, nil)
}
