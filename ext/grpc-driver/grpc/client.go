package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

const defaultTimeOut = 3

// Dial client连接server
func Dial(target string, opts ...DialOption) (*ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	// TODO opentrace-go
	return grpc.Dial(target, opts...)
}

// DefaultContext 默认
func DefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeOut*time.Second)
}
