package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	for i, arg := range os.Args {
		fmt.Println("Argument", i, ":", arg)
	}

	// Define flags
	var name string
	var age int
	var isMale bool

	// pointer to variable, name of flag, default value of flag, string that descirbe usage of flag
	flag.StringVar(&name, "name", "Sulabh", "Name of the user")
	flag.IntVar(&age, "age", 24, "Age of the user")
	flag.BoolVar(&isMale, "isMale", true, "Is the user male")

	flag.Parse()

	fmt.Println("Name:", name, "Age:", age, "Is User Male", isMale)

}
