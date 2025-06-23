# `sync.Once`

`sync.Once` is a synchronization primitive in Go that provides a guarantee: a specific piece of code will be executed **exactly once**, no matter how many goroutines try to execute it simultaneously or sequentially.

Its primary purpose is for "lazy initialization"—deferring the creation of an object or the execution of a setup task until the first time it is needed. This avoids the overhead of initializing resources that may never be used, while ensuring that the initialization is done safely in a concurrent environment.

## Key Concepts

### The `Do` Method

`sync.Once` has only one exported method:

`func (o *Once) Do(f func())`

- It takes a single argument: a function `f` with the signature `func()`.
- The first time `Do` is called on a given `sync.Once` instance, it invokes the function `f`.
- All subsequent calls to `Do` on the **same** `sync.Once` instance will do nothing. They will not execute `f` again.
- `Do` is blocking. If one goroutine calls `Do(f)` and `f` is running, other goroutines that call `Do(f)` will block until `f` has completed. This ensures that any caller of `Do` can be certain the initialization is finished when `Do` returns.

## How `sync.Once` Works (Under the Hood)

A `sync.Once` object contains two fields: a boolean flag (implemented as an atomic `uint32`) to track if the function has been run, and a `sync.Mutex` to handle the race condition of multiple goroutines calling `Do` at the same time.

The logic inside `Do` is roughly:

1.  **Fast Check (Atomic):** First, it performs a quick, lock-free check of the `done` flag. If it's already set, the method returns immediately. This makes subsequent calls extremely cheap.
2.  **Slow Path (Mutex):** If the flag is not set, it acquires the mutex.
3.  **Second Check (Double-Checked Locking):** After acquiring the lock, it checks the `done` flag _again_. This is crucial because another goroutine might have completed the initialization while this one was waiting for the lock.
4.  **Execute and Mark Done:** If the flag is still not set, it executes the function `f`. After `f` returns, it sets the `done` flag.
5.  **Unlock:** Finally, it releases the mutex.

This double-checked locking pattern makes `sync.Once` highly efficient.

## Example Usage: Lazy-Initialized Singleton

This is the most common use case for `sync.Once`. A singleton is a design pattern where only one instance of a type is ever created.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Logger struct {
	level string
}

var (
	once     sync.Once
	instance *Logger
)

// GetInstance creates and returns the singleton Logger instance.
// The initialization code inside the Do() function will only run once.
func GetInstance() *Logger {
	once.Do(func() {
		fmt.Println("Initializing logger instance now...")
		// Simulate a slow initialization (e.g., reading a config file)
		time.Sleep(1 * time.Second)
		instance = &Logger{level: "INFO"}
	})
	return instance
}

func main() {
	var wg sync.WaitGroup
	// Start 10 goroutines that all try to get the logger instance concurrently.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			logger := GetInstance()
			fmt.Printf("Goroutine %d got logger instance: %p\n", id, logger)
		}(i)
	}
	wg.Wait()
}
```

**Output:**

```
Initializing logger instance now...
Goroutine 8 got logger instance: 0x1400010e018
Goroutine 0 got logger instance: 0x1400010e018
Goroutine 1 got logger instance: 0x1400010e018
Goroutine 3 got logger instance: 0x1400010e018
Goroutine 4 got logger instance: 0x1400010e018
Goroutine 5 got logger instance: 0x1400010e018
Goroutine 2 got logger instance: 0x1400010e018
Goroutine 6 got logger instance: 0x1400010e018
Goroutine 7 got logger instance: 0x1400010e018
Goroutine 9 got logger instance: 0x1400010e018
```

Notice that "Initializing logger instance now..." is printed only **once**, and all goroutines receive a pointer to the **same memory address**, proving it's a true singleton.

---

## Best Practices and Pitfalls

### Handling Initialization with Parameters or Errors

The function signature for `Do` is `func()`. It takes no arguments and returns no values. To handle initialization that requires parameters or might fail, use a closure to capture variables from the surrounding scope.

```go
var (
	dbOnce sync.Once
	db     *Database
	dbErr  error
)

// connectToDB needs to be called only once.
// It returns both the connection and a potential error.
func connectToDB(dsn string) (*Database, error) {
	dbOnce.Do(func() {
		// The actual connection logic is inside the closure.
		// The results are stored in the package-level variables db and dbErr.
		conn, err := realConnect(dsn) // realConnect is your actual DB connection func
		if err != nil {
			dbErr = err
			return
		}
		db = conn
	})
	return db, dbErr
}
```

### Handling Panics

A crucial detail of `sync.Once` is how it handles panics. If the function `f` passed to `Do` panics, `sync.Once` considers the call to have failed. It does **not** mark the action as "done." This allows subsequent calls to `Do` to attempt the initialization again.

```go
func main() {
	var once sync.Once
	var i int

	// First attempt will panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		once.Do(func() {
			i++
			fmt.Println("First attempt, i is", i)
			panic("something went wrong")
		})
	}()

	// Second attempt will succeed because the first one panicked
	once.Do(func() {
		i++
		fmt.Println("Second attempt, i is", i)
	})

	fmt.Println("Final value of i:", i)
}
```

**Output:**

```
First attempt, i is 1
Recovered from panic: something went wrong
Second attempt, i is 2
Final value of i: 2
```

### `sync.Once` vs. `init()`

The `init()` function is a language feature for package-level initialization. `sync.Once` is a library primitive for lazy initialization.

- **Use `init()` when:** The initialization is required for the package to work at all, is relatively fast, and has no failure conditions you need to handle gracefully. It's an **eager** initialization that runs when the package is imported.
- **Use `sync.Once` when:** The initialization is expensive and may not be needed, or it needs to happen "on-demand" the first time a certain function is called. It's a **lazy** initialization.
