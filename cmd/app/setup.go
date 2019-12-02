package main

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

// Host h
var Host = "localhost"

// Port p
var Port = "6831"
var senderAddr = Host + ":" + Port

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
	serviceName := "t-mk-openTrace"
	tracer, closer := jaeger.NewTracer(
		serviceName,
		jaeger.NewConstSampler(true),
		reporter,
		jaeger.TracerOptions.ZipkinSharedRPCSpan(true),
	)
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
