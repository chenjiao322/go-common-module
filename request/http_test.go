package request

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

type Response struct {
	Method string              `json:"method"`
	Path   string              `json:"path"`
	Header map[string][]string `json:"header"`
	Body   string              `json:"body"`
}

func (r *Response) ToJson() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return b
}

func NewResponse(request *http.Request) *Response {
	a, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	header := request.Header
	header.Del("Accept-Encoding")
	header.Del("Content-Length")
	header.Del("User-Agent")
	return &Response{
		Method: request.Method,
		Path:   request.RequestURI,
		Header: request.Header,
		Body:   string(a),
	}
}

func MockServer() func() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(NewResponse(request).ToJson())
	})
	http.HandleFunc("/slow", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Second * 2)
		_, _ = writer.Write(NewResponse(request).ToJson())
	})
	server := http.Server{Addr: "127.0.0.1:35000"}

	go func() {
		_ = server.ListenAndServe()
	}()
	return func() { _ = server.Close() }
}

func DefaultHeader() map[string][]string {
	return map[string][]string{}
}

var quick = "http://127.0.0.1:35000/"
var slow = "http://127.0.0.1:35000/slow"
var bad = "http://127.0.0.1:35198/"

func TestHttp(t *testing.T) {
	closeServer := MockServer()
	time.Sleep(time.Second)
	defer closeServer()
	tests := []struct {
		name    string
		request *Http
		want    *Response
		wantErr bool
	}{
		{name: "get", request: NewDefault().Get(quick), want: &Response{Method: "GET", Path: "/", Header: DefaultHeader()}},
		{name: "timeout", request: NewDefault().SetTimeout(time.Millisecond * 50).Get(slow), wantErr: true},
		{name: "bad uri", request: NewDefault().Get(" " + slow), wantErr: true},
		{name: "bad server", request: NewDefault().SetTimeout(time.Millisecond * 50).Get(bad), wantErr: true},
		{name: "get", request: NewDefault().Get(quick).SetJson(map[string]string{"1": "2"}).SetParams(Param{"1": "2"}),
			want: &Response{Method: "GET", Path: "/?1=2", Header: map[string][]string{"Content-Type": {"application/json"}}, Body: `{"1":"2"}`},
		},
		{name: "bad json", request: NewDefault().Get(quick).SetJson(map[struct{}]struct{}{{}: {}}), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tt.request.Do(nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("IP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			response := new(Response)
			binary, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf(err.Error())
			}
			err = json.Unmarshal(binary, response)
			if err != nil {
				t.Errorf(err.Error())
			}
			if !reflect.DeepEqual(response, tt.want) {
				t.Errorf("different\n %+v\n %+v", response, tt.want)
			}
		})
	}

	// Test Fetch Json

	j := &Response{}
	_, err := NewDefault().Get(quick).Fetch(j).Do(nil)
	if err != nil {
		t.Errorf("request error")
	}
	if !reflect.DeepEqual(j, &Response{
		Method: "GET",
		Path:   "/",
		Header: map[string][]string{},
		Body:   "",
	}) {
		t.Errorf("wrong reponse")
	}

}

func TestBuild(t *testing.T) {
	tests := []struct {
		name string
		got  *Http
		want *Http
	}{
		{name: "get", got: NewDefault().Get(quick), want: &Http{client: http.DefaultClient, method: GET, url: quick}},
		{name: "post", got: NewDefault().Post(quick), want: &Http{client: http.DefaultClient, method: POST, url: quick}},
		{name: "put", got: NewDefault().Put(quick), want: &Http{client: http.DefaultClient, method: PUT, url: quick}},
		{name: "delete", got: NewDefault().Delete(quick), want: &Http{client: http.DefaultClient, method: DELETE, url: quick}},
		{name: "Patch", got: NewDefault().Patch(quick), want: &Http{client: http.DefaultClient, method: PATCH, url: quick}},
		{name: "head", got: NewDefault().Head(quick), want: &Http{client: http.DefaultClient, method: HEAD, url: quick}},
		{name: "option", got: NewDefault().Option(quick), want: &Http{client: http.DefaultClient, method: OPTIONS, url: quick}},
		{name: "Trace", got: NewDefault().Trace(quick), want: &Http{client: http.DefaultClient, method: TRACE, url: quick}},
		{name: "setJson", got: NewDefault().Get(quick).SetJson(map[string]string{}),
			want: &Http{client: http.DefaultClient, header: map[string][]string{"Content-Type": {"application/json"}}, method: GET, url: quick, json: map[string]string{}}},
		{name: "SetAll", got: NewDefault().Get(quick).
			SetTimeout(time.Second).
			SetBody([]byte("123")).
			SetHeader(http.Header{"a": {"123"}}).SetParams(Param{"123": "321"}),
			want: &Http{
				client:  http.DefaultClient,
				timeout: time.Second,
				header:  http.Header{"a": {"123"}},
				param:   Param{"123": "321"},
				body:    []byte("123"),
				method:  GET,
				url:     quick,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("IsInt() = %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func TestParam_ToString(t *testing.T) {
	tests := []struct {
		name string
		p    Param
		want string
	}{
		{name: "1", p: Param{"123": "321"}, want: "123=321"},
		{name: "2", p: Param{"123": "321", "124": "321"}, want: "123=321&124=321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
