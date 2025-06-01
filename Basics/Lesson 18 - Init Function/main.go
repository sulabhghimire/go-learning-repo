package main

import "fmt"

// The init function is a special function that can be declared within any package
// Used to peform initilization task in a package before it is used
// The init function has no parameters and no return values
// Go executes init function automatically when the pacakge is initilized and happens before main function is executed
// And it occurs exactly once per package even if the package is imported multiple times

// Order of Execution
// Within a single package go exacutes the init function in the order which the packages are declared

/*
Used for tasks like initilizing varaibles, performing set of operations, registrating compoments or configurations and
initilizating state required for the package to function correctly
*/

func init() {
	fmt.Println("Initilizing package1...")
}

func init() {
	fmt.Println("Initilizing package2...")
}

func init() {
	fmt.Println("Initilizing package3...")
}

func main() {
	fmt.Println("Inside the main function")
}

/*
Practical Use Cases
a. Setup Task
b. Configuration
c. Registering Components
d. Database Initialization

Best Practices
a. Avoid Side Effects
b. Initialization Order
c. Documentation
*/
