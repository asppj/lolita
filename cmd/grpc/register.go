package grpc

import (
	g "t-mk-opentrace/ext/grpc-driver/grpc"
	task2 "t-mk-opentrace/interval/task"
	"t-mk-opentrace/proto/task"
)

// RegisterRPC 注册rpc服务
func RegisterRPC(s *g.Server) {
	task.RegisterTaskServiceServer(s, &task2.Tb{})
}
