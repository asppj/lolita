package grpc

import (
	"context"
	"t-mk-opentrace/ext/grpc-driver/grpc/middleware"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
)

const defaultTimeOut = 3

// DialStream stream
func DialStream(target string, opts ...DialOption) (*ClientConn, error) {
	opts = append(opts,
		grpc.WithInsecure(),
		// grpc.WithStreamInterceptor()
	)
	return grpc.Dial(target, opts...)
}

// Dial client连接server
func Dial(target string, opts ...DialOption) (*ClientConn, error) {
	opts = append(
		opts,
		grpc.WithInsecure(),
		// clientDialOption(),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(opentracing.GlobalTracer())),
	)
	return grpc.Dial(target, opts...)
}

// DefaultContext 默认
func DefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeOut*time.Second)
}

// clientDialOption 拦截器
func clientDialOption() grpc.DialOption {
	return grpc.WithUnaryInterceptor(middleware.GRPCClientInterceptor)
}
