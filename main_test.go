package main

import (
	"context"
	"testing"
	"time"
)

func Test_healthCheckIngest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := healthCheckIngest(ctx, "localhost", 1491, "SecretPassword")
	if err != nil {
		t.Error(err)
	}
}
