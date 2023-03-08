package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/expectedsh/go-sonic/sonic"
)

func main() {
	host := os.Getenv("HEALTH_CHECK_HOST")
	port := os.Getenv("PORT")
	pass := os.Getenv("PASSWORD")

	portI, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := healthCheckIngest(ctx, host, portI, pass); err != nil {
		log.Fatal(err)
	}
}

func healthCheckIngest(ctx context.Context, host string, port int, pass string) error {
	doneChan := make(chan error, 1)

	go func() {
		defer close(doneChan)

		ingester, err := sonic.NewIngester(host, port, pass)
		if err != nil {
			doneChan <- err
			return
		}

		doneChan <- ingester.Ping()
	}()

	select {
	case err := <-doneChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
