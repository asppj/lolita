package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"t-mk-opentrace/cmd/grpc"
	"t-mk-opentrace/cmd/http"
)

// CleanUp 清理
func CleanUp() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			shutDown()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

// shutDown 关闭
func shutDown() {
	grpc.RPCShotDown()
	if err := http.GinShutDown(); err != nil {
		log.Println("关闭http失败")
	}
}
