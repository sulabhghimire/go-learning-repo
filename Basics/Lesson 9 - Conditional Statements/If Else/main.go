package main

import "fmt"

func main() {

	age := 45
	if age >= 18 {
		fmt.Println("You are an adult")
	}

	temp := 25
	if temp >= 30 {
		fmt.Println("It's hot outside")
	} else {
		fmt.Println("It's cool outside")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade : A")
	} else if score >= 80 {
		fmt.Println("Grade : B")
	} else if score >= 70 {
		fmt.Println("Grade : C")
	} else {
		fmt.Println("Grade : D")
	}

}
