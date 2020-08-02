package grpc

import (
	"context"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/asppj/lolita/ext/grpc-driver/grpc/middleware"

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
			grpc_middleware.ChainUnaryClient(
				otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
				grpc_prometheus.UnaryClientInterceptor,
			)),
		grpc.WithStreamInterceptor(
			grpc_middleware.ChainStreamClient(
				otgrpc.OpenTracingStreamClientInterceptor(opentracing.GlobalTracer()),
				grpc_prometheus.StreamClientInterceptor,
			)),
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
