package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultTimeOut = 5

// WithTimeOut ctx超时时间
func WithTimeOut() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), defaultTimeOut*time.Second)
		defer func() {
			if c.Err() == context.DeadlineExceeded {
				ctx.Writer.WriteHeader(http.StatusGatewayTimeout)
				ctx.Abort()
			}
			cancel()
		}()
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}
