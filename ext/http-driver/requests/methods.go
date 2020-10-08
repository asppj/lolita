package requests

import (
	"context"
	"io"
	"net/http"
)

// Cookie set header Cookie
const Cookie = "Cookie"
const contentType = "Content-Type"

// Get get
func Get(ctx context.Context, url string, body io.Reader, v interface{}) error {
	c := &RequestClient{}
	return c.Get(ctx, url, body, v)
}

// Post post
func Post(ctx context.Context, url string, body io.Reader, v interface{}) error {
	c := &RequestClient{}
	return c.Post(ctx, url, body, v)
}

// Delete Delete
func Delete(ctx context.Context, url string, body io.Reader, v interface{}) error {
	c := &RequestClient{}
	return c.Delete(ctx, url, body, v)
}

// Put Put
func Put(ctx context.Context, url string, body io.Reader, v interface{}) error {
	c := &RequestClient{}
	return c.Put(ctx, url, body, v)
}

// Patch Patch
func Patch(ctx context.Context, url string, body io.Reader, v interface{}) error {
	c := &RequestClient{}
	return c.Patch(ctx, url, body, v)
}

func setHeader(r *http.Request, k string, v string) {
	r.Header.Set(k, v)
}
