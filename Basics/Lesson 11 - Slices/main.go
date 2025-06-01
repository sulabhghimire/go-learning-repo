package main

import (
	"fmt"
	"slices"
)

func main() {

	// var sliceName[]ElementType declare slices
	// different ways of declearing and initiliazing array

	// var numbers []int
	// var number1 []int = []int{1, 2, 3, 4}

	// numbers3 := []int{9, 8, 7, 6}

	// Slices are refernece to the underlying arrays.
	// Don't store any data themselves but provide an window into array's elements
	// can grow an shrink dynamically
	// same len function can be used to check the length of a slice
	// cap function can be used to check capacity of an array
	// checks the number of elements in underlying array starting from slice's first element

	// we can also initialize slice using make function
	// slice := make([]int, 5)

	// we can also make slice from existing arry
	a := [5]int{1, 2, 10, 12, 5}
	slice := a[1:4] // make a slice starting from index 1 to 4; including index 1 but not index 4

	fmt.Println(slice)

	// we can also add to slice with append function
	slice = append(slice, 6, 7)
	fmt.Println(slice)

	// Copying a slice
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Println(sliceCopy)

	// nil slices
	// nil slice has zero values and doesn't reference any underlying array

	var nilSlice []int
	fmt.Println(nilSlice)

	// we can loop over slices llike as in array
	// we can also access the element of slice like as in array
	// we can also modify elemtn of slick like as in array

	// We ccn compare slice equality via
	if slices.Equal(slice, sliceCopy) {
		fmt.Println("Slice 1 is equal to slice copy")
	}

	// We can also make multidimensional slices with slices as well
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := range innerLength {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// Slice also support slice operator
	// slice[low:high]
	sliceA := slice[2:4]
	fmt.Println(sliceA)
	sliceB := slice[:]
	fmt.Println(sliceB)

	// Capacity of slice
	fmt.Println("Capacity of slice is :", cap(slice))
	fmt.Println("Length of slice is :", len(slice))

}
