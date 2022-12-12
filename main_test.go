package main

import (
	"context"
	"testing"
)

func Test_handler(t *testing.T) {
	handler(context.Background(), LambdaInput{
		RequestID:  "me-first",
		XRay:       "superpowers",
		ResourceID: "that-one",
	})
}
