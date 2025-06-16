package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop() // Ensure the ticker is stopped when done

	stop := time.After(5 * time.Second)
Loop:
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at", tick)
		case <-stop:
			fmt.Println("Stopping ticker after 5 seconds")
			break Loop
		}
	}

}

// // =========SCHEDULING TASKS WITH TICKERS========
// func periodicTask() {
// 	fmt.Println("Periodic task executed at", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop() // Ensure the ticker is stopped when done

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// func main() {

// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop() // Ensure the ticker is stopped when done

// 	i := 0
// 	for tick := range ticker.C {
// 		i++
// 		fmt.Println("Tick", i, "at", tick)
// 		if i == 9 {
// 			ticker.Stop() // Stop the ticker after 9 ticks
// 			break
// 		}
// 	}

// }
