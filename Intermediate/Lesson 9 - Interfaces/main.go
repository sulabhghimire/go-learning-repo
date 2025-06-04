package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type cicle struct {
	radius float64
}

// Rectangle
// rect struct area method
func (r rect) area() float64 {
	return r.height * r.width
}

// rect struct perimeter method
func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// Circle
// circle struct area method
func (c cicle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// circle struct perimeter method
func (c cicle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// circle struct diameter method --> not part of geometry interface
func (c cicle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {

	r := rect{width: 2, height: 4}
	c := cicle{radius: 5}

	measure(r)
	measure(c)

}
