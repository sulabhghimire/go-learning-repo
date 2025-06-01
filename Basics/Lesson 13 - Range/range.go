package main

import "fmt"

func main() {

	// range keyboard provides a easy way to iterate over various data types in go like arrays, slices, maps and channels

	// Using range over string
	message := "Hello world"
	for i, v := range message { // Give index and rune
		fmt.Println("Index:", i, "Value:", v) // gives index and unicode value of the character in the string
	}

	// to get actual character
	for i, v := range message {
		fmt.Printf("Index: %d ,Value: %c\n", i, v)
	}

	// Imp: range keyboard iterates over the copy of data structure it iterates over so modifying the index or value inside
	// loop doesn't effect the original data structure

	// range iterates over channels until it is closed

}
