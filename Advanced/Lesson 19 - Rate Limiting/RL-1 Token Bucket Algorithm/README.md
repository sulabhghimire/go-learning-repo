# The Token Bucket Algorithm

## 1. What is the Token Bucket Algorithm?

The Token Bucket is a rate limiting algorithm that controls the rate of traffic by modeling a bucket that holds "tokens." Each token represents permission to perform an action, like making an API request.

Its key characteristic is the ability to allow **bursts of traffic** up to the bucket's capacity, while still enforcing a sustained, average rate over time.

### The Analogy

Imagine a bucket with a fixed capacity:

1.  **The Bucket**: Represents the client's allowance. It has a maximum size (e.g., it can hold at most 100 tokens).
2.  **The Tokens**: Each token is a permit for one request. To make a request, a client must "spend" one token.
3.  **The Request Process**: When a request arrives, the system checks if there is at least one token in the bucket.
    - If **yes**, a token is removed, and the request is allowed.
    - If **no**, the bucket is empty, and the request is rejected.
4.  **The Refill Process**: An independent process adds new tokens to the bucket at a fixed rate (e.g., 10 tokens per second). If the bucket is already full, new tokens are discarded.

This mechanism allows an idle client to accumulate tokens, which can then be spent all at once in a burst.

## 2. The Go Implementation: An In-Depth Look

This implementation uses a buffered channel (`chan struct{}`) as the token bucket. This is an idiomatic and powerful approach in Go.

- **The Bucket** is a `chan struct{}`.
- **Bucket Capacity** is the buffer size of the channel.
- **A Token** is an empty struct `struct{}{}` sent to the channel.
- **Consuming a Token** means receiving a value from the channel.
- **Refilling the Bucket** means sending values to the channel from a background goroutine.

### The Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucketRateLimiter holds the state for our rate limiter.
type TokenBucketRateLimiter struct {
	tokens         chan struct{}     // The bucket, implemented as a buffered channel.
	refillTime     time.Duration   // The interval between refills.
	stopRefillChan chan struct{}     // A channel to signal the refill goroutine to stop.
	refillRate     int             // The number of tokens to add per refill interval.
}

// NewTokenBucketRateLimiter creates and starts a new rate limiter.
func NewTokenBucketRateLimiter(bucketCapacity int, refillRate int, refillDuration time.Duration) *TokenBucketRateLimiter {
	// Create a rate limiter with a buffered channel for tokens.
	rl := &TokenBucketRateLimiter{
		tokens:         make(chan struct{}, bucketCapacity),
		refillTime:     refillDuration,
		stopRefillChan: make(chan struct{}),
		refillRate:     refillRate,
	}

	// Initially fill the bucket with the maximum allowable number of tokens.
	for i := 0; i < bucketCapacity; i++ {
		rl.tokens <- struct{}{}
	}

	// Start the refill process in a separate goroutine.
	go rl.refill()

	return rl
}

// refill runs in the background, adding tokens periodically.
func (rl *TokenBucketRateLimiter) refill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// A tick has occurred, time to add tokens.
			count := 0
			for i := 0; i < rl.refillRate; i++ {
				select {
				case rl.tokens <- struct{}{}:
					count++
				default:
					// The bucket is full, so we stop trying to add tokens.
					break
				}
			}
			if count == 0 {
				fmt.Println("Bucket full, can't add more tokens.")
			} else {
				fmt.Printf("Added %d of %d tokens. Total available tokens: %d\n", count, rl.refillRate, len(rl.tokens))
			}
		case <-rl.stopRefillChan:
			// A stop signal was received.
			fmt.Println("Stopping refill goroutine.")
			return
		}
	}
}

// consumeToken attempts to take one token from the bucket.
func (rl *TokenBucketRateLimiter) consumeToken(rID int) bool {
	select {
	case <-rl.tokens:
		// A token was available and consumed.
		fmt.Println("Request", rID, "consumed one token. Total available tokens:", len(rl.tokens))
		return true
	default:
		// No token was available.
		fmt.Println("Request", rID, "denied; no available token.")
		return false
	}
}

// pauseBucketRefill gracefully stops the background refill goroutine.
func (rl *TokenBucketRateLimiter) pauseBucketRefill() {
	close(rl.stopRefillChan)
}

func main() {
	var wg sync.WaitGroup

	// Limiter: 5 token capacity, refills 2 tokens every 2 seconds.
	rl := NewTokenBucketRateLimiter(5, 2, 2*time.Second)
	defer rl.pauseBucketRefill()

	// A queue of requests to process.
	requests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for {
		if len(requests) == 0 {
			fmt.Println("All requests processed.")
			break
		}

		val := requests[0]
		if rl.consumeToken(val) {
			// If a token was consumed, process the request.
			wg.Add(1)
			requests = requests[1:] // Remove from queue
			go func() {
				defer wg.Done()
				time.Sleep(time.Millisecond * 100) // Simulate work
			}()
		} else {
			// If denied, wait before retrying.
			time.Sleep(time.Second)
		}
	}

	wg.Wait()
}
```

```text
Request 1 consumed one token. Total available tokens: 4
Request 2 consumed one token. Total available tokens: 3
Request 3 consumed one token. Total available tokens: 2
Request 4 consumed one token. Total available tokens: 1
Request 5 consumed one token. Total available tokens: 0
Request 6 denied; no available token.
Request 6 denied; no available token.
Added 2 of 2 tokens. Total available tokens: 2
Request 6 consumed one token. Total available tokens: 1
Request 7 consumed one token. Total available tokens: 0
Request 8 denied; no available token.
Request 8 denied; no available token.
Added 2 of 2 tokens. Total available tokens: 2
Request 8 consumed one token. Total available tokens: 1
Request 9 consumed one token. Total available tokens: 0
Request 10 denied; no available token.
Request 10 denied; no available token.
Added 2 of 2 tokens. Total available tokens: 2
Request 10 consumed one token. Total available tokens: 1
All requests processed.
Stopping refill goroutine.
```

#### Step-by-Step Analysis:

1.  **Initial Burst (Requests 1-5):** The program starts. The bucket is full with 5 tokens. The first 5 requests are immediately successful and consume all the tokens.
2.  **Bucket Empty (Request 6):** When request 6 arrives, `consumeToken` finds the bucket empty and denies it. The `main` loop now waits for 1 second before retrying the same request.
3.  **The First Refill:** After approximately 2 seconds from the start, the `refill` goroutine's ticker fires. The log `Added 2 of 2 tokens` appears. The bucket now contains 2 tokens.
4.  **Processing Continues (Requests 6-7):** The next retry for request 6 succeeds, consuming one of the new tokens. Request 7 immediately follows and consumes the second new token. The bucket is empty again.
5.  **Another Cycle:** This pattern repeats. Request 8 is denied until the next refill cycle adds 2 more tokens, after which it succeeds, and so on until all 10 requests have been processed.
6.  **Graceful Shutdown:** Once all requests are processed, the `main` function exits. The `defer rl.pauseBucketRefill()` statement is executed, which closes the `stopRefillChan` and allows the background goroutine to terminate cleanly.

## 3. Pros and Cons of this Implementation

#### Pros

- **Idiomatic Go**: Using buffered channels is a powerful and elegant way to handle rate limiting.
- **Thread-Safe by Design**: No explicit `sync.Mutex` is needed, as channel operations are inherently safe for concurrent use.
- **Decoupled Logic**: The refill logic runs completely independently in the background, decoupled from the token consumption logic.
- **Graceful Shutdown**: The use of a `stopRefillChan` is a best practice for managing the lifecycle of goroutines.

#### Cons

- **Chunk-Based Refills**: The `time.Ticker` refills tokens in discrete chunks (e.g., 2 every 2 seconds). This is slightly different from a theoretical model where tokens are added continuously. For most use cases, this is perfectly acceptable.
- **`len(channel)` for Logging**: While useful for logging, calling `len()` on a channel in highly contentious scenarios can return a slightly stale value. For this example's purpose, it's perfectly fine.

## 4. When to Use the Token Bucket Algorithm

- **Public APIs**: Ideal for APIs where clients may need to send a batch of requests in a short time (e.g., syncing data, uploading multiple files).
- **Throttling Outgoing Calls**: When your application needs to call a third-party API that has its own rate limit, you can use a token bucket to ensure you don't exceed their limits.
- **General Purpose Rate Limiting**: It's an excellent default choice for rate limiting due to its flexibility in handling both sustained traffic and bursts.
