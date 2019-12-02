package ginopentrace

import (
	"math/rand"
	"net/http"
	"time"

	// "github.com/bilibili/kratos/pkg/net/metadata"
	"t-mk-opentrace/ext/log-driver/log"

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

// TracerWrapper tracer 中间件
func TracerWrapper(c *gin.Context) {
	// 提取
	spanCtx, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(c.Request.Header),
	)
	// child Span
	sp := opentracing.GlobalTracer().StartSpan(
		c.Request.URL.Path,
		opentracing.ChildOf(spanCtx),
		opentracing.Tags{"hostName": c.Request.Host},
	)

	defer sp.Finish()
	// 注入
	if err := opentracing.GlobalTracer().Inject(
		sp.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(c.Request.Header)); err != nil {
		log.Warn(err)
	}

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
	sp.LogKV("Request:", c.Request)
	// 异常取样
	if statusCode >= http.StatusInternalServerError {
		ext.Error.Set(sp, true)
	} else if rand.Intn(100) > sf {
		ext.SamplingPriority.Set(sp, 0)
	}
}
