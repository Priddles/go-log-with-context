package zerologger_test

import (
	"context"
	"testing"

	"github.com/Priddles/go-log-with-context/log/zerologger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Test_Handler_2_With_Context_Value(t *testing.T) {
	ctx := log.Logger.WithContext(context.Background())

	_, err := zerologger.Handler_2(ctx, events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/hello/world",
		RequestContext: events.APIGatewayProxyRequestContext{
			RequestID: "c0ffee",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Handler_2_With_Global_Default(t *testing.T) {
	globalDefault := zerolog.DefaultContextLogger
	t.Cleanup(func() {
		zerolog.DefaultContextLogger = globalDefault
	})

	zerolog.DefaultContextLogger = &log.Logger

	_, err := zerologger.Handler_2(context.Background(), events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/hello/world",
		RequestContext: events.APIGatewayProxyRequestContext{
			RequestID: "cafe",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
