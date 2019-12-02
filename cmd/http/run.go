package http

import (
	"context"
	"log"
	"net/http"
	"t-mk-opentrace/cmd/http/middleware"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/gin-gonic/gin"
)

var mode = "debug"

// Port http端口
var Port = "6006"

// Host 地址
var Host = ""

var _defaultServer *http.Server

// ginInitRouter s
func ginInitRouter() *gin.Engine {
	gin.SetMode(mode)
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
	routerHandel.Use(middleware.OpenTraceMiddleware(opentracing.GlobalTracer()))
	if err := RegisterRouter(routerHandel); err != nil {
		panic(err)
	}
	_defaultServer = &http.Server{
		Addr:         Host + ":" + Port,
		Handler:      routerHandel,
		ReadTimeout:  500 * time.Second,
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
