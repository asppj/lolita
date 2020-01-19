package grpc

import (
	g "github.com/asppj/t-go-opentrace/ext/grpc-driver/grpc"
	pr "github.com/asppj/t-go-opentrace/interval/plan"
	"github.com/asppj/t-go-opentrace/proto/plan"
)

// RegisterRPC 注册rpc服务
func RegisterRPC(s *g.Server) {
	plan.RegisterServiceServer(s, &pr.RPCPlan{})
}
