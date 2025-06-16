# Go Tickers: Scheduling Periodic Tasks

A Ticker in Go is a concurrency primitive that "ticks" at a regular interval, sending a signal (the current time) on a channel. Tickers are the idiomatic way to perform periodic or recurring tasks on a consistent schedule.

They are ideal for scenarios like:

- **Polling:** Periodically checking the status of a system or resource.
- **Regular Updates:** Refreshing a data cache or UI component every few seconds.
- **Periodic Logging:** Reporting application metrics or health checks at a fixed interval.

### Creating and Using a Ticker

You create a ticker using `time.NewTicker()`, which takes a `time.Duration` specifying the interval between ticks.

The returned `*time.Ticker` object has a public channel field, `C` (`<-chan time.Time`), which receives the time at each tick.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new ticker that ticks every 1 second.
	ticker := time.NewTicker(1 * time.Second)
	// When the function exits, stop the ticker to release resources.
	defer ticker.Stop()

	// A channel to signal completion.
	done := make(chan bool)

	// A separate goroutine to stop the process after 5 seconds.
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	fmt.Println("Ticker started. Waiting for ticks...")

	// The main loop waits for either a tick or a done signal.
	for {
		select {
		case <-done:
			fmt.Println("Done signal received. Exiting.")
			return
		case t := <-ticker.C:
			// This case is executed on every tick.
			fmt.Println("Tick at", t.Format("15:04:05"))
		}
	}
}
```

## Why Use Tickers?

While you could implement similar behavior with a `for` loop and `time.Sleep()`, tickers offer significant advantages in concurrent programs.

- **Consistency:** Tickers are managed by the Go runtime and are designed to maintain a steady rhythm. The runtime attempts to send ticks at the specified interval, adjusting for delays. A simple `time.Sleep(interval)` inside a loop does not account for the time the loop body takes to execute, leading to drift over time.
- **Simplicity and Idiomatic Code:** The `select` pattern with a ticker channel is the standard, readable, and idiomatic way to handle periodic tasks in Go. It integrates seamlessly with other channel operations, such as a shutdown signal.

## Best Practices for Using Tickers

To use tickers safely and efficiently, follow these best practices.

### 1. Stop Tickers When No Longer Needed

**This is the most important rule.** A ticker that is not stopped will leak resources. The Go runtime keeps the ticker active in the background, consuming memory and CPU, even if no code is listening on its channel.

Always call `ticker.Stop()` to release these resources. The most reliable way to do this is with a `defer` statement.

**Correct: Using `defer` to guarantee cleanup**

```go
func startPeriodicTask() {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop() // Guarantees the ticker is stopped when the function returns.

    for range ticker.C {
        fmt.Println("Performing periodic task...")
    }
}
```

### 2. Avoid Blocking Operations in the Ticker's `select` Case

If the task you perform on each tick is long-running, it can block the `select` loop from receiving the next tick. This can cause ticks to be missed or delayed, disrupting the schedule.

If a task might block or take a significant amount of time, run it in its own goroutine.

**Incorrect: Blocking operation**

```go
// AVOID THIS PATTERN
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

for range ticker.C {
    // If this takes > 1 second, you will miss the next tick.
    time.Sleep(2 * time.Second)
    fmt.Println("This task is too slow!")
}
```

**Correct: Non-blocking operation**

```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

for range ticker.C {
    // Launch the long-running task in its own goroutine.
    go func() {
        fmt.Println("Starting a long task...")
        time.Sleep(2 * time.Second) // Simulate work.
        fmt.Println("Long task finished.")
    }()
}
// Required to see the output from the goroutines in this example.
time.Sleep(5 * time.Second)
```

### 3. Handle Ticker Stopping Gracefully

In long-running applications, you need a way to tell a ticker loop to stop. The standard Go pattern is to use a dedicated `done` channel. The `select` statement listens on both the ticker channel and the `done` channel, allowing for a clean shutdown.

This pattern is demonstrated in the first example and is crucial for building robust, manageable services.

```go
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done: // Listen for the shutdown signal.
			fmt.Println("Worker: stopping.")
			return
		case t := <-ticker.C:
			fmt.Println("Worker: processing tick at", t.Format("15:04:05.000"))
		}
	}
}

func main() {
	done := make(chan bool)
	go worker(done)

	// Let the worker run for a few seconds.
	time.Sleep(2 * time.Second)

	// Send the shutdown signal.
	close(done) // Closing the channel broadcasts the signal to all listeners.

	// Give the worker a moment to shut down gracefully.
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main: finished.")
}
```
