package zerologger

import (
	"context"
)

type key struct{}

var logContextKey key

type LogContext struct {
	RequestID string
}

func WithLogContext(ctx context.Context, val LogContext) context.Context {
	return context.WithValue(ctx, logContextKey, &val)
}

func GetLogContext(ctx context.Context) (*LogContext, bool) {
	val, ok := ctx.Value(logContextKey).(*LogContext)
	return val, ok
}
