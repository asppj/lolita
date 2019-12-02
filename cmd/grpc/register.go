package grpc

import (
	"t-mk-opentrace/api/proto/plan"
	"t-mk-opentrace/api/proto/task"
	g "t-mk-opentrace/ext/grpc-driver/grpc"
	pr "t-mk-opentrace/interval/plan"
	task2 "t-mk-opentrace/interval/task"
)

// RegisterRPC 注册rpc服务
func RegisterRPC(s *g.Server) {
	task.RegisterTaskServiceServer(s, &task2.Tb{})
	plan.RegisterServiceServer(s, &pr.RPCPlan{})
}
