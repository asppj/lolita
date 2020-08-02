package app

import (
	"sync"

	"github.com/asppj/lolita/cmd/grpc"
	"github.com/asppj/lolita/cmd/http"
	ant "github.com/asppj/lolita/ext/ants-driver/ants"
	"github.com/asppj/lolita/ext/log-driver/log"
)

func init() {
	InitDBs()
}

// Main app
func Main() {
	// tracer, conn := middleware.NewOpenTraceClient()
	_, closer := NewOpenTraceClient()
	defer func() {
		if err := closer.Close(); err != nil {
			log.Warn("opentracer collecter 关闭失败", err)
		}
	}()
	defer ant.Release()
	group := sync.WaitGroup{}
	defer group.Wait()
	// http
	group.Add(1)
	ant.Go(func() {
		if err := http.GinInitServer(); err != nil {
			panic(err)
		}
		group.Done()
	})
	// grpc
	group.Add(1)
	ant.Go(func() {
		if err := grpc.RPCServer(); err != nil {
			panic(err)
		}
		group.Done()
	})
	// 启动内网
	// group.Add(1)
	// ant.Go(func() {
	// 	server.NewServer()
	// 	group.Done()
	// })
	ant.Go(func() {
		CleanUp()
	})
}
