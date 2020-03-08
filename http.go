package mollie

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

func NewHTTPClient(n time.Duration) *http.Client {

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   n * time.Second,
			KeepAlive: n * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: n * time.Second,

		ExpectContinueTimeout: n * time.Second,
		ResponseHeaderTimeout: n * time.Second,
		MaxIdleConns:          3,
		MaxConnsPerHost:       2,
	}
	cli := &http.Client{
		Transport: tr,
		Timeout:   n * time.Second,
	}

	return cli
}

type params map[string]interface{}

type request struct {
	method   string
	endpoint string
	query    url.Values
	form     url.Values
	header   http.Header
	body     io.Reader
	fullURL  string
}

// setFormParam set param with key/value to request form body
func (r *request) setFormParam(key string, value interface{}) *request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setFormParams set params with key/values to request form body
func (r *request) setFormParams(m params) *request {
	for k, v := range m {
		r.setFormParam(k, v)
	}
	return r
}

func (r *request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
	return nil
}
