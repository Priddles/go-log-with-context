package zerologger

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Handler_3 handles logging by:
//
//  1. Overwriting the global logger with a new logger.
//
// As simple as it is, this method has limited use cases.
func Handler_3(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Str("requestId", request.RequestContext.RequestID).Logger()

	log.Info().
		Str("uniqueCode", "12345678").
		Str("httpMethod", request.HTTPMethod).
		Str("httpPath", request.Path).
		Msg("handler_3")

	err := rabbithole_3(request.Body)
	if err != nil {
		log.Error().Err(err).
			Str("uniqueCode", "12345678").
			Msg("handler_3 error")

		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func rabbithole_3(greeting string) error {
	log.Debug().
		Str("uniqueCode", "12345678").
		Msg("rabbithole_3")

	return errors.New("not implemented")
}
