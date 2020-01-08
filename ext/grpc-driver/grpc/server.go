package grpc

import (
	"github.com/asppj/t-go-opentrace/ext/grpc-driver/grpc/middleware"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	"google.golang.org/grpc"
)

// NewServer 创建服务
func NewServer(opt ...ServerOption) *Server {
	opt = append(
		opt,
		// serverOption(),
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(opentracing.GlobalTracer())),
	)
	return grpc.NewServer(opt...)
}

// serverOption serverOption
func serverOption() grpc.ServerOption {
	return grpc.UnaryInterceptor(middleware.GRPCServerInterceptor)
}
