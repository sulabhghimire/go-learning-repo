package main

import (
	"fmt"
	"os"
)

func main() {

	// In go os.Exit(statusCode) is a function that terminates the program immediately with the given status code
	// Useful in situation where we need to stop the execution of the program immdediately without deferring any function
	// or peforming any cleanup operations
	// It stops the program immediately and any deferred functions will not be executed
	// Function accepts and integer argument that returns the statusCode to the os
	// 0 indicates successfull completion
	// non-zero errors or abnormal exit

	defer fmt.Println("Deferred statement")

	fmt.Println("Starting the main function")

	// Exit with status code 1
	os.Exit(1)

	// This will never be executed
	fmt.Println("End of the main function")

}

/*

Practical Use Cases
a. Error Handaling
b. Termination Conditions
c. Exit Codes

Best Practices
a. Avoid Deferred Actions
b. Status Codes
c. Avoid Abusive Use

*/
