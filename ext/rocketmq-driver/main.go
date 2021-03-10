package main

import (
	"cgo-test/rocketmq"
	"context"
	"fmt"
	"time"
)

func main() {
	
	err := rocketmq.InitProducer()
	// if err != nil {
	// 	panic(err)
	// }
	err = rocketmq.InitPushConsumer()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case t := <-ticker.C:
			rocketmq.SendOneWay(ctx, fmt.Sprintf("send:%v", t))
		}
	}
	
}
