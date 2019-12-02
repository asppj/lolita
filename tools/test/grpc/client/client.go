package main

import (
	"t-mk-opentrace/api/proto/task"
	g "t-mk-opentrace/cmd/grpc"
	"t-mk-opentrace/ext/grpc-driver/grpc"
	"t-mk-opentrace/ext/log-driver/log"
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
