package main

import (
	"fmt"
	"time"
)

func main() {

	var err error

	fmt.Println("Begining program")
	go sayHello()
	fmt.Println("After sayHello function")

	// err = go doWork() is not acceptes
	go func() {
		err = doWork()
	}()

	go printNumbers()
	go printLettters()

	time.Sleep(2 * time.Second) // WAIT FOR GO ROUTINE TO FINISH

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully.")
	}

}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := range 5 {
		fmt.Println("Number:", i, "Time:", time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLettters() {
	for _, letter := range "abcde" {
		fmt.Println("Letter:", string(letter), "Time:", time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate the work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("Error occured in doWork")
}
