package main

import (
	"fmt"
	"math"
)

const PI = 3.14
const GRAVITY = 9.8 // untyped constant

func main() {
	const days int = 7 // typed constant

	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
		thrusday  = 4
	)

	// overflow with signed integers
	var maxInt int64 = 9223372036854775807 // maximum value for int64
	fmt.Println("maxInt: ", maxInt)        // prints: maxInt:  9223372036854775807
	maxInt++                               // this will cause an overflow error

	fmt.Println("maxInt: ", maxInt) // prints: maxInt:  -9223372036854775808 (overflow wraps around to negative)

	// overflow with unsigned integers
	var maxUint uint64 = 18446744073709551615 // maximum value for uint64
	fmt.Println("maxUint: ", maxUint)         // prints: maxUint:  18446744073709551615
	maxUint++                                 // this will cause an overflow error
	fmt.Println("maxUint: ", maxUint)         // prints: maxUint:  0 (overflow wraps around to zero)

	// underflow with small float
	var smallFloat float64 = 1.0e-323       // small float
	fmt.Println("smallFloat: ", smallFloat) // prints: smallFloat:  1e-323
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("smallFloat: ", smallFloat) // prints: smallFloat:  0 (underflow to zero)

}
