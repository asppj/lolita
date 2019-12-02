package main

import (
	"context"
	"t-mk-opentrace/ext/http-driver/requests"
	"t-mk-opentrace/ext/log-driver/log"
)

// NewRequest test get /
func NewRequest() {
	ctx := context.Background()
	v := &struct {
		StatusCode interface{} `json:"status_code"`
	}{}
	if err := requests.Get(ctx, "http://localhost:6006", nil, v); err != nil {
		log.Warn(err)
	}
	log.Info(v)
}

func main() {
	NewRequest()
}
