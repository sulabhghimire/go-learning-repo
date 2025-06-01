package main

import (
	"fmt"
)

func main() {

	// function decleration
	// func <name>(parameters list) (returnTypes) {
	//	return
	//}

	// func name should be valid identifier
	// public func should start with capital letter and private with small letter

	// zero or more parameters can be defined following their type
	// we can return multiple values from a function or don't return any values
	// if return statement is omitted then functions return default zero values for their return types

	// Arguments passed into the functions are copied into the function parameters
	// Modifications of the parameters inside the functions don't affect the original arguument passed

	sum := add(1, 2)
	fmt.Println("Sum is :", sum)

	// anynomous functions: defined without name directly where they are called
	func() {
		fmt.Println("Hello anynomous function")
	}()
	// this is called immediately

	greet := func() {
		fmt.Println("Hello anynomous function greet")

	}

	greet()

	// function can be used as types and the functions in go can be assigned to variables and
	// passed as arguments af functions and retruns as funcitons
	operation := add
	result := operation(2, 5)
	println(result)

	// Passing as an argument here
	resApplyOperation := applyOperation(5, 4, add)
	println("Apply operation reuslt:", resApplyOperation)

	// Returning a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))

	// So functions are called first class citizens in go

	// Functions in go can have multiple return values
	// func functionName(parameter1 type1, parameter2 type2, ....) (returnType1, returnType2, ...) {
	// 	return returnValue1, returnValue2, ...
	// }

	q, r := divide(2, 7)
	fmt.Println("Q:", q, "R:", r)

	// Variadic functions in go can accept a variable number of arguments
	// ... ellipsis indicates that func can accept zero or more arguments of that typr
	// func funcName(param1 type1, param2 type2, param3 ...type3) returnType
	// param3 is a varadic parameters --> can be zero or more parameters for type 3

	// fmt.Println("Variadic function : ", sumOfIntegers(1, 2, 3, 4))

	statement, total := sumOfIntegers("Sum of 1-10 is:", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(statement, total)

	// using slices as a variadic parameters
	numbers := []int{1, 2, 3, 4, 5}
	statement, total = sumOfIntegers("Sum of 1-5 is:", numbers...)
	fmt.Println(statement, total)

	// defer is a meachanism that allows us to postpond the execution of a function until the surronding function returns
	// it helps to ensure that certain cleanup actions or finalizing tasks are peformed regardless how functions exists
	// weather it returns normally or panics
	// defer statement is followed by function or method call, the function is evaluated immediately but the
	//  execution is defered until the surrounding func returns
	process()

	// we can have multiple defer statemts in a single func and will be executed in last in first out order when func returns
}

/*
Use case of defer
1. Resource cleanup like file or database connections are closed
2. While using mutexes to synchronize go routine --> to enusre mutxes are unlocked even if functions panics
3. For logging and tracing

Best practices
1. Keep deferred actions short
2. Understand the evaluation timing
3. Avoid complex control flow
*/

func process() {
	defer fmt.Println("Deferring the statement first")
	defer fmt.Println("Deferring the statement second")
	fmt.Println("Processing normal executing")
	return
}

func sumOfIntegers(returnString string, nums ...int) (string, int) {

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return returnString, sum

}

// func sumOfIntegers(nums ...int) (sum int) {

// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return

// }

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func add(a, b int) int {
	return a + b
}

// a function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// a function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
