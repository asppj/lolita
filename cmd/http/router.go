package http

import (
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.Engine) error {
	router.GET("", func(ctx *gin.Context) {
		tr := opentracing.GlobalTracer()
		carrier := opentracing.HTTPHeadersCarrier(ctx.Request.Header)
		span := opentracing.SpanFromContext(ctx.Request.Context())
		if span == nil {
			ctx.String(http.StatusInternalServerError, "Span is nil\n")
			// return
			span = opentracing.StartSpan("sp-lsp")
		}
		cs := opentracing.StartSpan("lsp-root-sp",
			opentracing.ChildOf(span.Context()))
		_ = tr.Inject(cs.Context(), "c-lsp-root-sp", carrier)
		defer cs.Finish()
		req, err := http.NewRequest("GET", "http://localhost:6006/user", nil)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error non-nil %v", err)
		}
		client := &http.Client{
			Timeout: 0,
		}
		_, err = client.Do(req)
		if err != nil {
			log.Println(err)
		}
		// carrier := opentracing.HTTPHeadersCarrier(req.Header)
		// Verify the context was populated as expected by the middleware
		// err = opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)
		ctx.String(http.StatusOK, "router oks")

	})
	router.GET("/user", func(ctx *gin.Context) {
		span := opentracing.SpanFromContext(ctx.Request.Context())
		if span == nil {
			ctx.String(http.StatusInternalServerError, "user Span is nil\n")
			return
		}
		ctx.String(http.StatusOK, "user oks")

	})
	return nil
}
