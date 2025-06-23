# `sync.Cond`

`sync.NewCond` is a factory function in Go's `sync` package that creates a new **condition variable**. A condition variable is a synchronization primitive that allows goroutines to wait for a specific condition to become true.

It provides a way for goroutines to suspend their execution (go to sleep) until they are notified that something they are waiting for has happened. This is far more efficient than busy-waiting or polling (repeatedly acquiring a lock, checking a condition, and releasing the lock).

Condition variables are essential for more complex synchronization scenarios where a simple `Mutex` or `RWMutex` is not enough. They are used when a goroutine needs to wait for a change in the state of shared data that is protected by a mutex.

## Key Concepts of `sync.Cond`

### 1. Condition Variables

A condition variable doesn't protect data itself. Instead, it works in conjunction with a lock to manage goroutines that are waiting for a certain _condition_ (a specific state of the shared data) to be met. It essentially acts as a waiting room or queue for goroutines.

- A goroutine enters the "waiting room" (`Wait()`) if the condition is not met.
- Another goroutine that changes the state to meet the condition can "signal" the waiting room (`Signal()` or `Broadcast()`) to wake up one or all of the waiting goroutines.

### 2. Mutex and Condition Variables

A `sync.Cond` is always associated with a `sync.Locker` (which is typically a `*sync.Mutex` or `*sync.RWMutex`). This lock serves a critical purpose: **it protects the condition being checked.**

The workflow is:

1.  A goroutine acquires the lock.
2.  It checks the shared data to see if the condition is true.
3.  If the condition is **false**, it calls `Wait()`. The `Wait()` method will **atomically release the lock and put the goroutine to sleep.**
4.  When another goroutine signals it, the `Wait()` method wakes up and **atomically re-acquires the lock** before returning.
5.  The goroutine can now safely re-check the condition and proceed.

This atomic release-and-wait, wake-and-reacquire behavior is what prevents race conditions and makes `sync.Cond` so powerful.

## Methods of `sync.Cond`

A `sync.Cond` instance has three primary methods. To use them, you must first create a `Cond` with `sync.NewCond(locker)`.

### 1. `Wait()`

`func (c *Cond) Wait()`

- **Prerequisite:** The goroutine calling `Wait()` **must be holding the lock** `c.L`.
- **Action:** It atomically unlocks `c.L` and suspends the execution of the goroutine.
- **Wake-up:** When woken by `Signal()` or `Broadcast()`, it re-acquires the lock `c.L` before it returns.
- **Usage:** Because of "spurious wakeups" and other race conditions, `Wait()` must always be called inside a `for` loop that checks the condition.

  ```go
  // Correct usage of Wait()
  c.L.Lock()
  for !condition { // The loop is essential!
      c.Wait()
  }
  // ... Now the condition is met and the lock is held ...
  c.L.Unlock()
  ```

### 2. `Signal()`

`func (c *Cond) Signal()`

- Wakes up **one** goroutine that is waiting on the condition `c`, if there is one.
- It is recommended to hold the lock `c.L` when calling `Signal()` or `Broadcast()`.
- **Use Case:** Use `Signal()` when you know that only one waiting goroutine can make progress after the condition change (e.g., adding a single item to a queue can only be consumed by one consumer).

### 3. `Broadcast()`

`func (c *Cond) Broadcast()`

- Wakes up **all** goroutines that are waiting on the condition `c`.
- **Use Case:** Use `Broadcast()` when a state change might allow multiple waiting goroutines to proceed (e.g., a "gate" is opened) or when different goroutines are waiting on different aspects of the same state, and you can't be sure which one should be woken up.

---

## Example Usage: Producer-Consumer

This is the classic example for `sync.Cond`. We have a shared queue. Producers add items to the queue, and consumers remove them. Consumers must wait if the queue is empty.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	queue := make([]int, 0)

	// Consumer
	go func() {
		for {
			cond.L.Lock() // Acquire the lock
			// The condition is "queue is not empty" (len(queue) > 0)
			// We must wait while the condition is false (len(queue) == 0)
			for len(queue) == 0 {
				fmt.Println("Consumer is waiting, queue empty.")
				cond.Wait() // Atomically unlocks and waits. Re-locks on wakeup.
			}

			// At this point, the queue is not empty and we hold the lock.
			item := queue[0]
			queue = queue[1:]
			fmt.Printf("Consumer processed: %d, queue size: %d\n", item, len(queue))

			cond.L.Unlock() // Release the lock
			time.Sleep(1 * time.Second)
		}
	}()

	// Producer
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		cond.L.Lock() // Acquire the lock to modify the queue

		queue = append(queue, i)
		fmt.Printf("Producer added: %d, new queue size: %d\n", i, len(queue))

		// The condition "queue is not empty" is now true.
		// Signal a waiting consumer.
		cond.Signal()

		cond.L.Unlock() // Release the lock
	}

	time.Sleep(5 * time.Second) // Wait for some processing to happen
}
```

---

## Best Practices for Using `sync.Cond`

### 1. Ensure Mutex is Held

Always hold the associated lock (`c.L`) when calling `Wait()`, `Signal()`, or `Broadcast()`. Failing to do so for `Wait()` will cause a panic. While not a panic for `Signal`/`Broadcast`, not holding the lock breaks the guaranteed atomicity and can lead to race conditions.

### 2. Avoid Spurious Wakeups (Use a `for` loop)

A goroutine waiting on `cond.Wait()` might wake up even if no `Signal` or `Broadcast` was sent (a "spurious wakeup"). Furthermore, even if woken by a signal, another goroutine might have acquired the lock first and changed the state back.
The only robust way to handle this is to re-check the condition in a loop.

**BAD:**

```go
if !condition {
    c.Wait() // Unsafe! What if it wakes up but condition is still false?
}
```

**GOOD:**

```go
for !condition {
    c.Wait() // Safe! Will re-check the condition on every wakeup.
}
```

### 3. Use Condition Variables Judiciously

For many common synchronization patterns in Go (like producer-consumer), channels are often a simpler, more idiomatic, and safer choice. `sync.Cond` is a lower-level primitive. Use it when you need more complex control that channels don't easily provide, such as:

- Needing to `Broadcast` to multiple goroutines.
- Coordinating around complex state that isn't just a simple queue.
- Integrating with code that already uses mutexes extensively.

### 4. Balance `Signal` and `Broadcast`

- Prefer `Signal()` when only one goroutine can make progress. It is more efficient as it avoids waking up a thundering herd of goroutines that will just check the condition and go back to sleep.
- Use `Broadcast()` when multiple goroutines can proceed or when you are unsure which of the waiting goroutines should be woken.

---

## Advanced Use Cases

### 1. Task Scheduling

A scheduler can have a pool of worker goroutines. When there are no tasks in the queue, the workers can `Wait()` on a condition variable. When a new task is added to the queue, the scheduler can `Signal()` one worker to wake up and process it. `Broadcast()` could be used to signal all workers to shut down.

### 2. Resource Pools

A connection pool (e.g., for a database) can use a `Cond` to manage goroutines waiting for a connection to become available.

- **Borrower:** Locks the pool, checks for an available connection. If none, it calls `Wait()`.
- **Returner:** Locks the pool, adds a connection back, and calls `Signal()` to wake up one waiting borrower.

### 3. Event Notification Systems

An event bus can have multiple listeners waiting for specific events. When an event is published, the bus can lock, update its state, and then use `Broadcast()` to notify all waiting listeners. Each listener, upon waking up, will re-check the event state to see if it's relevant to them.
