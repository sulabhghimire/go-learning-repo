# The Fixed Window Counter Algorithm in Go

## 1. What is the Fixed Window Counter Algorithm?

The Fixed Window Counter is a rate limiting algorithm that tracks requests within a fixed time window (e.g., one minute, one hour). It works by maintaining a counter for the current window and resetting it every time a new window begins.

### The Analogy

Imagine a bouncer at a club door who wants to let in a maximum of 100 people per hour.

1.  **The Window:** The bouncer looks at the clock. The current window is from 12:00 PM to 12:59 PM.
2.  **The Counter:** The bouncer uses a simple clicker counter, starting at 0.
3.  **The Process:** Every time a person enters, the bouncer clicks the counter. If the counter reaches 100, no one else is allowed in until the next hour.
4.  **The Reset:** At 1:00 PM, the bouncer resets the clicker back to 0 and starts counting again for the new window (1:00 PM to 1:59 PM).

## 2. The Core Problem: The Edge Burst Issue

The main weakness of the Fixed Window algorithm is its vulnerability to a burst of traffic at the "edge" where two windows meet.

Consider a limit of **100 requests per minute**:

- **Window 1:** Ends at `12:00:59`.
- **Window 2:** Starts at `12:01:00`.

A malicious or aggressive client could send:

- 100 requests at `12:00:59`. These are **allowed** as they are within the limit for Window 1.
- 100 requests at `12:01:00`. The window has just reset, so these are also **allowed**.

The result is that **200 requests** are processed in just two seconds, temporarily doubling the intended rate limit. For many systems, this is an unacceptable flaw.

## 3. The Go Implementation: An In-Depth Look

The provided code implements a thread-safe Fixed Window Counter.

### The Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter holds the state for the fixed window counter.
type RateLimiter struct {
	mu        sync.Mutex    // Ensures thread safety.
	count     int           // The number of requests in the current window.
	limit     int           // The maximum number of requests allowed per window.
	window    time.Duration // The duration of the time window.
	resetTime time.Time     // The time when the counter should be reset.
}

// NewRateLimiter creates a new fixed window rate limiter.
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

// Allow checks if a request is permitted.
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	// This is the core logic: check if the current window has expired.
	if now.After(rl.resetTime) {
		// If it has, start a new window.
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	// Check if the current request is within the limit for the current window.
	if rl.count < rl.limit {
		rl.count++
		return true
	}

	// The limit has been reached for this window.
	return false
}

func main() {
	// Limiter: 3 requests per 1 second.
	rateLimiter := NewRateLimiter(3, 1*time.Second)

	// Send 10 requests, spaced 200ms apart.
	for i := 0; i < 10; i++ {
		if rateLimiter.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(time.Millisecond * 200)
	}
}
```

### Code Breakdown

- **`RateLimiter` Struct**:

  - `mu sync.Mutex`: A mutex to make the `Allow` method thread-safe, preventing race conditions when multiple goroutines call it.
  - `count int`: Tracks the number of requests made in the _current_ time window.
  - `limit int`: The configured request limit for the window.
  - `window time.Duration`: The length of the window (e.g., `1 * time.Second`).
  - `resetTime time.Time`: The key state variable. It stores the timestamp when the current window expires.

- **`Allow()` Method**:
  1.  It first acquires a lock to ensure exclusive access.
  2.  It gets the current time, `now`.
  3.  The crucial check: `if now.After(rl.resetTime)`. This determines if the current time has passed the expiration point of the previous window.
      - If `true`, a **new window starts**. The `resetTime` is updated to be one `window` duration in the future, and the `count` is reset to `0`.
  4.  It then checks if the `count` for the current window is still less than the `limit`.
      - If `true`, the `count` is incremented, and the request is allowed.
      - If `false`, the limit has been reached, and the request is denied.

## 4. Pros and Cons

#### Pros

- **Simple to Implement:** The logic is straightforward and easy to understand.
- **Memory Efficient:** It only needs to store a counter and a timestamp, making it very lightweight.

#### Cons

- **The Edge Burst Problem:** Its primary weakness. It can allow up to twice the intended number of requests at the boundary of two windows.
- **Inflexible:** It can lead to a "starvation" effect where traffic that arrives early in a window can use up all the quota, blocking all subsequent requests for the rest of the window.

## 5. When to Use It

The Fixed Window Counter is suitable for:

- Internal services or low-stakes applications where perfect accuracy is not required.
- Scenarios where simplicity and performance are more important than protecting against edge bursts.
- Educational purposes to understand the basics of rate limiting before moving to more advanced algorithms like the Sliding Window or Token Bucket.
