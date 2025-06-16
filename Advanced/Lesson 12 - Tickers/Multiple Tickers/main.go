package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func startTicker(ctx context.Context, name string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Simulate work
			println("Fetching data for", name, "at", time.Now().Format(time.RFC3339))
		case <-ctx.Done():
			println(name, "stopped")
			return
		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("\nInterrupt received. Stopping tickers...")
		cancel()
	}()

	go startTicker(ctx, "AAPL", 2*time.Second)
	go startTicker(ctx, "GOOG", 3*time.Second)
	go startTicker(ctx, "TSLA", 5*time.Second)

	// Keep main alive until context is done
	<-ctx.Done()
	time.Sleep(1 * time.Second) // Give goroutines time to log shutdown
	fmt.Println("All tickers stopped. Exiting.")

}
