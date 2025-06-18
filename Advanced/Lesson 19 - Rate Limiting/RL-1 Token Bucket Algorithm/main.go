package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucketRateLimiter struct {
	tokens         chan struct{}
	refillTime     time.Duration
	stopRefillChan chan struct{}
	refillRate     int
}

func NewTokenBucketRateLimiter(bucketCapacity int, refillRate int, refillDuration time.Duration) *TokenBucketRateLimiter {
	// Create a rate limiter with tokens buffered channel of bucket size
	rl := &TokenBucketRateLimiter{
		tokens:         make(chan struct{}, bucketCapacity),
		refillTime:     refillDuration,
		stopRefillChan: make(chan struct{}),
		refillRate:     refillRate,
	}

	// Initially fill the bucket with maximum allowable number of token
	for range bucketCapacity {
		rl.tokens <- struct{}{}
	}

	// Start the refill process
	go rl.refill()

	return rl
}

// create a method to refill the token periodically for TokenBucketLimiter
func (rl *TokenBucketRateLimiter) refill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			count := 0
			for range rl.refillRate {
				select {
				case rl.tokens <- struct{}{}:
					count++
				default:
					break
				}
			}
			if count == 0 {
				fmt.Println("Bucket full can't add more tokens.")
			} else {
				fmt.Printf("Added %d of %d tokens. Total remaining tokens: %d\n", count, rl.refillRate, len(rl.tokens))
			}
		case <-rl.stopRefillChan:
			fmt.Println("Stopping refill")
			return
		}
	}
}

// consumeToken func consumes a token if available and return true or false
// if consumed returns true
// if bucket is empty returns false
func (rl *TokenBucketRateLimiter) consumeToken(rID int) bool {

	select {
	case <-rl.tokens:
		fmt.Println("Request", rID, "consumed one token. Total remaining tokens:", len(rl.tokens))
		return true
	default:
		fmt.Println("Request", rID, "denied no available token.")
		return false
	}

}

// pauseBucketRefill stops process od refilling the bucket with tokens
func (rl *TokenBucketRateLimiter) pauseBucketRefill() {
	close(rl.stopRefillChan)
}

func main() {

	var wg sync.WaitGroup

	rl := NewTokenBucketRateLimiter(5, 2, 2*time.Second)
	defer rl.pauseBucketRefill()

	requests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for {
		if len(requests) == 0 {
			fmt.Println("All requests processed")
			break
		}
		val := requests[0]
		if rl.consumeToken(val) {
			wg.Add(1)
			requests = requests[1:]
			go func() {
				defer wg.Done()
				time.Sleep(time.Second)
			}()
		} else {
			time.Sleep(time.Second)
		}

	}

	wg.Wait()

}
