package main

import "fmt"

func main() {

	fmt.Print("Hello", "a") // no space as space is added between two when neiter is a string
	fmt.Print(12, 456)      // space

	fmt.Println("Hello")
	fmt.Println(12, 456)

	name := "John"
	age := 25
	fmt.Printf("Name :%s, Age: %d\n", name, age)
	fmt.Printf("Binaray :%b, Hexadecimal: %X\n", age, age)

	// Fomratting Functions
	s := fmt.Sprint("Hello", "World!", 123, 456) // no space as space is added between two when neiter is a string and no new line character
	fmt.Println(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456) // space is added and new line is added
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age: %d", name, age) // doesn't add new line character
	fmt.Println(sf)
	fmt.Println(sf)

	// Scanning functions
	var name2 string
	var age2 int

	fmt.Print("Enter your name and age:")
	//fmt.Scan(&name2, &age2)   // Will keep asking for input
	//fmt.Scanln(&name2, &age2) // Stops scanning at a new line and tells use there be exactly one item per input
	fmt.Scanf("%s, %d", &name2, &age2)

	fmt.Println("Name is", name2, "and your age is", age2)

	// Error formatting Functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	}
	return nil
}
