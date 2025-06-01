# Go Variable Declarations and Scope

This Go program demonstrates how to declare variables using different techniques, explains default values, and describes variable scoping rules in Go.

## Variable Declaration

- Variables can be declared using the `var` keyword with or without initialization.
  Example: var age int  
   var name string = "John"

- Go supports type inference, allowing you to omit explicit types when initializing.
  Example: var name1 = "Doe"

- Inside functions, the short declaration operator (`:=`) can be used for both declaring and initializing variables.
  Example: count := 10  
   lastName := "Smith"

## Default Values

When a variable is declared but not initialized, Go assigns it a default zero value:

- Numeric types: 0
- String: empty string ""
- Boolean: false
- Pointers, functions, maps, structs, and slices: nil

## Variable Scope

- Go variables have block-level scope.

- Variables can be declared at the package level or inside functions.

- Package-level variables are accessible to all functions in the package.

- Variables declared inside a function are only accessible within that function.

- Variables declared inside a block (e.g., if, for, switch) are only accessible within that block.

- Variables declared inside a loop are only accessible within that loop.

- The short declaration operator (`:=`) can only be used inside functions.

- To declare variables at the package level, use the `var` keyword.

- The lifetime of a variable is limited to its scope: it is created when the scope begins and destroyed when it ends.
