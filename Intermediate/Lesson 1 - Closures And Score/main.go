package main

import "fmt"

func main() {

	sequence := adder("a") // as called once so i will be 0 only once

	fmt.Println(sequence()) // executes the returned function when called i -> 1
	fmt.Println(sequence()) // executes the returned function when called i -> 2
	fmt.Println(sequence()) // executes the returned function when called i -> 3
	fmt.Println(sequence()) // executes the returned function when called i -> 4

	sequence2 := adder("b")
	fmt.Println(sequence2()) // executes the returned function when called i -> 1
	fmt.Println(sequence2()) // executes the returned function when called i -> 2
	fmt.Println(sequence2()) // executes the returned function when called i -> 3
	fmt.Println(sequence2()) // executes the returned function when called i -> 4

	subtractor := func() func(int) int {
		countDown := 99
		return func(i int) int {
			countDown -= i
			return countDown
		}
	}()

	// Using the closure subtractor
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))
	fmt.Println(subtractor(4))

}

func adder(a string) func() int {
	i := 0 // the variable i is scoped to the adder function --> will execute each time adder function is called
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added one to i:", i, "in sequence", a)
		return i
	}
}
