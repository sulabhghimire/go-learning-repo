package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func initialize() {
	fmt.Println("This will not be repeated no matter how many times we call this function using once.Do")
}

func main() {

	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine #:", i+1)
			once.Do(initialize)
		}()
	}
	wg.Wait()

}
