package main

import (
	g "t-mk-opentrace/cmd/grpc"
	"t-mk-opentrace/ext/grpc-driver/grpc"
	"t-mk-opentrace/ext/log-driver/log"
	"t-mk-opentrace/proto/task"
)

var rpcServer = g.RPCAddr

// TaskDial dial
func TaskDial() (*grpc.ClientConn, error) {
	return grpc.Dial("127.0.0.1" + rpcServer)
}

// TestTask 测试rpc连接
func TestTask() {
	cc, err := TaskDial()
	if err != nil {
		panic(err)
	}
	c := task.NewTaskServiceClient(cc)
	ctx, cancel := grpc.DefaultContext()
	defer cancel()
	res, err := c.Search(ctx, &task.TaskRequest{
		Request:   "req",
		NameReq:   "name",
		AgeReq:    18,
		LocalTime: "timeLocal",
	})
	if err != nil {
		log.Warn(err)
		return
	}
	log.Info(res.AgeRes, res.NameRes, res.GetResponse())
}

func main() {
	TestTask()
}
