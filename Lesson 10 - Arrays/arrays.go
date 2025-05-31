package main

import "fmt"

func main() {

	//var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers) // initializes with default value 0,0,0,0,0

	// index starts with 0 so 1st elm indx 0 and last elm index len(array)-1
	numbers[4] = 30
	fmt.Println(numbers) // 0, 0, 0, 0, 30

	fruits := [4]string{"Apple", "Banana", "Mango", "Grapes"}
	fmt.Println("Fruits Array: ", fruits)

	// can access using index
	fmt.Println(fruits[2]) // getting third element of the array

	// In go arrays are value types when assigned to new variables or pass an array as an argument ot func copy of the original
	// array is created and changes to copy doesn't affect the original array

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[1] = 100

	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array:", copiedArray)

	// we can iteratate over array using a for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Num:", numbers[i])
	}

	// can also iterate using range
	for index, value := range numbers {
		fmt.Println("Range Index:", index, "Num:", value)

	}
	// we can use _ to discard a value also called blank identifier
	// eg : for _, index := range numbers
	// if we don't want to use any values returned from anywhere be ir range, or a func, we can assign _ to that value
	// and we don't get any errors

	// Also we can do if some needs may come up
	b := 2
	_ = b

	// Length of array using len func
	fmt.Println("The length of number array is :", len(numbers))

	// Comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Array 1 is equal to array 2 :", array1 == array2)

	// Array can be used to represent multidimensional data
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Maxtrix is:", matrix)

	originalArray2 := [3]int{1, 2, 3}
	var copiedArray2 *[3]int
	copiedArray2 = &originalArray2

	copiedArray2[1] = 100

	fmt.Println("Original Array: ", originalArray2)
	fmt.Println("Copied Array:", copiedArray2)

}
