package main

import "fmt"

func main() {

	/*
		Simple iteration over a range
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
		}
	*/

	// iterate over collection
	// %v for general values
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	// Get odd numbers
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // continue the loop but skip rest of statements
	// 	}
	// 	fmt.Println("Odd number: ", i)
	// 	if i == 5 {
	// 		break // break out of the loop
	// 	}
	// }

	//Asterics Layou
	// rows := 5
	// // outer loop
	// for i := 1; i <= rows; i++ {
	// 	// inner loop for spaces before starts
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loop for starts
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// iteration over integer ranges
	// for i := range 10 {
	// 	i++
	// 	fmt.Println(i)
	// }

	// For loop as while loop
	// i := 1
	// for i <= 5 {
	// 	fmt.Println("Iteration:", i) // When we print two values seperated by comma, a space is added
	// 	i++
	// }

	// for as while with break
	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Sum:", sum)
	// 	if sum == 50 {
	// 		break
	// 	}
	// }

	// for as while with continue
	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd number:", num)
		num++ // ++ increment operator and -- decrement operator
	}

}
