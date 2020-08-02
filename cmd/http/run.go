package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/asppj/lolita/cmd/http/middleware"
	"github.com/asppj/lolita/config"

	"github.com/gin-gonic/gin"
)

// var mode = "debug"
//
// // Port http端口
// var Port = "16006"
//
// // Host 地址
// var Host = ""

var _defaultServer *http.Server

// ginInitRouter s
func ginInitRouter() *gin.Engine {
	mode := config.Get().Mode
	gin.SetMode(string(mode))
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	if gin.Mode() == gin.DebugMode {
		gin.ForceConsoleColor()
	}
	return r
}

// GinInitServer init
func GinInitServer() error {
	routerHandel := ginInitRouter()
	// tracer, conn := middleware.NewOpenTraceClient()
	// defer func() {
	// 	if err := conn.Close(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	conf := config.Get().Server
	Host := conf.Host
	Port := conf.Port
	routerHandel.Use(middleware.WithTimeOut())
	routerHandel.Use(middleware.OpenTraceMiddleware())
	routerHandel.Use(middleware.PromMiddle())
	if err := RegisterRouter(routerHandel); err != nil {
		panic(err)
	}
	_defaultServer = &http.Server{
		Addr:         Host + ":" + Port,
		Handler:      routerHandel,
		ReadTimeout:  5000 * time.Second,
		WriteTimeout: 10000 * time.Second,
	}
	_defaultServer.RegisterOnShutdown(func() {
		log.Printf("关闭http服务...%s:%s", Host, Port)
	})
	var host = Host
	if host == "" {
		host = "0.0.0.0"
	}
	log.Printf("正在启动http服务，监听%s:%s\n", host, Port)
	return _defaultServer.ListenAndServe()
}

// GinShutDown 关闭
func GinShutDown() error {
	log.Println("正在关闭http服务")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := _defaultServer.Shutdown(ctx); err != nil {
		return err
	}
	log.Println("http服务成功关闭")
	return nil
}
