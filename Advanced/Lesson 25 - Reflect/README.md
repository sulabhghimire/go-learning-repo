# Go Reflection (`reflect` package)

Reflection is a mechanism that allows a program to inspect and manipulate its own structure and behavior at runtime. In Go, reflection is provided by the `reflect` package, which contains the types and functions necessary to work with arbitrary types dynamically.

The two most important types in the `reflect` package are `reflect.Type` and `reflect.Value`.

- **`reflect.Type`**: Represents a Go type. It provides metadata about the type, such as its name, kind (struct, int, slice, etc.), and methods.
- **`reflect.Value`**: Represents the actual value of a variable. It allows for inspection of the data and, in some cases, modification.

## Why Use Reflection?

Reflection is a powerful tool, but it should be used with care as it can make code more complex and slower. Common use cases include:

- **Dynamic Type Inspection**: Examining variables of an unknown type at runtime to determine their structure and properties.
- **Generic Programming**: Writing functions that can operate on values of any type, such as a generic data validator or object mapper.
- **Serialization/Deserialization**: Building custom encoders and decoders (like for JSON, XML, or ORMs) that can convert data structures to and from a serialized format without knowing their specific type at compile time.

## Core Functions and Concepts

| Function/Concept     | Description                                                                                                                                                                                                                    |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `reflect.TypeOf(i)`  | Takes any `interface{}` as a parameter and returns its dynamic type as a `reflect.Type`.                                                                                                                                       |
| `reflect.ValueOf(i)` | Takes any `interface{}` as a parameter and returns a `reflect.Value` representing the runtime data.                                                                                                                            |
| `Value.Kind()`       | Returns the underlying kind of a type (e.g., `reflect.Struct`, `reflect.Int`, `reflect.Ptr`). This is different from its `Type`, which is the specific user-defined type.                                                      |
| `Value.Elem()`       | Used on a `reflect.Value` that holds a pointer. It "dereferences" the pointer, returning a `Value` representing the data it points to. This is essential for modifying the original variable.                                  |
| **Settability**      | A `reflect.Value` is "settable" only if it is **addressable** and was created from an **exported** field. To modify a variable using reflection, you must pass a pointer to it to `reflect.ValueOf()` and then call `.Elem()`. |

---

## Examples

### 1. Inspecting Type, Kind, and Modifying Values

This example demonstrates the basics: getting the `Type` and `Kind` of a variable and modifying its value. To modify a value, we must pass a pointer to `reflect.ValueOf`.

```go
package main

import (
	"fmt"
	"reflect"
)

// UID is a custom type with an underlying kind of `int`.
type UID int

func main() {
	// --- Inspecting a Custom Type ---
	var x UID = 42

	v := reflect.ValueOf(x)
	t := v.Type()
	k := v.Kind()

	fmt.Println("--- Inspecting UID ---")
	fmt.Println("Value:", v) // Output: 42
	fmt.Println("Type:", t)  // Output: main.UID
	fmt.Println("Kind:", k)  // Output: int
	fmt.Println("Is kind 'int'?:", k == reflect.Int) // Output: true
	fmt.Println("Is Zero Value?:", reflect.ValueOf(UID(0)).IsZero()) // Output: true
	fmt.Println()

	// --- Modifying a Value ---
	// To modify a value, we need a pointer to it.
	y := 10
	v1 := reflect.ValueOf(&y).Elem() // Get the Value the pointer refers to

	fmt.Println("--- Modifying an int ---")
	fmt.Println("Is v1 settable?:", v1.CanSet()) // Output: true
	fmt.Println("Original value:", v1.Int())     // Output: 10

	// Set the value
	v1.SetInt(18)
	fmt.Println("Modified value:", y) // Output: 18
	fmt.Println()


	// --- Working with Interfaces ---
	var itf any = "Hello, Reflection!"
	v3 := reflect.ValueOf(itf)

	fmt.Println("--- Inspecting an Interface ---")
	fmt.Println("Interface Type:", v3.Type()) // Output: string
	fmt.Println("Interface Kind:", v3.Kind()) // Output: string
	if v3.Kind() == reflect.String {
		fmt.Println("String value:", v3.String()) // Output: Hello, Reflection!
	}
}
```

### 2. Working with Structs and Fields

Reflection is commonly used to iterate over the fields of a struct. Note that only **exported** fields (those starting with an uppercase letter) are accessible and modifiable.

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string // Exported field
	age  int    // Unexported field
}

func main() {
	p := Person{Name: "Alice", age: 30}

    // --- Reading Struct Fields ---
    // A copy of p is passed to ValueOf, so we can't modify it here.
	v := reflect.ValueOf(p)
	fmt.Println("--- Reading from a struct copy ---")
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		typeField := v.Type().Field(i)
		fmt.Printf("Field %d: Name='%s', Type='%v', Value='%v'\n", i, typeField.Name, field.Type(), field.Interface())
	}
	fmt.Println()

	// --- Modifying Struct Fields ---
	// Pass a pointer to make the struct settable.
	v1 := reflect.ValueOf(&p).Elem()
	fmt.Println("--- Modifying a struct via pointer ---")

	nameField := v1.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Jane")
	}

	ageField := v1.FieldByName("age")
	// The 'age' field is unexported, so it cannot be set.
	fmt.Printf("Can set 'Name' field?: %t\n", nameField.CanSet()) // true
	fmt.Printf("Can set 'age' field?: %t\n", ageField.CanSet())   // false

	fmt.Println("Modified Person:", p) // Output: {Jane 30}
}
```

### 3. Working with Methods

You can also use reflection to discover and call methods on a type. Arguments must be passed as a slice of `reflect.Value`.

```go
package main

import (
	"fmt"
	"reflect"
)

type Greeter struct{}

// Greet is an exported method.
func (g Greeter) Greet(fname, lname string) string {
	return "Hello, " + fname + " " + lname
}

func main() {
	g := Greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)

	fmt.Println("Type:", t)

	// --- Discover Methods ---
	fmt.Println("--- Discovering Methods ---")
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}
	fmt.Println()

	// --- Call a Method by Name ---
	fmt.Println("--- Calling a Method ---")
	m := v.MethodByName("Greet")

	// Prepare arguments for the call
	args := []reflect.Value{
		reflect.ValueOf("Alice"),
		reflect.ValueOf("Davidson"),
	}

	// Call the method
	results := m.Call(args)

	// Process results
	fmt.Println("Greet result:", results[0].String()) // Output: Hello, Alice Davidson
}
```
