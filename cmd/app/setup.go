package main

import (
	"io"
	"t-mk-opentrace/config"
	"t-mk-opentrace/ext/es-driver/es"
	mongo2 "t-mk-opentrace/ext/mongo-driver/mongo"
	"time"

	"github.com/go-errors/errors"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

// // Host h
// var Host = "localhost"
//
// // Port p
// var Port = "6831"
// var senderAddr = Host + ":" + Port

// NewOpenTraceClient 连接
func NewOpenTraceClient() (p opentracing.Tracer, i io.Closer) {
	conf := config.Get().OpenTrace
	senderAddr := conf.Host + ":" + conf.Port
	sender, err := jaeger.NewUDPTransport(senderAddr, 0)
	if err != nil {
		panic(err)
	}
	reportOpt := jaeger.ReporterOptions.BufferFlushInterval(1 * time.Second)
	reporter := jaeger.NewRemoteReporter(
		sender,
		reportOpt)
	serviceName := "t-mk-openTrace-16006c"
	tracer, closer := jaeger.NewTracer(
		serviceName,
		jaeger.NewConstSampler(true),
		reporter,
		jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
	)
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

// InitDBs 连接数据库
func InitDBs() {
	err := config.Init()
	if err != nil {
		panic(errors.New(err))
	}
	conf := *config.Get()

	// 初始化mongo连接
	c, err := mongo2.InitClient(conf.ToMongoMKBizConfig())
	if err != nil {
		panic(err)
	}
	if c == nil {
		panic(errors.New("mk_biz mongo 连接失败"))
	}

	c, err = mongo2.InitClient(conf.ToMongoMKWatConfig())
	if err != nil {
		panic(err)
	}
	if c == nil {
		panic(errors.New("mk_wat mongo 连接失败"))
	}

	c, err = mongo2.InitClient(conf.ToMongoWPConfig())
	if err != nil {
		panic(err)
	}
	if c == nil {
		panic(errors.New("wp mongo 连接失败"))
	}
	if err := es.Init(conf.ToESConfig()); err != nil {
		panic(err)
	}
}
