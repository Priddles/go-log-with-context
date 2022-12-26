package zerologger

import (
	"context"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Handler_1 handles logging by:
//
//  1. Creating a child logger of the global logger.
//  2. Embedding the new logger into the given context.
//  3. Taking that context down the rabbit hole.
//  4. Retrieving the new logger from the context.
//
// The new logger should be retrieved from the context by using [zerolog.Ctx]
// (or [log.Ctx] which is an alias).
func Handler_1(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := log.With().Str("requestId", request.RequestContext.RequestID).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().
		Str("uniqueCode", "12345678").
		Str("httpMethod", request.HTTPMethod).
		Str("httpPath", request.Path).
		Msg("handler_1")

	err := rabbithole_1(ctx, request.Body)
	if err != nil {
		logger.Error().Err(err).
			Str("uniqueCode", "12345678").
			Msg("handler_1 error")

		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func rabbithole_1(ctx context.Context, greeting string) error {
	// N.B. This is the same as calling [log.Ctx].
	logger := zerolog.Ctx(ctx)

	logger.Debug().
		Str("uniqueCode", "12345678").
		Msg("rabbithole_1")

	return errors.New("not implemented")
}
