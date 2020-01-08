package main

import (
	"context"
	"io"
	"time"

	"github.com/asppj/t-go-opentrace/ext/http-driver/requests"
	"github.com/asppj/t-go-opentrace/ext/log-driver/log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

var senderAddr = "localhost:6831"

// NewRequest test get /
func NewRequest() {
	NewOpenTraceClient()
	ctx := context.Background()
	sp := opentracing.StartSpan("t-client-http")
	defer sp.Finish()
	ctx = opentracing.ContextWithSpan(ctx, sp)
	v := &struct {
		Data interface{} `json:"data"`
	}{}
	if err := requests.Get(ctx, "http://localhost:6006/v1/pc/article", nil, v); err != nil {
		log.Warn(err)
	}
	log.Info(v)
}

func main() {
	NewRequest()
}

// NewOpenTraceClient 连接
func NewOpenTraceClient() (p opentracing.Tracer, i io.Closer) {
	sender, err := jaeger.NewUDPTransport(senderAddr, 0)
	if err != nil {
		panic(err)
	}
	reportOpt := jaeger.ReporterOptions.BufferFlushInterval(1 * time.Second)
	reporter := jaeger.NewRemoteReporter(
		sender,
		reportOpt)
	serviceName := "t-mk-client-test"
	tracer, closer := jaeger.NewTracer(
		serviceName,
		jaeger.NewConstSampler(true),
		reporter,
		jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
	)
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
