package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go producer(ch)
	consumer(ch)
}

func producer(ch chan<- int) {
	for i := range 5 {
		time.Sleep(2 * time.Second)
		ch <- i
	}
	close(ch)
}

// Receive only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value, ":", time.Now())
	}
}
