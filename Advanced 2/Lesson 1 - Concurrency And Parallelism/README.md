# Go Concurrency and Parallelism Explained

This document provides an introduction to the concepts of concurrency and parallelism, highlighting their differences, challenges, and a practical implementation in Go.

## Core Concepts

### Concurrency

Concurrency is the ability of a system to handle multiple tasks or processes at the same time. It's about **dealing with** lots of things at once. In a concurrent system, tasks can start, run, and complete in overlapping time periods, but they are not necessarily executing at the exact same instant. The system manages and makes progress on multiple tasks by switching between them.

- **Analogy:** A chef in a kitchen juggling multiple tasks. They might start boiling water for pasta, then chop vegetables, then check the sauce, then go back to the pasta. They are making progress on all dishes concurrently, but only doing one specific action at a time.

### Parallelism

Parallelism is the ability of a system to execute multiple tasks or parts of a single task **simultaneously**. It's about **doing** lots of things at once. This requires hardware with multiple processing units, such as a multi-core processor.

- **Analogy:** A team of chefs in a kitchen. One chef is dedicated to making the pasta, another is chopping vegetables, and a third is preparing the sauce. All these tasks are happening at the exact same time, leading to the meal being prepared much faster.

## Concurrency vs. Parallelism: Key Differences

| Feature        | Concurrency                                                                                     | Parallelism                                           |
| -------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------- |
| **Definition** | Managing multiple tasks, not necessarily at the same time.                                      | Executing multiple tasks simultaneously.              |
| **Focus**      | Task management and coordination.                                                               | Performance through simultaneous execution.           |
| **Execution**  | Tasks might be interleaved or scheduled on a single core.                                       | Tasks run at the same time on different cores.        |
| **Use Case**   | Handling I/O-bound tasks (e.g., web requests, database queries), managing multiple connections. | Computation-heavy tasks, large-scale data processing. |

## Challenges and Considerations

Both concurrency and parallelism introduce unique challenges that developers must manage.

### Concurrency Challenges

- **Synchronization**: Ensuring that shared resources are accessed by only one task at a time to prevent data corruption.
- **Deadlocks**: A situation where two or more tasks are blocked forever, each waiting for the other to release a resource.

### Parallelism Challenges

- **Data Sharing**: Safely and efficiently sharing data between tasks running on different cores can be complex.
- **Overhead**: The cost of creating and managing threads or processes can sometimes outweigh the performance benefits for small tasks.

## Practical Example in Go

The following Go code demonstrates how to achieve parallelism for CPU-bound tasks.

### The Code

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// heavyTask simulates a CPU-intensive operation.
// It accepts an id for identification and a WaitGroup to signal completion.
func heavyTask(id int, wg *sync.WaitGroup) {
	// Defer wg.Done() to ensure it's called when the function exits.
	defer wg.Done()

	fmt.Printf("Task %d is starting...\n", id)

	// A simple, long loop to consume CPU time.
	for i := 0; i < 1_000_000_000; i++ {
	}

	fmt.Printf("Task %d is finished.\n", id)
}

func main() {
	// Record the start time to measure total execution time.
	startTime := time.Now()

	// Set the number of OS threads that can execute Go code simultaneously.
	// For parallelism, this should be > 1.
	numThreads := 4
	runtime.GOMAXPROCS(numThreads)

	// A WaitGroup is used to wait for a collection of goroutines to finish.
	var wg sync.WaitGroup

	fmt.Printf("Starting %d parallel tasks...\n", numThreads)

	// Launch multiple goroutines.
	for i := 0; i < numThreads; i++ {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a new goroutine to run heavyTask.
		go heavyTask(i+1, &wg)
	}

	// Block the main function until the WaitGroup counter is zero.
	// This means we wait for all heavyTask goroutines to finish.
	wg.Wait()

	fmt.Printf("All tasks completed in %v\n", time.Since(startTime))
}
```

### Code Breakdown

1.  **`heavyTask(id int, wg *sync.WaitGroup)`**: This function simulates a workload that keeps a CPU core busy.

    - `defer wg.Done()`: This is a crucial line. `defer` ensures that `wg.Done()` is called right before the function returns. `wg.Done()` decrements the `WaitGroup` counter, signaling that this goroutine has completed its work.

2.  **`runtime.GOMAXPROCS(numThreads)`**: This function sets the maximum number of CPU cores that the Go runtime can use simultaneously. By setting it to `4`, we are telling Go to run our goroutines in parallel on up to 4 different cores if available.

3.  **`var wg sync.WaitGroup`**: We create a `WaitGroup` to synchronize our program. The `main` goroutine will wait until all `heavyTask` goroutines have finished.

4.  **The Loop (`for i := ...`)**:

    - `wg.Add(1)`: Before launching a goroutine, we increment the `WaitGroup` counter by one.
    - `go heavyTask(i+1, &wg)`: We start a new goroutine using the `go` keyword. This task will now run concurrently (and in this case, in parallel) with the `main` goroutine.

5.  **`wg.Wait()`**: This call blocks the execution of `main`. The program will not proceed past this line until the `WaitGroup` counter becomes zero. Since we call `wg.Done()` in each `heavyTask`, the counter will reach zero only when all tasks are complete.

### How to Run

1.  Save the code above into a file named `main.go`.
2.  Open your terminal and navigate to the directory where you saved the file.
3.  Run the program using the following command:
    ```sh
    go run main.go
    ```

### Expected Output

You will see the tasks start at roughly the same time and, if you have a multi-core CPU, finish at roughly the same time. The total duration will be significantly less than if the tasks were run sequentially.

```
Starting 4 parallel tasks...
Task 1 is starting...
Task 4 is starting...
Task 2 is starting...
Task 3 is starting...
Task 4 is finished.
Task 2 is finished.
Task 1 is finished.
Task 3 is finished.
All tasks completed in 1.15s
```

_(Note: The order of task start/finish and the total time will vary based on your system's scheduler and CPU.)_
