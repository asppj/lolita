package grpc

import (
	"github.com/asppj/t-go-opentrace/ext/grpc-driver/grpc/middleware"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
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
			grpc_middleware.ChainUnaryServer(
				otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
				grpc_prometheus.UnaryServerInterceptor,
			)),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				otgrpc.OpenTracingStreamServerInterceptor(opentracing.GlobalTracer()),
				grpc_prometheus.StreamServerInterceptor,
			)),
	)
	server := grpc.NewServer(opt...)
	grpc_prometheus.Register(server)
	return server
}

// serverOption serverOption
func serverOption() grpc.ServerOption {
	return grpc.UnaryInterceptor(middleware.GRPCServerInterceptor)
}
