# Goroutines in Go

This guide provides a comprehensive overview of Goroutines, a fundamental feature of the Go programming language for building concurrent applications.

## Table of Contents

- [What are Goroutines?](#what-are-goroutines)
- [Why Use Goroutines?](#why-use-goroutines)
- [Basics of Goroutines](#basics-of-goroutines)
  - [Creating a Goroutine](#creating-a-goroutine)
  - [Goroutine Lifecycle](#goroutine-lifecycle)
- [Goroutine Scheduling](#goroutine-scheduling)
  - [M:N Scheduling Model](#mn-scheduling-model)
  - [Efficient Multiplexing](#efficient-multiplexing)
- [Concurrency vs. Parallelism](#concurrency-vs-parallelism)
- [Common Pitfalls and Best Practices](#common-pitfalls-and-best-practices)
  - [Synchronization](#1-synchronization)
  - [Avoiding Goroutine Leaks](#2-avoiding-goroutine-leaks)
  - [Limiting Goroutine Creation](#3-limiting-goroutine-creation)
  - [Proper Error Handling](#4-proper-error-handling)
- [Handling Errors in Goroutines](#handling-errors-in-goroutines)

## What are Goroutines?

- **Lightweight Threads:** Goroutines are lightweight, concurrently executing functions managed by the Go runtime, not the OS.
- **Enable Concurrency:** They allow you to perform multiple tasks concurrently within a single Go program.
- **Non-Blocking:** Goroutines run in the background without blocking the main program flow.
- **A Key Go Feature:** Goroutines are central to Go's philosophy, making it simple to write highly concurrent and parallel programs.

## Why Use Goroutines?

- **Simplify Concurrent Programming:** They provide a high-level, easy-to-use abstraction over threads.
- **Efficient Task Handling:** Ideal for handling I/O operations, complex calculations, or any task that can run in parallel.
- **No Manual Thread Management:** You can perform tasks concurrently without the complexity of creating and managing OS threads yourself.

## Basics of Goroutines

### Creating a Goroutine

You can start a new goroutine by using the `go` keyword followed by a function call. This tells the Go runtime to execute the function concurrently with the rest of the program.

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from the goroutine!")
}

func main() {
	// Start a new goroutine
	go sayHello()

	fmt.Println("Hello from the main function!")

	// The main goroutine needs to wait, otherwise it will exit
	// before the new goroutine has a chance to run.
	// NOTE: This is a naive way to wait. Use sync.WaitGroup for real applications.
	time.Sleep(1 * time.Second)
}
```

### Goroutine Lifecycle

- A goroutine **starts** when it is created with the `go` keyword.
- It runs **concurrently** with other goroutines in the same address space.
- A goroutine **exits** when the function it is executing completes.
- The Go runtime scheduler is responsible for managing the entire lifecycle, including scheduling and execution.

## Goroutine Scheduling

The Go runtime has a sophisticated scheduler to manage thousands of goroutines efficiently.

### M:N Scheduling Model

The scheduler uses an **M:N model**, which means it maps **M** goroutines onto **N** operating system (OS) threads. Typically, `N` is a small number (equal to the number of CPU cores). This allows Go to manage a massive number of goroutines with a minimal number of expensive OS threads, improving both efficiency and scalability.

### Efficient Multiplexing

The Go runtime scheduler efficiently **multiplexes** (or switches) goroutines onto the available OS threads. It dynamically schedules and reschedules goroutines as needed, ensuring that no thread is idle if there is work to be done. This leads to high concurrency and excellent performance.

## Concurrency vs. Parallelism

It's important to understand the distinction between these two concepts:

- **Concurrency:** Multiple tasks making progress simultaneously, but not necessarily executing at the exact same time. Think of a chef juggling multiple cooking tasks in one kitchen.
- **Parallelism:** Tasks are executed at the exact same time on multiple processors or cores. Think of multiple chefs, each with their own kitchen, working simultaneously.

Goroutines facilitate **concurrency**. The Go runtime scheduler enables **parallelism** by distributing these concurrent goroutines across available CPU cores when possible.

## Common Pitfalls and Best Practices

### 1. Synchronization

Since goroutines run independently, you must coordinate them when they need to share data or when one must wait for another to finish.

- **`sync.WaitGroup`**: To wait for a collection of goroutines to complete their execution.
- **`sync.Mutex`**: To protect shared data from being accessed by multiple goroutines at the same time (race conditions).

**Example with `sync.WaitGroup`:**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    // Decrement the counter when the goroutine completes
    defer wg.Done()

    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        // Increment the WaitGroup counter
        wg.Add(1)
        go worker(i, &wg)
    }

    // Block until the WaitGroup counter is zero
    wg.Wait()
    fmt.Println("All workers have finished.")
}
```

### 2. Avoiding Goroutine Leaks

A goroutine leak occurs when a goroutine is started but never finishes, consuming memory and CPU resources indefinitely. Always ensure every goroutine has a clear exit path. Leaks often happen with channels that are never written to or read from, causing the goroutine to block forever.

### 3. Limiting Goroutine Creation

While goroutines are lightweight, creating an unlimited number can still exhaust system resources. For tasks like processing items from a large list, use a **worker pool pattern** to limit the number of concurrently running goroutines.

### 4. Proper Error Handling

Errors that occur inside a goroutine are isolated to that goroutine and will not automatically crash your program or propagate to the calling function. You must implement a mechanism to communicate errors back.

## Handling Errors in Goroutines

The concept of **error propagation** is crucial. Since goroutines run concurrently, any errors they encounter must be communicated back to the main thread or a controlling goroutine.

- **Use Channels (Idiomatic Approach):** The most common and idiomatic way to handle errors is to use a dedicated channel. The goroutine can send an error on this channel, and the main thread can receive it.

- **Use Shared Error Variables:** Another method is to use a shared variable protected by a mutex. The goroutine can write an error to this variable, and the main thread can read it after synchronization. However, channels are generally preferred as they are a core and safer part of Go's concurrency model.
