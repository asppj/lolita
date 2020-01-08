package middleware

import (
	got "github.com/asppj/t-go-opentrace/ext/gin-opentrace"

	"github.com/gin-gonic/gin"
)

// OpenTraceMiddleware Extract header
// Extract
func OpenTraceMiddleware() gin.HandlerFunc {
	// return got.Middleware(trace)
	return got.TracerWrapper
}
