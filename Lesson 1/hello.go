package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}

// run using --> go run filename --> must be main package and contain a main func
// build --> go build filename --> builds an executable file
// run file with ./fielname
// with go build we are making a presistent executable file but with go run we aren't making a presistent executable file

// what happens when we run our program with go run
// --> go run is used to compile and execute a go porgram in a single step without explicitly creating a binary executable
// --> first compiles out source code .go files into a temporary executable file in memory (RAM)
// --> after sucessfull compilation go immediately runs the compiled executable
// --> convinent for quickly running and testing small programs
// --> useful during developmenet and testing phases
// --> after program completes the temporary executable created by program is discarded
