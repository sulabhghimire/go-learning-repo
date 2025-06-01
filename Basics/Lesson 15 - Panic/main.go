package main

import "fmt"

func main() {

	// panic
	// panic is aa built in func that stops normal execution of a function immediately.
	// when a function encounters a panic it stops executing its current activities, unwinds the stack and executes any deferred functions
	// this process continues up the stack until all the functions have returned and program terminates
	// used to indicate any un-expected error condition where the program can't continue safely
	// anthing after panic will not be executed by runtime

	// panic function is called with optional argument of any type -> that represents the value associated with the panic
	// Syntax: panic(interface{}) OR panic(any)

	// valid input
	process(102)

	// Invalid input
	process(-10)

}

func process(input int) {

	defer fmt.Println("Deferred One")
	defer fmt.Println("Deferred Two")

	if input < 0 {
		panic("Input must be an non-negative number")
	}

	fmt.Println("Processing input :", input)

}
