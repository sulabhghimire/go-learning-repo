package main

import "time"

func main() {
	timer1 := time.NewTimer(1 * time.Second)
	defer timer1.Stop() // Ensure the timer is stopped to prevent resource leaks
	timer2 := time.NewTimer(2 * time.Second)
	defer timer2.Stop() // Ensure the timer is stopped to prevent resource leaks

	select {
	case <-timer1.C:
		println("Timer 1 expired")
	case <-timer2.C:
		println("Timer 2 expired")
	}

}

// // === Scheduling delayed execution
// func main() {

// 	timer := time.NewTimer(2 * time.Second)

// 	go func() {
// 		<-timer.C // Block until the timer expires
// 		println("Delayed operation executed")
// 	}()

// 	fmt.Println("Waiting...")
// 	time.Sleep(3 * time.Second) // blocking timer
// 	fmt.Println("Main function completed")
// }

// // ======= Time After
// func longRunningOperation() {

// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}

// }

// func main() {

// 	timeout := time.After(7 * time.Second)
// 	done := make(chan bool)

// 	go func() {
// 		longRunningOperation()
// 		done <- true
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation timed out")
// 	case <-done:
// 		fmt.Println("Operation completed successfully")

// 	}

// }

// // ======= Time Tick
// func main() {

// 	fmt.Println("Starting app")
// 	timer := time.NewTimer(5 * time.Second) // Non blocking unlike time.Sleep

// 	fmt.Println("Waiting for timer.C")

// 	stopped := timer.Stop() // Stop the timer before it expires
// 	if stopped {
// 		fmt.Println("Timer stopped before expiration")
// 	}

// 	fmt.Println("Timer reset")
// 	timer.Reset(time.Second) // Reset the timer to 1 second
// 	<-timer.C                // Block until the timer expires
// 	println("Timer expired after 2 seconds")

// }
