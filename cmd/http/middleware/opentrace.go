package middleware

import (
	got "t-mk-opentrace/ext/gin-opentrace"

	"github.com/gin-gonic/gin"
)

// OpenTraceMiddleware Extract header
// Extract
func OpenTraceMiddleware() gin.HandlerFunc {
	// return got.Middleware(trace)
	return got.TracerWrapper
}
