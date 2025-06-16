# Contexts in Go

The `context` package, introduced in Go 1.7, is a standard library feature that is indispensable for writing robust, production-grade concurrent applications, especially servers and microservices.

A `Context` is a standard interface that carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between goroutines. Its primary purpose is to manage the lifecycle of a request or a unit of work, allowing for graceful shutdowns, timeouts, and the prevention of resource leaks.

---

### Basic Concepts

#### Context Creation: The Roots

Every context tree starts from a base context. There are two primary functions for this:

- **`context.Background()`**: This is the most common starting point. It returns an empty, non-cancellable context that is never done. It is meant to be the root of all context chains, typically used in `main()`, `init()`, and at the top level of a request handler.

- **`context.TODO()`**: This also returns an empty context, functionally identical to `Background()`. By convention, you use `context.TODO()` as a **placeholder** when you are unsure which context to use or when a function's signature has been updated to accept a context but the calling code hasn't been updated yet. It signals to static analysis tools and other developers that the context is not properly wired up and needs attention. **In production code, you should aim to replace all `context.TODO()` calls with a more appropriate context.**

#### Context Hierarchy: Deriving New Contexts

You almost never use `Background()` or `TODO()` directly. Instead, you derive new, "child" contexts from a parent context. This creates a tree of contexts where cancelling a parent also cancels all of its children.

1.  **`context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)`**

    - Creates a child context that can be explicitly cancelled.
    - It returns a `CancelFunc`. Calling this function cancels the context and all other contexts derived from it.
    - **Crucially, you must always call the `cancel` function to release resources associated with the context, even if the operation completes successfully. The common practice is to use `defer cancel()`.**

    ```go
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // Ensures cleanup happens
    ```

2.  **`context.WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)`**

    - A convenience function that creates a context that will be automatically cancelled after a specified duration.
    - It's equivalent to calling `context.WithDeadline()` with `time.Now().Add(timeout)`.
    - It also returns a `cancel` function that should be called to clean up resources early if the operation finishes before the timeout.

    ```go
    // This context will be cancelled after 2 seconds.
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    ```

3.  **`context.WithDeadline(parent Context, d time.Time) (Context, CancelFunc)`**

    - Creates a context that will be automatically cancelled at a specific point in time (`d`).
    - Useful when propagating a deadline from an incoming request.

    ```go
    deadline := time.Now().Add(5 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()
    ```

4.  **`context.WithValue(parent Context, key, val interface{}) Context`**

    - Creates a context that carries a request-scoped key-value pair.
    - This is for passing data that is not part of a function's explicit parameters, such as a request ID or a trace token.
    - **This should be used sparingly.** Misusing it can lead to unclear APIs. The key should be a custom, unexported type to avoid collisions.

    ```go
    type key int
    const requestIDKey key = 0
    ctx := context.WithValue(context.Background(), requestIDKey, "trace-12345")
    ```

---

### Practical Usage

#### Context Cancellation

A function that performs a long-running task should accept a context and use a `select` statement to listen for cancellation on the context's `Done()` channel.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // This channel is closed when the context is cancelled.
			fmt.Println("Worker: Cancellation signal received. Shutting down.")
			return // Exit the goroutine.
		default:
			fmt.Println("Worker: Doing some work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println("Main: It's time to stop the worker.")
	cancel() // Signal the worker to stop.

	time.Sleep(1 * time.Second) // Give the worker time to print its shutdown message.
	fmt.Println("Main: Finished.")
}
```

#### Timeouts and Deadlines

This is identical to cancellation, but the `cancel()` signal is sent automatically by Go's runtime when the time is up.

```go
func operationWithTimeout(ctx context.Context) {
    // This operation takes 3 seconds, but the context has a 2-second timeout.
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Operation completed successfully.")
    case <-ctx.Done():
        fmt.Println("Operation timed out!")
        // ctx.Err() will return the reason for cancellation (e.g., context.DeadlineExceeded)
        fmt.Println("Error:", ctx.Err())
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Always call cancel, even on timeout contexts.

    operationWithTimeout(ctx)
}
```

#### Context Values

This is how you pass and retrieve request-scoped data.

```go
type key string
const userIDKey key = "userID"

func processRequest(ctx context.Context) {
    // Retrieve the value. Type assertion is required.
    userID, ok := ctx.Value(userIDKey).(int)
    if !ok {
        fmt.Println("Could not find userID in context.")
        return
    }
    fmt.Println("Processing request for user ID:", userID)
}

func main() {
    // Store a value in the context.
    ctx := context.WithValue(context.Background(), userIDKey, 123)
    processRequest(ctx)
}
```

---

### Best Practices

1.  **Propagate Context Properly**: A context should be the **first argument** to a function, conventionally named `ctx`. `func doSomething(ctx context.Context, ...)`
2.  **Avoid Storing Contexts in Structs**: Storing a context in a struct makes it ambiguous when the context is active and can hide the flow of control. Pass it explicitly to each function that needs it.
3.  **Handle Context Cancellation**: Any function that accepts a context must be prepared to handle its cancellation. This means returning early and cleaning up any resources it has acquired.
4.  **Handle Context Values Carefully**:
    - Use `WithValue` only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
    - Use a custom, unexported type for context keys to prevent collisions between packages.
5.  **Avoid Creating Contexts in Loops (A Common Pitfall)**: Be careful when deriving contexts inside a loop. Calling `defer cancel()` is wrong because the `defer` will only execute when the _function_ returns, not when the loop iteration ends, leading to a resource leak.

    **WRONG:**

    ```go
    for _, item := range items {
        ctx, cancel := context.WithTimeout(parentCtx, 1*time.Second)
        // This defer will not run until the surrounding function exits!
        // This leaks timers and goroutines associated with the context.
        defer cancel()
        go doWork(ctx, item)
    }
    ```

    **RIGHT:**

    ```go
    for _, item := range items {
        ctx, cancel := context.WithTimeout(parentCtx, 1*time.Second)
        // Call cancel at the end of the iteration or in the worker.
        go func() {
            doWork(ctx, item)
            cancel() // Cancel when work is done to free resources.
        }()
    }
    ```

---

### Common Pitfalls

1.  **Ignoring Context Cancellation (Goroutine Leaks)**: If a function receives a context but never checks `ctx.Done()`, it may continue running long after the caller has given up, wasting CPU and memory. This is a classic source of goroutine leaks.
2.  **Misusing Context Values**: Using `WithValue` to pass required parameters to a function makes your API implicit and hard to understand. If a function needs a `userID`, it should be an explicit parameter: `func processUser(userID int)`. Use `WithValue` for "out-of-band" data only.
3.  **A `nil` Context**: Passing a `nil` context to functions that expect a `context.Context` will cause a panic. Always start with `context.Background()` or `context.TODO()` if you don't have another context to pass.
