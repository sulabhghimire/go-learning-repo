# The Leaky Bucket Algorithm

## 1. What is the Leaky Bucket Algorithm?

The Leaky Bucket algorithm's primary goal is to **smooth out bursts of traffic** into a constant, predictable stream. It is typically implemented using a First-In-First-Out (FIFO) queue.

### The Analogy

Imagine a bucket with a small hole at the bottom.

1.  **The Bucket (Queue):** Incoming requests are "poured" into the bucket. The bucket has a finite capacity.
2.  **The Leak (Processor):** The bucket "leaks" requests out through the hole at a constant, fixed rate (e.g., 1 request per second). This represents processing the requests.
3.  **The Process:**
    - When a request arrives, the system checks if there is space in the bucket (queue).
    - If **yes**, the request is added to the queue.
    - If **no**, the bucket is full, and the new request is discarded (it overflows).
4.  **The Output:** The output from the bucket is always a steady stream, regardless of how bursty the input is.

The key takeaway is that the Leaky Bucket **does not allow bursts**; it forces all traffic into a uniform flow.

## 2. Important Clarification: Leaky Bucket vs. Token Bucket

The provided Go code demonstrates a common and understandable confusion between the Leaky Bucket and Token Bucket algorithms. **The code implements a Token Bucket.**

Here is a quick comparison to clarify the difference:

| Feature            | True Leaky Bucket                                   | Token Bucket (Implemented in the Go code)      |
| ------------------ | --------------------------------------------------- | ---------------------------------------------- |
| **Primary Goal**   | Enforce a strict, **constant output rate**.         | Enforce an **average rate**, but allow bursts. |
| **Output Traffic** | Smoothed out and constant.                          | Bursty and variable.                           |
| **Allows Bursts?** | **No**, all bursts are smoothed into a steady flow. | **Yes**, up to the bucket's capacity.          |
| **Implementation** | A queue with a fixed processing rate.               | A counter of "tokens" that refills over time.  |

With this clarification, we will now analyze the provided Go code as a **Token Bucket** implementation.

## 3. The Go Implementation: A Token Bucket Example

This code uses a counter (`tokens`) that is periodically refilled, which is the defining characteristic of the Token Bucket algorithm.

### The Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// NOTE: This struct implements the Token Bucket algorithm.
type LeakyBucket struct {
	capacity int
	leakRate time.Duration // In this implementation, this is the time it takes to gain 1 token.
	tokens   int
	lastLeak time.Time     // The timestamp of the last token calculation.
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		lastLeak: time.Now(),
		tokens:   capacity, // Start with a full bucket of tokens.
	}
}

// Allow checks if a request is permitted by consuming a token.
func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)

	// Calculate how many tokens should have been refilled since the last check.
	tokensToAdd := int(elapsedTime / lb.leakRate)
	if tokensToAdd > 0 {
		lb.tokens += tokensToAdd
		if lb.tokens > lb.capacity {
			lb.tokens = lb.capacity // Don't exceed capacity.
		}
		// Advance the lastLeak timestamp by the amount of time accounted for.
		// This prevents losing "fractional" time for more accurate refills.
		lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
	}

	// Check if a token is available to be consumed.
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	return false
}

func main() {
	var wg sync.WaitGroup
	// Limiter: 5 token capacity, refills 1 token every 500ms (2 tokens/sec).
	lb := NewLeakyBucket(5, 500*time.Millisecond)

	// Launch 20 goroutines in two quick bursts to test the limiter.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(reqID int) {
			defer wg.Done()
			if lb.Allow() {
				fmt.Printf("Request %d: Allowed\n", reqID)
			} else {
				fmt.Printf("Request %d: Denied\n", reqID)
			}
		}(i + 1)
		// Small delay to create two semi-distinct bursts
		if i == 9 {
			time.Sleep(500 * time.Millisecond)
		}
	}
	wg.Wait()
}
```

## 4. How to Run

1.  Save the code above as `main.go`.
2.  Execute it from your terminal:
    ```sh
    go run .
    ```

## 5. Example Walkthrough and Output Analysis

The `main` function creates a limiter with a **capacity of 5** and a refill rate of **1 token every 500ms** (equivalent to 2 tokens per second). It then fires off 20 concurrent requests.

**Important Note:** Because the requests are launched in goroutines, the exact order of the output is **non-deterministic** and will vary slightly on each run. However, the overall behavior will be consistent.

#### Predicted Behavior:

1.  **Initial Burst:** The bucket starts with 5 tokens. The first 5 goroutines that win the race to call `Allow()` will be **allowed**.
2.  **Initial Denials:** The next several requests that arrive before the bucket has time to refill will be **denied**.
3.  **Refilling:** Every 500 milliseconds, the logic will calculate that 1 new token should be added to the bucket.
4.  **Sustained Rate:** After the initial burst, requests will be allowed at a steady rate of approximately 2 per second, as tokens become available. Any requests arriving faster than this rate will be denied.

#### A Possible Output Might Look Like This:

```text
Request 2: Allowed
Request 1: Allowed
Request 4: Allowed
Request 3: Allowed
Request 5: Allowed
Request 6: Denied
Request 8: Denied
Request 7: Denied
Request 9: Denied
Request 10: Denied
Request 11: Denied
Request 12: Allowed  // A token was refilled
Request 13: Denied
Request 14: Allowed  // Another token was refilled
Request 15: Denied
Request 16: Allowed
Request 17: Denied
Request 18: Allowed
Request 19: Denied
Request 20: Allowed
```

## 6. Pros and Cons of the Leaky Bucket Algorithm

Here we discuss the pros and cons of the _actual_ Leaky Bucket algorithm (the queue-based one).

#### Pros

- **Smooth Output:** Its main advantage. It guarantees a constant, predictable output stream, which is ideal for services that require a steady processing rate.
- **Predictable:** The behavior is easy to reason about, as the outflow is always fixed.

#### Cons

- **Throttles Bursts:** A burst of requests will fill the queue and sit there, increasing their latency. The system doesn't speed up to accommodate the burst.
- **Less Flexible:** It's not ideal for use cases where allowing occasional bursts is desirable and beneficial.

## 7. When to Use the (True) Leaky Bucket Algorithm

The Leaky Bucket algorithm is best suited for scenarios where a constant output rate is a requirement.

- **Media Streaming:** Delivering video or audio data at a stable rate to ensure smooth playback.
- **Job Schedulers:** Processing jobs from a queue at a predictable pace to avoid overwhelming downstream services.
- **Throttling Network Packets:** Controlling data flow in network devices to ensure stable bandwidth usage.
