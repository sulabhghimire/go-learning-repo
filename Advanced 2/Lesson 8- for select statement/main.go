package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(quit)
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Received a tick:")
		case <-quit:
			fmt.Println("Quitting...")
			return
		}
	}

}
