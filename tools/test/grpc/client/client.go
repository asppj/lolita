package main

import (
	"github.com/asppj/t-go-opentrace/api/proto/task"
	g "github.com/asppj/t-go-opentrace/cmd/grpc"
	"github.com/asppj/t-go-opentrace/ext/grpc-driver/grpc"
	"github.com/asppj/t-go-opentrace/ext/log-driver/log"
)

var rpcServer = g.RPCAddr

// TaskDial dial
func TaskDial() (*grpc.ClientConn, error) {
	return grpc.Dial("192.168.253.73" + rpcServer)
}

// TestTask 测试rpc连接
func TestTask() {
	cc, err := TaskDial()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := cc.Close(); err != nil {
			log.Warn(err)
		}
	}()
	c := task.NewTaskServiceClient(cc)
	ctx, cancel := grpc.DefaultContext()
	defer cancel()
	res, err := c.PlanDetail(ctx, &task.PlanRequest{
		PlanID: "planID-test-1",
	})
	if err != nil {
		log.Warn(err)
		return
	}
	log.Info(res.PlanID, res.GetCode())
}

func main() {
	TestTask()
}
