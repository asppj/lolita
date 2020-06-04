package test_ai

import (
	"net/http"
	"testing"
)

// TestSinaURLShort 测试
func TestSinaURLShort(t *testing.T) {
	originurl := "https://github.com"
	resp, err := http.Get(originurl)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.Status)
}
