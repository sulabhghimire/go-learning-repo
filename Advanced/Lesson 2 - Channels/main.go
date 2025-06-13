package main

import "fmt"

// This is mistake as we can't make it work inside function
/*
Get fatal error all go routines are deadlock as we are not using any goroutines
As channels in Go are blocking
*/
func IntentionalMistake() {
	greeting := make(chan string)
	greetingString := "Hello"

	greeting <- greetingString // blocking because it is continuously trying to recieve values, it is ready to receive continuous flow of data.

	reciever := <-greeting

	fmt.Println("Channel greeting:", reciever)
}

func main() {

	// variable := make(chan type)
	// IntentionalMistake()
	greeting := make(chan string)
	greetingString := "Hello"

	go func() {
		greeting <- greetingString
		greeting <- "Random"

		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}

	}()

	reciever := <-greeting
	fmt.Println("Channel greeting:", reciever)

	reciever = <-greeting
	fmt.Println("Channel greeting second:", reciever)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	anotherTest() // Here as soon as both are executed the function returns and doesn't wait for the process to complete
}

func anotherTest() {
	greeting := make(chan string)
	greetingString := "Hello"

	go func() {
		greeting <- greetingString
	}()

	go func() {
		reciever := <-greeting
		fmt.Println("Channel greeting:", reciever)
	}()

}
