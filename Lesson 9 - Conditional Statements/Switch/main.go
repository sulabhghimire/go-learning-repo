package main

import "fmt"

func main() {

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's and banand")
	default:
		fmt.Println("It's an unknown fruit")
	}

	// Multiple condition in switch case
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday":
		fmt.Println("It's a week day")
	case "Sunday", "Saturday":
		fmt.Println("It's weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Switch case with expression
	num := 15
	switch {
	case num < 10:
		fmt.Println("Number is less then 10")
	case num >= 10 && num < 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Number is 20 or more")
	}

	// Fallthrough in switch case
	number := 2
	switch {
	case number > 1:
		fmt.Println("Number greater less then 1")
		fallthrough
	case number == 2:
		fmt.Println("Number is two")
	default:
		fmt.Println("Not two")
	}

	checkType(10)
	checkType(2.14)
	checkType("Hello")
	checkType(true)

}

// Switch for type assertion
func checkType(x any) {
	switch x.(type) {
	case int:
		fmt.Println("It's integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's string")
	default:
		fmt.Println("Unknown type")
	}
}
