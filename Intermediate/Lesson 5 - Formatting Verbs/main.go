package main

import "fmt"

func main() {

	i := 111_505.5
	string := "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)

	fmt.Println("Integer formatting")
	int := 45
	fmt.Printf("%b\n", int)
	fmt.Printf("%d\n", int)
	fmt.Printf("%+d\n", int)
	fmt.Printf("%o\n", int)
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int)
	fmt.Printf("%X\n", int)
	fmt.Printf("%#x\n", int)
	fmt.Printf("%#X\n", int)
	fmt.Printf("%4d\n", int)
	fmt.Printf("%-4d\n", int)
	fmt.Printf("%04d\n", int)

	fmt.Println("String formatter ----------------")
	txt := "World"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	fmt.Println("Boolean formatter ----------------")
	t := true
	f := false
	fmt.Printf("%t\n", t)
	fmt.Printf("%v\n", f)

	fmt.Println("Floating point formatter ----------------")
	flt := 9.15
	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)

}
