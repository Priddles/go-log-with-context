package log

import (
	"context"
	"errors"
)

type logContextKey string

// Context keys are not exported.
const (
	// tracingContextKey is the key for TracingContext in Contexts.
	tracingContextKey logContextKey = "tracing"
)

// ErrContextNotEmpty is returned when a context is already carrying information.
var ErrContextNotEmpty = errors.New("log: context not empty")

// ContextWithTracing returns a new Context that carries tracing information.
//
// If the context is already carrying tracing information, it will return
// ErrContextNotEmpty.
func ContextWithTracing(ctx context.Context, tracing *TracingContext) (context.Context, error) {
	if _, exists := GetTracing(ctx); exists {
		return ctx, ErrContextNotEmpty
	}

	return context.WithValue(ctx, tracingContextKey, tracing), nil
}

// Must is a helper function to ensure there was no error when calling the
// ContextWithTracing function.
func Must(ctx context.Context, err error) context.Context {
	if err != nil {
		panic(err)
	}

	return ctx
}

// GetTracing returns the tracing information carried by ctx, if any.
func GetTracing(ctx context.Context) (*TracingContext, bool) {
	tracing, ok := ctx.Value(tracingContextKey).(*TracingContext)
	return tracing, ok
}
