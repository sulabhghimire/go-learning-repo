package main

import "fmt"

func main() {

	// Recover is a built in function that is used to regain the control of panicing go routine
	// Its only useful inside defer functions and used to manage behaviour of panicing go routine of abrubt behaviour
	process()
	fmt.Println("Returned from process")
}

/*
 Uses of Recover
 a. Graceful Recovery
 b. Cleanup
 c. Logging and Reporting

 Best Practices
 a. Use with defer
 b. Avoid silent recovery
 c. Avoid overuse
*/

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start process")
	panic("Something went wrong")
}
