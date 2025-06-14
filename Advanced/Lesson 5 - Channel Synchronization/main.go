package main

import (
	"fmt"
	"time"
)

// func main() {

// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	<-done

// 	fmt.Println("Finished.")

// }

// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		ch <- 102 // Blocking until value is received
// 		fmt.Println("sent value")
// 	}()
// 	value := <-ch // Blocking until a value is sent
// 	fmt.Println(value)
// }

// ================= SYNCHRONIZING MULTIPLE GO ROUTINES AND ENSURING ALL GOROUTINES ARE COMPLETE
// func main() {
// 	numGroupRoutines := 3
// 	done := make(chan int, 3)

// 	for id := range numGroupRoutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working ...\n", id)
// 			time.Sleep(time.Second)
// 			done <- id // Sending signal of completion
// 		}(id)
// 	}

// 	for range numGroupRoutines {
// 		<-done // Wait for each goroutine to finish, WAIT FOR ALL GOROUTINES TO SIGNAL COMPLETION
// 	}

// 	fmt.Println("All go routines are finished")
// }

// ==============SYNCHRONIZING DATA EXCHANGE
func main() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- fmt.Sprintf("Hello %d", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data) // Close so no errors
	}()
	// close(data) // Channel closed before Goroutine could send a value ot the channel

	// When we loop over the channel we create a receiver each time
	// As long as the channel is open the for loop is continuously looping over the channel to receive values
	// We get values because we loop over last time when there is no value in channel
	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	} // loops over only on active channel, creates receiver each time and stops creating receiver(looping) once the cannel is closed

}
