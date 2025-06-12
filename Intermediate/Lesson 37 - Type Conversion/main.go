package main

import "fmt"

func main() {

	// Numeric type conversions
	var a int = 32
	b := int32(a)
	c := float64(b)

	e := 3.14
	f := int(e) // truncates the decimal part
	fmt.Println(f, c)

	// Type(value)

	g := "Hello"
	var h []byte // byte is unit8 0 to 255
	h = []byte(g)
	fmt.Println(h)

	i := []byte{255, 72, 'a', 'b'}
	fmt.Println(string(i))
}
