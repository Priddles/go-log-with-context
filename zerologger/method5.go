package zerologger

import (
	"context"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// contextLogger returns a func that will log a [LogContext] embedded in a [context.Context].
func contextLogger(ctx context.Context) func(e *zerolog.Event) {
	return func(e *zerolog.Event) {
		logContext, ok := GetLogContext(ctx)
		if !ok {
			return
		}

		e.Str("requestId", logContext.RequestID)
	}
}

// Handler_5 handles logging by:
//
//  1. Embedding data into the given context.
//  2. Taking that context down the rabbit hole.
//  3. Passing a thunk to the logger that will log the data in the context.
func Handler_5(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logContext := LogContext{RequestID: request.RequestContext.RequestID}
	ctx = WithLogContext(ctx, logContext)

	log.Info().
		Func(contextLogger(ctx)).
		Str("uniqueCode", "12345678").
		Str("httpMethod", request.HTTPMethod).
		Str("httpPath", request.Path).
		Msg("handler_5")

	err := rabbithole_5(ctx, request.Body)
	if err != nil {
		log.Error().Err(err).
			Func(contextLogger(ctx)).
			Str("uniqueCode", "12345678").
			Msg("handler_5 error")

		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func rabbithole_5(ctx context.Context, greeting string) error {
	log.Debug().
		Func(contextLogger(ctx)).
		Str("uniqueCode", "12345678").
		Msg("rabbithole_5")

	return errors.New("not implemented")
}
