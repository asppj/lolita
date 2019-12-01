package grpc

import "google.golang.org/grpc"

// NewServer 创建服务
func NewServer(opt ...ServerOption) *Server {
	return grpc.NewServer(opt...)
}
