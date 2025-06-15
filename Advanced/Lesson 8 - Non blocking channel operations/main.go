package main

import "time"

func main() {

	// // === Non Blocking receive Operations
	// ch := make(chan int)
	// select {
	// case v := <-ch:
	// 	// This case will not block, but will not execute since the channel is empty
	// 	println("Received:", v)
	// default:
	// 	// This case will execute since the channel is empty
	// 	println("No value received, channel is empty")
	// }

	// // === Non Blocking send Operations
	// select {
	// case ch <- 42:
	// 	// This case will not block, but will not execute since the channel is unbuffered and no receiver is ready
	// 	println("Sent value 42 to channel")
	// default:
	// 	// This case will execute since the channel is unbuffered and no receiver is ready
	// 	println("No value sent, channel is unbuffered and no receiver")
	// }

	// === Non Blocking Operations in Real Time Systems

	data := make(chan int, 1)
	quit := make(chan bool)

	go func() {

		for {
			select {
			case v := <-data:
				println("Data Received:", v)
			case <-quit:
				println("Quit signal received, exiting goroutine")
				return
			default:
				// No data received, continue processing
				println("No data received, continuing processing")
				time.Sleep(500 * time.Millisecond)
			}
		}

	}()

	for i := range 5 {
		data <- i
		time.Sleep(1 * time.Second)
	}

	quit <- true
}
