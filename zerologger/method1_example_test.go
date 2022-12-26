package zerologger_test

import (
	"context"
	"testing"

	"github.com/Priddles/go-log-with-context/log/zerologger"
	"github.com/aws/aws-lambda-go/events"
)

func Test_Handler_1(t *testing.T) {
	_, err := zerologger.Handler_1(context.Background(), events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/hello/world",
		RequestContext: events.APIGatewayProxyRequestContext{
			RequestID: "deadbeef",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
