package main

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(5))
	fmt.Println(sumOfDigits(56))
	fmt.Println(sumOfDigits(123879832))

}

func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}

	// Recursive case : factorial of n is n*factorail(n-1)
	return n * factorial(n-1)

}

func sumOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
