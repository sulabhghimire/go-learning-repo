package main

import (
	"fmt"
	"reflect"
)

// Working with Methods
type Greeter struct{}

func (Greeter) Greet(fname, lname string) string {

	return "Hello " + fname + " " + lname

}

func main() {

	g := Greeter{}

	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)

	var method reflect.Method
	fmt.Println("Type:", t)
	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("Method: %d: %s\n", i, method.Name)
	}

	m := v.MethodByName(method.Name)
	results := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("Davidson")})
	fmt.Println("Greet result:", results[0].String())

}

// // ======== WORKING WITH STRUCT AND FIELDS
// type person struct {
// 	Name string
// 	age  int
// }

// func main() {
// 	p := person{Name: "Alice", age: 30}
// 	v := reflect.ValueOf(p)

// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}

// 	v1 := reflect.ValueOf(&p).Elem()
// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet() { // Only exported fields can be accessed and modified
// 		nameField.SetString("Jane")
// 	} else {
// 		fmt.Println("Can't set")
// 	}

// 	fmt.Println("Modified Person:", p)
// }

// type UID int

// func main() {

// 	var x UID
// 	x = 43

// 	v := reflect.ValueOf(x)
// 	t := v.Type()
// 	k := v.Kind()

// 	fmt.Println("Value:", v)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", k)
// 	fmt.Println("Is int:", k == reflect.Int)
// 	fmt.Println("Is string:", k == reflect.String)
// 	fmt.Println("Is Zero:", v.IsZero())

// 	y := 10
// 	v1 := reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)
// 	fmt.Println("V2 Type:", v2.Type())

// 	fmt.Println("Original value:", v1.Int())

// 	v1.SetInt(18)
// 	fmt.Println("Modified value:", v1.Int())

// 	var itf any = "Hello"
// 	v3 := reflect.ValueOf(itf)

// 	fmt.Println("V3 Type:", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println("String value:", v3.String())
// 	}
// }
