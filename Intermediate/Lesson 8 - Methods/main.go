package main

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with value reciever
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer reciever --> used to change type or to copy large structs
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

// Method on user defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string { // not acessing the value of myInt
	return "Welcome to MYiNT TYPE"
}

func main() {

	react := Rectangle{length: 10, width: 9}
	area := react.Area()
	fmt.Println("Area is :", area)

	react.Scale(2)
	area = react.Area()
	fmt.Println("Area after scale is :", area)

	num := MyInt(-5)
	num2 := MyInt(7)
	fmt.Println("Is num positive", num.IsPositive())
	fmt.Println("Is num2 positive", num2.IsPositive())
	fmt.Println(num.welcomeMessage())

	// Struct embedding allows the methods of any embedded struct to be promoted to the outer struct
	shape := Shape{
		Rectangle: Rectangle{length: 10, width: 9},
	}
	fmt.Println(shape.Area()) // since we are embedding rectangle in shape so method associated with rectangle will promoted to Shape
	fmt.Println(shape.Rectangle.Area())
}
