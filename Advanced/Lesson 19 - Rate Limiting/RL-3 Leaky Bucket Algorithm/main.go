package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		lastLeak: time.Now(),
		tokens:   capacity,
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)

	// fmt.Printf("Tokens added %d, tokens subtracted %d, Total tokens: %d\n", tokensToAdd, 1, lb.tokens)
	// fmt.Printf("Last leak time: %v\n", lb.lastLeak)
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false
}

func main() {

	var wg sync.WaitGroup
	lb := NewLeakyBucket(5, 500*time.Millisecond)

	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if lb.Allow() {
				// fmt.Println("Current time:", time.Now())
				fmt.Println("Request allowed", i)
			} else {
				// fmt.Println("Current time:", time.Now())
				fmt.Println("Request denied", i)
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}

	time.Sleep(500 * time.Millisecond)

	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if lb.Allow() {
				// fmt.Println("Current time:", time.Now())
				fmt.Println("Request allowed", i)
			} else {
				// fmt.Println("Current time:", time.Now())
				fmt.Println("Request denied", i)
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}

	wg.Wait()

}

// type RateLimiter struct {
// 	capacity     int
// 	leakDuration time.Duration
// 	leakRate     int
// 	queue        chan int
// 	stop         chan struct{}
// 	wg           sync.WaitGroup
// }

// func NewRateLimiter(capacity int, leakDuration time.Duration, leakRate int) *RateLimiter {
// 	rl := &RateLimiter{
// 		capacity:     capacity,
// 		leakDuration: leakDuration,
// 		queue:        make(chan int, capacity),
// 		stop:         make(chan struct{}),
// 		leakRate:     leakRate,
// 	}
// 	rl.wg.Add(1)
// 	go rl.leakRequest()
// 	return rl
// }

// func (rl *RateLimiter) leakRequest() {

// 	defer rl.wg.Done()

// 	ticker := time.NewTicker(rl.leakDuration)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			count := 0
// 			for range rl.leakRate {
// 				select {
// 				case id := <-rl.queue:
// 					count++
// 					fmt.Printf("⏳ Processing request %d. Remaining: %d\n", id, len(rl.queue))
// 				default:
// 					fmt.Println("💤 Nothing to process...")
// 				}
// 			}
// 		case <-rl.stop:
// 			fmt.Println("Request leaking stopped")
// 			return
// 		}

// 	}

// }

// func (rl *RateLimiter) AllowRequest(id int) bool {
// 	select {
// 	case rl.queue <- id:
// 		fmt.Printf("✅ Request %d queued. Bucket fill: %d/%d\n", id, len(rl.queue), cap(rl.queue))
// 		return true
// 	default:
// 		fmt.Printf("❌ Request %d dropped. Bucket full!\n", id)
// 		return false
// 	}
// }

// func (lb *RateLimiter) Stop() {
// 	close(lb.stop)
// 	lb.wg.Wait()
// }

// func main() {
// 	lb := NewRateLimiter(5, 1*time.Second, 1)
// 	defer lb.Stop()

// 	requests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	for  {

// 		if len(requests) == 0
// 		go func(rID int) {
// 			lb.AllowRequest(rID)
// 		}(id)
// 		// time.Sleep(200 * time.Millisecond) // simulate incoming rate
// 	}

// 	// Allow some time for all requests to be processed
// 	// time.Sleep(7 * time.Second)
// }
