package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	pId := os.Getpid()
	fmt.Println("Process ID:", pId)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify channel on interrupt or terminate signals
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		sig := <-sigs
		fmt.Println("We received signal:", sig)
		done <- true
	}()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal.")
				//os.Exit(0)
				return
			default:
				fmt.Println("Working..")
				time.Sleep(time.Second)
			}
		}
		// for sig := range sigs {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("Received SIGINT (Interrupt)")
		// 	case syscall.SIGTERM:
		// 		fmt.Println("Received SIGTERM (Terminate)")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("Received SIGHUP (Hangup)")
		// 		// case syscall.SIGUSR1:
		// 		// 	fmt.Println("Received SIGNUSR1 (User defined signal 1)")
		// 		// 	fmt.Println("User defined function is executed")
		// 		// continue
		// 	}
		// 	fmt.Println("Graceful exit.")
		// 	os.Exit(0)
		//}
	}()

	// Simulate some work
	fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}
