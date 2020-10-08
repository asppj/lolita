package requests

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

const clientTimeOut = 5

// HeaderOption func set header
type HeaderOption = func(r *http.Request)

// RequestClient RequestClient
type RequestClient struct {
	ContentType string
	Cookie      string
	URL         string
	opts        []HeaderOption
	ins         *http.Client
	insOnce     sync.Once
}

// SetClient init
func (c *RequestClient) SetClient(client *http.Client) {
	c.ins = client
}

// client 初始化client
func (c *RequestClient) client() *http.Client {
	if c.ins == nil {
		c.insOnce.Do(func() {
			c.ins = &http.Client{
				Timeout: clientTimeOut * time.Second,
			}
		})
	}
	return c.ins
}

// SetHeader ss
func (c *RequestClient) SetHeader(k, v string) {
	c.opts = append(c.opts, func(r *http.Request) {
		setHeader(r, k, v)
	})
}

// Request 请求
func (c *RequestClient) Request(
	ctx context.Context,
	url string,
	in io.Reader,
	out interface{},
	method string,
	fnOpts ...HeaderOption) error {
	opts := make([]HeaderOption, 0, len(fnOpts)+len(c.opts))
	opts = append(opts, fnOpts...)
	opts = append(opts, c.opts...)
	return Request(ctx, c.client(), url, in, out, method, opts...)
}

// Post sd
func (c *RequestClient) Post(ctx context.Context, url string, in io.Reader, out interface{}) error {
	fn := func(r *http.Request) {
		setHeader(r, contentType, "application/x-www-form-urlencoded")
	}
	return c.Request(ctx, url, in, out, "POST", fn)
}

// Get sd
func (c *RequestClient) Get(ctx context.Context, url string, in io.Reader, out interface{}) error {
	return c.Request(ctx, url, in, out, "GET")
}

// Delete Delete
func (c *RequestClient) Delete(ctx context.Context, url string, in io.Reader, out interface{}) error {
	fn := func(r *http.Request) {
		setHeader(r, contentType, "application/json")
	}
	return c.Request(ctx, url, in, out, "DELETE", fn)
}

// Put Put
func (c *RequestClient) Put(ctx context.Context, url string, in io.Reader, out interface{}) error {
	fn := func(r *http.Request) {
		setHeader(r, contentType, "application/json")
	}
	return c.Request(ctx, url, in, out, "PUT", fn)
}

// Patch Patch
func (c *RequestClient) Patch(ctx context.Context, url string, in io.Reader, out interface{}) error {
	fn := func(r *http.Request) {
		setHeader(r, contentType, "application/json")
	}
	return c.Request(ctx, url, in, out, "PATCH", fn)
}
