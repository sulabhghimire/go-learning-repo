# Go `sync.WaitGroup`: A Guide to Synchronizing Goroutines

A `sync.WaitGroup` is a synchronization primitive provided by Go's `sync` package. Its purpose is to block a goroutine until a collection of other goroutines have completed their execution. It is one of the most common and essential tools for managing concurrency in Go.

## Why Use WaitGroups?

- **Synchronization:** The primary use case is to synchronize the execution flow. It allows a parent goroutine (often `main`) to wait for multiple child goroutines to finish their work before proceeding. This prevents the main program from exiting prematurely.
- **Coordination:** WaitGroups help coordinate the completion of concurrent tasks. They act as a "barrier," ensuring that a set of operations are all complete before the program moves on to the next stage, which might depend on the results or side effects of those operations.
- **Resource Management:** They are crucial for safely managing resources. For example, you can ensure that all goroutines using a database connection have finished their work before the connection is closed.

## Core Methods

To use a WaitGroup, you create an instance and use its three core methods.

```go
import "sync"

var wg sync.WaitGroup
```

### 1. `Add(delta int)`

This method increments the WaitGroup's internal counter by the `delta` value. This number represents the number of goroutines the `WaitGroup` should wait for.

- **When to call:** Call `Add` _before_ you launch the goroutine you intend to wait for.

```go
numWorkers := 5
wg.Add(numWorkers) // Set the counter to 5
```

### 2. `Done()`

This method decrements the WaitGroup's counter by one.

- **When to call:** Each goroutine must call `Done()` when it has finished its work. It's a signal from the worker that it has completed its task.

```go
// Inside a worker goroutine
fmt.Println("Work finished.")
wg.Done() // Signal completion
```

### 3. `Wait()`

This method blocks the goroutine that calls it until the WaitGroup's internal counter becomes zero.

- **When to call:** Call `Wait()` in the goroutine that needs to wait for the workers to finish (usually the `main` goroutine).

```go
// In the main goroutine, after launching all workers
wg.Wait() // Block until the counter is zero
fmt.Println("All workers have completed.")
```

### Complete Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// worker function takes an id and a pointer to a WaitGroup
func worker(id int, wg *sync.WaitGroup) {
	// Use defer to ensure Done is called when the function exits.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	// Simulate some work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// 1. Create a new WaitGroup instance.
	var wg sync.WaitGroup

	numWorkers := 3

	// 2. Call Add to set the number of goroutines to wait for.
	// This should be done *before* launching the goroutines.
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// Launch a goroutine for each worker.
		// Pass a pointer to the WaitGroup.
		go worker(i, &wg)
	}

	fmt.Println("Main: Waiting for workers to finish...")
	// 3. Call Wait to block until all workers have called Done.
	wg.Wait()

	fmt.Println("Main: All workers have completed.")
}
```

## Best Practices

- **Use `defer` to Call `Done`:** Always call `wg.Done()` inside a `defer` statement at the top of your goroutine's function. This guarantees that the counter is decremented, even if the goroutine panics.
  ```go
  func worker(wg *sync.WaitGroup) {
      defer wg.Done()
      // ... rest of the work ...
      // If a panic occurs here, Done() is still called.
  }
  ```
- **Call `Add` Before Launching the Goroutine:** Call `Add` in the parent goroutine _before_ the `go` statement. If you call `Add` inside the child goroutine, you create a race condition where `Wait()` might execute before `Add` is ever called, causing the program to continue without actually waiting.

  ```go
  // CORRECT
  for i := 0; i < 5; i++ {
      wg.Add(1)
      go doWork(i, &wg)
  }

  // INCORRECT (RACE CONDITION)
  for i := 0; i < 5; i++ {
      go func(i int) {
          wg.Add(1) // Too late! Wait() might run before this.
          doWork(i, &wg)
      }(i)
  }
  ```

- **Use `defer` for Unlocking Mutexes:** If your goroutines use shared resources protected by a mutex, use `defer` to unlock it. This ensures resources are not left in a locked state, which could prevent other goroutines from finishing and calling `Done`.
  ```go
  func criticalWork(wg *sync.WaitGroup, mu *sync.Mutex) {
      defer wg.Done()
      mu.Lock()
      defer mu.Unlock() // Guarantees the mutex is unlocked.
      // ... critical section ...
  }
  ```

## Common Pitfalls

- **Mismatch Between `Add` and `Done` Calls:** The number of `Done` calls must exactly match the number you added to the counter.

  - **`Add` > `Done`:** If you add more to the counter than you have `Done` calls, the counter will never reach zero. `wg.Wait()` will block forever, causing a **deadlock**.
  - **`Done` > `Add`:** If you call `Done` more times than the current counter value, the program will **panic** with a "negative WaitGroup counter" error.

- **Creating Deadlocks:** The most common cause of a `WaitGroup` deadlock is forgetting to call `Done` in all possible execution paths of a goroutine (e.g., inside an `if/else` block where one path misses the `Done` call). Using `defer wg.Done()` is the best way to prevent this.
