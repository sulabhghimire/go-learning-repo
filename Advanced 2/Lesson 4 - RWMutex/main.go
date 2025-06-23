package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu    sync.RWMutex
	counter int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	rwmu.RLock()
	defer rwmu.RUnlock()

	fmt.Println("Read Counter:", counter)
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()

	rwmu.Lock()
	defer rwmu.Unlock()

	counter = value
	fmt.Printf("Written value %d for counter\n", value)
}

func main() {

	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	time.Sleep(1 * time.Microsecond)
	go writeCounter(&wg, 18)

	wg.Wait()

}
