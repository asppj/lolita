package middleware

import (
	"context"

	"github.com/asppj/t-mk-opentrace/ext/log-driver/log"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// TextMapReader read
type TextMapReader struct {
	metadata.MD
}

// ForeachKey 读取metadata中的span信息
func (t TextMapReader) ForeachKey(handler func(key, val string) error) error { // 不能是指针
	for key, val := range t.MD {
		for _, v := range val {
			if err := handler(key, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// GRPCServerInterceptor 拦截器
func GRPCServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 从context中获取metadata。md.(type) == map[string][]string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		// 如果对metadata进行修改，那么需要用拷贝的副本进行修改。（FromIncomingContext的注释）
		md = md.Copy()
	}
	carrier := TextMapReader{md}
	tracer := opentracing.GlobalTracer()
	spanContext, e := tracer.Extract(opentracing.TextMap, carrier)
	if e != nil {
		log.Debug("GRPCServerInterceptor-context 未提取到span", e)
	}
	span := tracer.StartSpan(info.FullMethod, opentracing.ChildOf(spanContext))
	defer span.Finish()
	ctx = opentracing.ContextWithSpan(ctx, span)

	return handler(ctx, req)
}
