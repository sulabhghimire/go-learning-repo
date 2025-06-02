package main

import "fmt"

func main() {

	// var ptr *int

	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(ptr)  // stores the memory address where a is stored in the RAM
	fmt.Println(*ptr) // dereferencing a pointer value

	// zero value of a pointer is nil and doesn't store memory address of any varaible

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
