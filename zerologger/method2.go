package zerologger

import (
	"context"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
)

// Handler_2 handles logging in a manner almost identical to [Handler_1],
// but has some caveats that make it worth listing separately.
//
// It handles logging by:
//
//  1. Retrieving a logger from the given context.
//  2. Creating a child logger from the context's logger.
//  3. Overwriting the context's logger with the new logger.
//  4. Taking that context down the rabbit hole.
//  5. Retrieving the new logger from the context.
//
// WARNING - This method relies on the assumption that either:
//
//  1. The given context has an embedded logger, or
//  2. The global variable [zerolog.DefaultContextLogger] is not nil.
//
// If both of these assumptions are false, then zerolog will silently return a DISABLED logger.
//
// The new logger should be retrieved from the context by using [zerolog.Ctx]
// (or [log.Ctx] which is an alias).
func Handler_2(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := zerolog.Ctx(ctx).With().Str("requestId", request.RequestContext.RequestID).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().
		Str("uniqueCode", "12345678").
		Str("httpMethod", request.HTTPMethod).
		Str("httpPath", request.Path).
		Msg("handler_2")

	err := rabbithole_2(ctx, request.Body)
	if err != nil {
		logger.Error().Err(err).
			Str("uniqueCode", "12345678").
			Msg("handler_2 error")

		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func rabbithole_2(ctx context.Context, greeting string) error {
	logger := zerolog.Ctx(ctx)

	logger.Debug().
		Str("uniqueCode", "12345678").
		Msg("rabbithole_2")

	return errors.New("not implemented")
}
