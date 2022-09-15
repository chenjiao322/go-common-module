package request

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

type HttpMethod string

const (
	GET     = HttpMethod("GET")
	POST    = HttpMethod("POST")
	PUT     = HttpMethod("PUT")
	DELETE  = HttpMethod("DELETE")
	HEAD    = HttpMethod("HEAD")
	OPTIONS = HttpMethod("OPTIONS")
	PATCH   = HttpMethod("PATCH")
	TRACE   = HttpMethod("TRACE")
)

type Param map[string]string

type Http struct {
	client *http.Client
	method HttpMethod

	url    string
	param  Param
	header http.Header

	json interface{}
	body []byte

	// 用于接收body格式是json的返回值
	resp interface{}

	timeout time.Duration
}

func NewDefault() *Http {
	return New(http.DefaultClient)
}

func New(client *http.Client) *Http {
	return &Http{client: client}
}

func (h *Http) Get(url string) *Http {
	h.SetMethod(GET)
	return h.SetUrl(url)
}

func (h *Http) Post(url string) *Http {
	h.SetMethod(POST)
	return h.SetUrl(url)
}

func (h *Http) Option(url string) *Http {
	h.SetMethod(OPTIONS)
	return h.SetUrl(url)
}

func (h *Http) Put(url string) *Http {
	h.SetMethod(PUT)
	return h.SetUrl(url)
}

func (h *Http) Patch(url string) *Http {
	h.SetMethod(PATCH)
	return h.SetUrl(url)
}

func (h *Http) Delete(url string) *Http {
	h.SetMethod(DELETE)
	return h.SetUrl(url)
}

func (h *Http) Head(url string) *Http {
	h.SetMethod(HEAD)
	return h.SetUrl(url)
}

func (h *Http) Trace(url string) *Http {
	h.SetMethod(TRACE)
	return h.SetUrl(url)
}

func (h *Http) SetMethod(method HttpMethod) *Http {
	h.method = method
	return h
}

func (h *Http) SetUrl(url string) *Http {
	h.url = url
	return h
}

func (h *Http) SetParams(p Param) *Http {
	h.param = p
	return h
}

func (h *Http) SetHeader(header http.Header) *Http {
	h.header = header
	return h
}

func (h *Http) SetJson(j interface{}) *Http {
	if h.header == nil {
		h.header = http.Header{}
	}
	h.header.Set("Content-Type", "application/json")
	h.json = j
	return h
}

func (h *Http) SetBody(b []byte) *Http {
	h.body = b
	return h
}

func (h *Http) SetTimeout(t time.Duration) *Http {
	h.timeout = t
	return h
}

// Fetch 接受json格式的返回值,调用该接口后, resp.body 会 close.
func (h *Http) Fetch(i interface{}) *Http {
	h.resp = i
	return h
}

func (h *Http) Do(ctx context.Context) (*http.Response, error) {
	req, cancel, err := h.makeRequest(ctx)
	if err != nil {
		return nil, err
	}
	defer cancel()
	var resp *http.Response
	resp, err = h.client.Do(req)
	if err != nil {
		return nil, err
	}
	if h.resp != nil {
		defer func() { _ = resp.Body.Close() }()
		out, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(out, h.resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	return resp, nil
}

func (h *Http) GetBody() ([]byte, error) {
	if h.json != nil {
		data, err := json.Marshal(h.json)
		if err != nil {
			return []byte{}, err
		}
		h.body = data
	}
	return h.body, nil
}

func (h *Http) makeRequest(ctx context.Context) (*http.Request, func(), error) {
	body, err := h.GetBody()
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest(string(h.method), h.url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	AddParams(req, h.param)
	AddHeader(req, h.header)
	cancel := func() {
		// 如果不设置timeout,那么就不用cancel
	}
	if h.timeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, h.timeout)
		req = req.WithContext(ctx)
	}
	return req, cancel, nil
}

func AddParams(r *http.Request, param Param) {
	q := r.URL.Query()
	for k, v := range param {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()
}

func AddHeader(r *http.Request, header http.Header) {
	r.Header = header
}

func (p Param) ToString() string {
	var params [][2]string
	for k, v := range p {
		params = append(params, [2]string{k, v})
	}
	sort.Slice(params, func(i, j int) bool {
		return params[i][0] < params[j][0]
	})
	sb := strings.Builder{}
	for i := 0; i < len(params); i++ {
		k, v := params[i][0], params[i][1]
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(v)
		if i != len(params)-1 {
			sb.WriteString("&")
		}
	}
	return sb.String()
}
