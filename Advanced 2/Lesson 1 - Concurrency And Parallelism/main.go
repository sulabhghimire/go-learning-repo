package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func printNumbers() {
// 	for i := range 5 {
// 		fmt.Println("Number", i, "at", time.Now())
// 		time.Sleep((500 * time.Millisecond))
// 	}
// }

// func printLetters() {
// 	for _, letter := range "ABCDE" {
// 		fmt.Println("Character", string(letter), "at", time.Now())
// 		time.Sleep((500 * time.Millisecond))
// 	}
// }

func heavyTask(id int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("Task %d is starting at %v\n", id, time.Now())
	for range 100_000_000 {
	}
	fmt.Printf("Task %d is finished at %v\n", id, time.Now())

}

func main() {
	// go printNumbers()
	// go printLetters()
	// time.Sleep(3 * time.Second)

	numThreads := 4
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	for i := range numThreads {
		wg.Add(1)
		go heavyTask(i+1, &wg)
	}

	wg.Wait()
}
