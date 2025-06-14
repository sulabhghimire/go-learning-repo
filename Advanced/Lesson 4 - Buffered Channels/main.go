package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// variable := make(chan type, capacity int)
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	// ch <- 3 // error deadlock --> no more space in the channel
// 	fmt.Println("Buffered channels")
// 	fmt.Println("Value1", <-ch)
// 	fmt.Println("Value2", <-ch)
// 	ch <- 3
// 	fmt.Println("Buffered channels")

// 	// no error as doesn't need immediate receiver

// }

func main() {
	// ==Blocking on send only if the buffer is full
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from Buffer")
	go func() {
		fmt.Println("Go routine 2 sec timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch) // ends <----- starts
	}()
	fmt.Println("Blocking Starts")
	ch <- 3 // Blocks because buffer is full
	fmt.Println("Blocking Ends")
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)

	fmt.Println("Buffered Channels")

}

// func main() {
// 	// ==== BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()

// 	fmt.Println("Value:", <-ch)
// 	fmt.Println("Value:", <-ch)
// 	fmt.Println("End of Program")
// }
