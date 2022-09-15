package tracer

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/propagation"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"strconv"
)

func empty() {
	// Nothing
}

type HttpTracer struct {
	tracer      *go2sky.Tracer
	projectName string
}

func NewHttpTracer(tracer *go2sky.Tracer) *HttpTracer {
	return &HttpTracer{tracer: tracer}
}

func (h *HttpTracer) BeforeRequest(_ http.ResponseWriter, r *http.Request) (func(), *http.Request) {
	if h.tracer == nil {
		return empty, r
	}
	if r.Header.Get(propagation.Header) == "" && r.Header.Get(propagation.HeaderCorrelation) == "" {
		return empty, r
	}
	span, ctx, err := h.tracer.CreateEntrySpan(r.Context(),
		"HttpHandleStart",
		func(headerKey string) (string, error) {
			return r.Header.Get(headerKey), nil
		})
	if err != nil {
		return empty, r
	}
	span.SetPeer(h.projectName)
	span.Tag(go2sky.TagURL, r.RequestURI)
	span.Tag(go2sky.TagHTTPMethod, r.Method)
	span.SetComponent(componentIDGOHttpServer)
	return func() { span.End() }, r.WithContext(ctx)
}

func (h *HttpTracer) AfterRequest(w http.ResponseWriter, r *http.Request) func() {
	if h.tracer == nil {
		return empty
	}
	if !HasTraceId(r.Context()) {
		return empty
	}
	span, err := h.tracer.CreateExitSpan(r.Context(), "HttpHandleEnd", h.projectName,
		func(headerKey, headerValue string) error {
			w.Header().Add(headerKey, headerValue)
			return nil
		})
	if err != nil {
		return empty
	}
	span.Tag(go2sky.TagURL, r.RequestURI)
	span.Tag(go2sky.TagHTTPMethod, r.Method)
	span.SetComponent(componentIDGOHttpServer)
	status := reflect.ValueOf(w).Elem().FieldByName("status")
	if status.IsValid() {
		span.Tag(go2sky.TagStatusCode, strconv.Itoa(int(status.Int())))
	}
	return func() { span.End() }
}

func (h *HttpTracer) WrapHandle(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Errorf("http request error: %v", err)
			}
		}()
		beforeEnd, r := h.BeforeRequest(w, r)
		handler(w, r)
		afterEnd := h.AfterRequest(w, r)
		beforeEnd()
		afterEnd()
	}
}
