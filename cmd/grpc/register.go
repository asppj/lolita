package grpc

import (
	g "github.com/asppj/lolita/ext/grpc-driver/grpc"
	pr "github.com/asppj/lolita/interval/plan"
	"github.com/asppj/lolita/proto/plan"
)

// RegisterRPC 注册rpc服务
func RegisterRPC(s *g.Server) {
	plan.RegisterServiceServer(s, &pr.RPCPlan{})
}
