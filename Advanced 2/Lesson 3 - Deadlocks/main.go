package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(1 * time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
	}()

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mu1")
		time.Sleep(1 * time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 2 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
	}()

	time.Sleep(3 * time.Second)
	// select {}
	fmt.Println("Main func completed")

}
