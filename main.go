package main

import (
	"context"
	"fmt"

	"github.com/Priddles/go-log-with-context/log/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaInput struct {
	RequestID  string
	XRay       string
	ResourceID string
}

func handler(ctx context.Context, input LambdaInput) (*events.APIGatewayProxyResponse, error) {
	ctx = log.MustContextWithTracing(ctx, &log.TracingContext{
		// EventID: input.EventID,
		// ExecutionID: input.ExecutionID,
		RequestID: input.RequestID,
		TraceID:   input.XRay,
	})

	log.Log(ctx, log.ERROR, "serving API request", map[string]any{
		"resourceId": input.ResourceID,
	})

	return nil, fmt.Errorf("always fails =P")
}

func main() {
	lambda.Start(handler)
}
