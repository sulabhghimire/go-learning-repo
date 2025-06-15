package main

import "fmt"

// If we are waiting for only one channel response
// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {

// 		time.Sleep(time.Second)
// 		ch1 <- 1

// 	}()

// 	go func() {

// 		time.Sleep(time.Second)
// 		ch2 <- 2

// 	}()

// 	// Wait when used default because the default case will be executed without waiting for goroutine to finish
// 	// Not needed if do default case in select
// 	time.Sleep(2 * time.Second)

// 	select {
// 	case msg := <-ch1:
// 		fmt.Println("Received from Channel 1:", msg)

// 	case msg := <-ch2:
// 		fmt.Println("Received from Channel 2:", msg)

// 	default:
// 		fmt.Println("No channels ready")
// 	}

// }

// Receiving for all channels
// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {

// 		time.Sleep(time.Second)
// 		ch1 <- 1

// 	}()

// 	go func() {

// 		time.Sleep(time.Second)
// 		ch2 <- 2

// 	}()

// 	for range 2 {
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("Received from Channel 1:", msg)

// 		case msg := <-ch2:
// 			fmt.Println("Received from Channel 2:", msg)
// 		}
// 	}

// }

// Use select with timeout to handle operations that take too long
// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Received data:", msg)
// 	case <-time.Tick(1 * time.Second):
// 		fmt.Println("Time out")
// 	}

// }

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

Loop:
	for {
		select {
		case msg, ok := <-ch: // For ok to be false the channel needs to be closed and empty
			if !ok {
				fmt.Println("Channel closed")
				break Loop
			}
			fmt.Println("Received data:", msg)
		}
	}
	fmt.Println("Operation completed:")

}
