package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// Simulate the processing of ticket requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d tickets for person %d with cost %d\n", req.numTickets, req.personID, req.cost)
		// Simulate some processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func main() {
	numRequests := 5
	price := 5
	ticketsRequests := make(chan ticketRequest, numRequests)
	ticketsResults := make(chan int)

	// Start the ticket processor
	for range 3 {
		go ticketProcessor(ticketsRequests, ticketsResults)
	}

	// Send ticket requests to the channel
	for i := range numRequests {
		ticketsRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketsRequests) // Close the channel to signal no more requests

	for range numRequests {
		fmt.Printf("Ticker for personID %d processed successfully\n", <-ticketsResults)
	}
}

// //=== BASIC WORKER POOL EXAMPLE ===//
// func worker(id int, tasks <-chan int, results chan<- int) {

// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		// Simulate some work
// 		time.Sleep(time.Second)
// 		results <- task * 2 // Example processing: double the task value
// 	}

// }

// func main() {

// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)

// 	results := make(chan int, numJobs)

// 	// Create Workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	// Send tasks to the tasks channels for workers
// 	for i := range numJobs {
// 		tasks <- i
// 	}
// 	close(tasks) // Close the tasks channel to signal no more tasks

// 	// Collect results
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}
// }
