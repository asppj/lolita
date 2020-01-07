package grpc

import (
	"t-mk-opentrace/api/proto/plan"
	g "t-mk-opentrace/ext/grpc-driver/grpc"
	pr "t-mk-opentrace/interval/plan"
)

// RegisterRPC 注册rpc服务
func RegisterRPC(s *g.Server) {
	plan.RegisterServiceServer(s, &pr.RPCPlan{})
}
