package zerologger

import (
	"context"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Compile-time checks.
var (
	_ zerolog.LogObjectMarshaler = LogContext{}
)

// Extend [LogContext] to implement [zerolog.LogObjectMarshaler]
func (ctx LogContext) MarshalZerologObject(e *zerolog.Event) {
	e.Str("requestId", ctx.RequestID)
}

func logWithContext(e *zerolog.Event, ctx context.Context) *zerolog.Event {
	logContext, ok := GetLogContext(ctx)
	if ok {
		e.EmbedObject(logContext)
	}

	return e
}

// Handler_4 handles logging by:
//
//  1. Creating data that implements [zerolog.LogObjectMarshaler].
//  2. Embedding that data into the given context.
//  3. Taking that context down the rabbit hole.
//  4. Retrieving the data from the context.
//
// The data can be retrieved using [context.Context.Value],
// and added to logs using [zerolog.Event.EmbedObject].
//
// A helper function [logWithContext] (which calls EmbedObject) is used to avoid panics when the
// context is missing the expected data.
// (The panic happens because a nil pointer is returned when the context is missing data.
// When the typed pointer is assigned to an interface, a basic nil-check is no longer sufficient,
// and zerolog tries to dereference it.)
func Handler_4(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logContext := LogContext{RequestID: request.RequestContext.RequestID}
	ctx = WithLogContext(ctx, logContext)

	log.Info().
		EmbedObject(logContext).
		Str("uniqueCode", "12345678").
		Str("httpMethod", request.HTTPMethod).
		Str("httpPath", request.Path).
		Msg("handler_4")

	err := rabbithole_4(ctx, request.Body)
	if err != nil {
		log.Error().Err(err).
			EmbedObject(logContext).
			Str("uniqueCode", "12345678").
			Msg("handler_4 error")

		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func rabbithole_4(ctx context.Context, greeting string) error {
	logWithContext(log.Debug(), ctx).
		Str("uniqueCode", "12345678").
		Msg("rabbithole_4")

	return errors.New("not implemented")
}
