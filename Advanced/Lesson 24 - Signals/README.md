# A Guide to Handling OS Signals in Go

In Unix-like operating systems, **signals** are a fundamental form of inter-process communication (IPC). They are asynchronous notifications sent to a process to inform it of an event, such as a user request to interrupt or terminate.

Properly handling signals is the key to creating applications that can shut down gracefully, clean up resources, and interact predictably with their environment.

#### Why Use Signals?

- **Graceful Shutdowns:** This is the most common use case. When a user presses `Ctrl+C` (`SIGINT`) or a process manager (like Kubernetes or `systemd`) sends a termination signal (`SIGTERM`), your application can catch it, finish any in-progress work, save state, close database connections, and exit cleanly.
- **Resource Cleanup:** A graceful shutdown ensures that temporary files are deleted, network connections are closed properly, and locks are released, preventing resource leaks or corruption.
- **Configuration Reloading:** Some applications use the `SIGHUP` (hangup) signal as a trigger to reload their configuration files without requiring a full restart.

### How to Use Signals in Go: The `os/signal` Package

Go's standard library provides the `os/signal` package, which allows a program to listen for and respond to incoming OS signals. The core mechanism involves a Go channel.

The main function is `signal.Notify(c chan<- os.Signal, sig ...os.Signal)`.

- It directs the Go runtime to relay incoming signals of the specified types (`sig...`) to the channel `c`.
- If you don't specify any signals, all incoming signals will be relayed.

#### Common Signals in Unix-like Systems

| Signal        | Default Action | Common Use Case                                                            | Can be Handled? |
| :------------ | :------------- | :------------------------------------------------------------------------- | :-------------- |
| **`SIGINT`**  | Terminate      | Sent when you press `Ctrl+C`. "Interrupt, please stop gracefully."         | **Yes**         |
| **`SIGTERM`** | Terminate      | The standard "polite" request to terminate. Used by `kill <PID>`.          | **Yes**         |
| **`SIGHUP`**  | Terminate      | "Hang up." Often used to trigger configuration reloads.                    | **Yes**         |
| **`SIGKILL`** | Terminate      | The "un-ignorable" kill signal. The OS terminates the process immediately. | **No**          |
| **`SIGSTOP`** | Stop/Pause     | Pauses a process. Cannot be handled.                                       | **No**          |
| **`SIGCONT`** | Continue       | Resumes a process that was paused with `SIGSTOP`.                          | **Yes**         |

### Example 1: Basic Signal Handling (`SIGINT`)

Let's create a simple program that waits for a `Ctrl+C` and then prints a message before exiting.

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. Create a channel to receive OS signals.
	// A buffered channel of size 1 is recommended so the notifier
	// doesn't block if the program isn't ready to receive the signal.
	sigs := make(chan os.Signal, 1)

	// 2. Register the channel to receive notifications for specific signals.
	// We'll listen for SIGINT (Ctrl+C) and SIGTERM (standard termination).
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Program running. Press Ctrl+C to exit.")

	// 3. This is a blocking operation. The program will wait here until
	// a signal is received on the channel.
	sig := <-sigs
	fmt.Println() // For a newline after ^C
	fmt.Printf("Signal received: %s\n", sig)
	fmt.Println("Starting graceful shutdown...")

	// 4. Perform cleanup tasks here.
	time.Sleep(2 * time.Second) // Simulate cleanup work

	fmt.Println("Cleanup complete. Exiting.")
}
```

**How to run this:**

1.  Save the code as `main.go` and run `go run main.go`.
2.  Press `Ctrl+C` in your terminal.
3.  You will see the shutdown messages printed before the program exits.

### Example 2: Graceful HTTP Server Shutdown

A more practical example is a web server that needs to finish handling active requests before shutting down.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Set up the signal handling channel
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Create a server and a handler
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}

	// Start the server in a goroutine so it doesn't block
	go func() {
		log.Println("Server starting on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v\n", err)
		}
	}()

	// Block until a signal is received
	sig := <-sigs
	log.Printf("Signal received: %s. Initiating graceful shutdown...", sig)

	// Create a context with a timeout to allow for graceful shutdown.
	// This gives active requests 30 seconds to complete.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown() gracefully shuts down the server without interrupting any
	// active connections. It works by first closing all open listeners,
	// then closing all idle connections, and then waiting indefinitely for
	// connections to return to idle and then shut down.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}

// A simple handler that simulates work
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request...")
	time.Sleep(5 * time.Second) // Simulate a long-running task
	fmt.Fprintln(w, "Request handled successfully!")
	log.Println("Finished request.")
}
```

**How to test this:**

1.  Run the program: `go run main.go`.
2.  Open a web browser or use `curl http://localhost:8080`.
3.  While `curl` is waiting, go to the terminal running the server and press `Ctrl+C`.
4.  You will see the server print the "Initiating graceful shutdown..." message, but it will **wait** for the 5-second request to complete before finally exiting.

### Sending Signals with the `kill` Command

You can manually send signals to your running Go application from the terminal.

1.  **Find the Process ID (PID):**

    ```bash
    # Find the PID of your running program
    $ pgrep your_program_name
    12345
    ```

2.  **Send a Signal:**

    ```bash
    # Send SIGTERM (the default, graceful)
    $ kill 12345

    # Send SIGINT (same as Ctrl+C)
    $ kill -s SIGINT 12345
    # or
    $ kill -2 12345

    # Send SIGKILL (force kill, cannot be caught)
    $ kill -9 12345
    ```

### Debugging and Troubleshooting

- **Debugging Signal Handling:** The best way to debug is with logging. Add `log.Println()` statements at each stage: when the listener is set up, when a signal is received, when cleanup starts, and when it finishes. This allows you to trace the execution flow.

- **Common Issues:**
  - **Signal Lost:** A signal can be "lost" if it's sent to the process _before_ `signal.Notify()` is called. To mitigate this, set up your signal handling channel as one of the very first things in your `main()` function. Using a **buffered channel** (`make(chan os.Signal, 1)`) is also critical, as it allows the runtime to send a signal into the channel even if your program isn't immediately ready to receive it.
  - **Deadlocks During Shutdown:** A common bug is for the shutdown logic to wait for a goroutine to finish, but that goroutine is itself blocked, waiting for something that the shutdown logic holds a lock on. Using `context.WithTimeout` for shutdowns (as in the HTTP server example) provides a crucial safety net, ensuring your program will eventually exit even if the graceful shutdown fails.
