package app

import (
	"io"
	"time"

	"github.com/asppj/t-go-opentrace/config"

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
	_ = *config.Get()
}
