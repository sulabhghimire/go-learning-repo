package main

// Pipeline pattern in Go using channels
func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch) // Close the channel when done sending
}

func filter(in <-chan int, out chan<- int) {

	for val := range in {
		if val%2 == 0 { // Filter even numbers
			out <- val
		}
	}
	close(out) // Close the output channel when done sending

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		println(val) // Print the filtered values
	}
}

// Pipeline pattern ends here

// // Closing a channel more than once will cause a panic
// func main() {

// 	ch := make(chan int)
// 	go func() {
// 		close(ch)
// 		close(ch) // Multiple close calls on channel cause the program to panic
// 	}()
// 	time.Sleep(1 * time.Second)

// }

// // RANGE OVER A CLOSED CHANNEL
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch) // if not closed, the it will error
// 	}()

// 	for val := range ch {
// 		println(val)
// 	}

// }

// // Receiving from a closed channel
// func main() {

// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <-ch
// 	if !ok {
// 		println("Channel is closed, no value received")
// 	} else {
// 		println("Received value:", val)
// 	}

// }

// // Simple example of closing a channel in Go
// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		println(val)
// 	}

// }
