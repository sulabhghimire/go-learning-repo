package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 1

type buffer struct {
	items []int
	mu    sync.RWMutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
		// cond: sync.NewCond(&b.mu), Can't do here because we need instance of b and here instance is not created
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == bufferSize {
		b.cond.Wait()
		// releases the mutex temporarily
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.cond.Wait()
	}
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item)
	b.cond.Signal()

	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(time.Second)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 10 {
		b.consume()
		time.Sleep(2 * time.Second)
	}
}

func main() {
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()

}
