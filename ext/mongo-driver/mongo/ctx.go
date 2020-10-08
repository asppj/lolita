package mongo

import (
	"context"
	"time"
)

// DefaultCtxWithOut 带超时
func DefaultCtxWithOut() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultQueryTimeout*time.Second)
}
