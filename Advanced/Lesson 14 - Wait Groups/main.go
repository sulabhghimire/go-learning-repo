package main

import (
	"fmt"
	"sync"
	"time"
)

// // === CONSTRUCTION EXAMPLE
type Worker struct {
	ID   int
	Task string
}

// Perform Task will simulates a worker performing task
func (w *Worker) PerformTask(wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker ID %d started %s task.\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("Worker ID %d finished %s task.\n", w.ID, w.Task)

}

func main() {

	var wg sync.WaitGroup

	tasks := []string{"Digging", "Laying Bricks", "Painting"}

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	wg.Wait()
}

// // === EXAMPLE WITH CHANNELS
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	fmt.Printf("Worker %d is starting.\n", id)

// 	totalCompletedTasks := 0
// 	for task := range tasks {
// 		fmt.Printf("Worker %d is working on task %d.\n", id, task)
// 		time.Sleep(3 * time.Second)
// 		results <- task * 2
// 		fmt.Printf("Worker %d is completed on task %d.\n", id, task)
// 		totalCompletedTasks++
// 	}
// 	fmt.Printf("Worker %d has completed.\n", id)

// 	fmt.Printf("Worker %d has completed a total of %d tasks.\n", id, totalCompletedTasks)
// }

// func main() {

// 	start := time.Now()

// 	var wg sync.WaitGroup
// 	numWorkers := 200
// 	numJobs := 200
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i + 1
// 	}
// 	close(tasks)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for res := range results {
// 		fmt.Println("Result:", res)
// 	}

// 	end := time.Now()

// 	fmt.Println("Total Time:", end.Sub(start).Seconds())

// }

// // ==== BASIC EXAMPLE WITHOUT USING CHANNELS
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	fmt.Printf("Worker %d starting.\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished.\n", id)

// }

// func main() {

// 	// Create new wait group
// 	wg := new(sync.WaitGroup)
// 	numberOfWorkers := 3

// 	// Add number of workers in wait group
// 	wg.Add(numberOfWorkers)

// 	for i := range numberOfWorkers {
// 		go worker(i, wg)
// 	}

// 	// Wait for all the workers to complete
// 	wg.Wait()

// 	fmt.Println("All worker competed the task.")
// }
