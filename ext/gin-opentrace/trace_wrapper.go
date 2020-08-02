package ginopentrace

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	// "github.com/bilibili/kratos/pkg/net/metadata"
	"github.com/asppj/lolita/ext/log-driver/log"

	"github.com/gin-gonic/gin"

	// "github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// const contextTracerKey = "Tracer-context"

// sf sampling frequency
var sf = 100

func init() {
	rand.Seed(time.Now().Unix())
}

// SetSamplingFrequency 设置采样频率
// 0 <= n <= 100
func SetSamplingFrequency(n int) {
	sf = n
}

var newSpanName = "ginMiddleware_"

// TracerWrapper tracer 中间件
func TracerWrapper(c *gin.Context) {
	// 提取
	spanCtx, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(c.Request.Header),
	)
	if spanCtx == nil {
		newSpan := opentracing.StartSpan(newSpanName + c.Request.Method)
		newSpan.SetTag("hasParent", false)
		spanCtx = newSpan.Context()
	}
	// child Span
	sp := opentracing.GlobalTracer().StartSpan(
		c.Request.Method+" "+c.Request.URL.Path,
		opentracing.ChildOf(spanCtx),
	)
	defer sp.Finish()
	// 注入
	if err := opentracing.GlobalTracer().Inject(
		sp.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(c.Request.Header)); err != nil {
		log.Warn(err)
	}
	c.Request = c.Request.WithContext(
		opentracing.ContextWithSpan(c.Request.Context(), sp))
	c.Next()
	// 收集
	statusCode := c.Writer.Status()
	// Tag
	ext.PeerHostname.Set(sp, c.Request.Host)
	ext.PeerAddress.Set(sp, c.Request.RemoteAddr)
	ext.PeerService.Set(sp, c.ClientIP())
	ext.HTTPStatusCode.Set(sp, uint16(statusCode))
	ext.HTTPMethod.Set(sp, c.Request.Method)
	ext.HTTPUrl.Set(sp, c.Request.URL.Path)
	// log
	sp.LogKV("Header:", c.Request.Header)
	// 异常取样
	if statusCode >= http.StatusInternalServerError {
		ext.Error.Set(sp, true)
		sp.LogKV("RequestURI:", c.Request.RequestURI)
		sp.LogKV("RemoteAddr:", c.Request.RemoteAddr)
	} else if rand.Intn(100) > sf {
		ext.SamplingPriority.Set(sp, 0)
	}
}

// GetCtxFromGinContext extract span from ctx
// http 请求中调用grpc服务时使用
func GetCtxFromGinContext(ctx *gin.Context) context.Context {
	spanCtx, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(ctx.Request.Header),
	)
	if err != nil {
		return context.Background()
	}
	sp := opentracing.GlobalTracer().StartSpan(
		"HTTPToGRPC",
		opentracing.FollowsFrom(spanCtx),
		opentracing.Tags{"hostName": ctx.Request.Host},
	)
	defer sp.Finish()
	// 注入
	return opentracing.ContextWithSpan(context.Background(), sp)
}

// SpanTransferFromContextToHeader extract span from ctx
func SpanTransferFromContextToHeader(ctx context.Context) context.Context {
	var parentCtx opentracing.SpanContext
	if parent := opentracing.SpanFromContext(ctx); parent != nil {
		parentCtx = parent.Context()
	}
	sp := opentracing.StartSpan(
		"GRPCToHTTP",
		opentracing.FollowsFrom(parentCtx),
		ext.SpanKindRPCClient,
	)
	defer sp.Finish()
	return opentracing.ContextWithSpan(context.Background(), sp)
}
