package grpc

import (
	"github.com/asppj/t-mk-opentrace/api/proto/plan"
	g "github.com/asppj/t-mk-opentrace/ext/grpc-driver/grpc"
	pr "github.com/asppj/t-mk-opentrace/interval/plan"
)

// RegisterRPC 注册rpc服务
func RegisterRPC(s *g.Server) {
	plan.RegisterServiceServer(s, &pr.RPCPlan{})
}
