package main

func main() {

	// var age int              // Declare an uninitilized variable named age of type int
	// var name string = "John" // Declare and initialize a variable named name of type string
	// var name1 = "Doe"        // Declare and initialize a variable named name1 of type string with type inference

	// count := 10         // Declare and initialize a variable named count of type int with type inference and with walrus operator
	// lastName := "Smith" // Declare and initialize a variable named lastName of type string with type inference and with walrus operator

	/*
		Default value
		Numeric Types: 0
		String Type: ""
		Boolean Type: false
		Pointers, Functions, Maps, Struct  and Slices: nil
	*/

	/*
		Scope
			Variables in go has block scope.

		We can decalre variables at package level or inside a function.
		Package level variables are accessible to all functions in the package.
		Variables declared inside a function are only accessible within that function.
		Variables declared inside a block (e.g. if, for, switch) are only accessible within that block.
		Variables declared inside a loop are only accessible within that loop.

		:= ie. the walrus operator can only be used inside a function.
		We need to use the var keyword to declare variables at package level.

		The lifetime of varaible is with in therir scope.
		They are created when their scope starts and destroyed when the scope ends.

	*/

}
