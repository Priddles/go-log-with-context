package zerologger_test

import (
	"context"
	"testing"

	"github.com/Priddles/go-log-with-context/log/zerologger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
)

func Test_Handler_3(t *testing.T) {
	globalLogger := log.Logger
	t.Cleanup(func() {
		log.Logger = globalLogger
	})

	_, err := zerologger.Handler_3(context.Background(), events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/hello/world",
		RequestContext: events.APIGatewayProxyRequestContext{
			RequestID: "5ca1ab1e",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
