package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numWorkers := 5
	wg.Add(numWorkers)

	increment := func() {
		defer wg.Done()
		for range 100 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numWorkers {
		go increment()
	}

	wg.Wait()

	fmt.Println("Final counter", counter)
}

// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	c.count++
// }

// func (c *counter) getValue() int {
// 	return c.count
// }

// func main() {

// 	var wg sync.WaitGroup
// 	counter := &counter{}

// 	numWorkers := 10

// 	for range numWorkers {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				counter.increment()
// 			}
// 		}()
// 	}

// 	wg.Wait()

// 	fmt.Println("Total:", counter.getValue())

// }
