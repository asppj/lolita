package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/asppj/lolita/tools/nbcs/server"

	"github.com/asppj/lolita/cmd/grpc"
	"github.com/asppj/lolita/cmd/http"
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
	server.StopServer()
}
