package main

import (
	"fmt"
	"os"
	"strings"
)

// Environment variables are the key value pairs that are part of the environment in which process runs
// Porvide a convivent way to pass configuration information, credentials and other parameters to applications without
// hardcoding those values into the application itself

// In Go environmental variables are accessed through os package

func main() {

	// Get environment variable
	user := os.Getenv("USER") // Works on unix like
	home := os.Getenv("HOME") // Works on unix like

	fmt.Println("User env is:", user)
	fmt.Println("Home env is:", home)

	fmt.Println("-------------Setting And Getting Envrionment Variables--------------------")
	// Set environment variable
	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
	}
	fmt.Println("FRUIT:", os.Getenv("FRUIT"))

	fmt.Println("-----------------------Using os.Environ------------------------------------")
	// Get all environmental variables
	for _, env := range os.Environ() {
		kvpair := strings.SplitN(env, "=", 2)
		fmt.Println(kvpair[0])
	}

	fmt.Println("-----------------Unsetting an environment variable-------------------------")
	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error un-setting environment variable:", err)
	}
	fmt.Println("FRUIT:", os.Getenv("FRUIT"))

	/*
		"a=b=c=d"
		n=1 returns "a=b=c=d"
		n=2 returns "a" and "b=c=d"
		n=3 returns "a" and "b" and "c=d"
		n=3 returns "a" and "b" and "c" and "d"
	*/
	str := "a=b=c=d=e"
	fmt.Println(strings.SplitN(str, "=", -1)) // [a b c d e]
	fmt.Println(strings.SplitN(str, "=", 0))  // []
	fmt.Println(strings.SplitN(str, "=", 1))  // [a=b=c=d=e]
	fmt.Println(strings.SplitN(str, "=", 2))  // [a b=c=d=e]
	fmt.Println(strings.SplitN(str, "=", 3))  // [a b c=d=e]
	fmt.Println(strings.SplitN(str, "=", 4))  // [a b c d=e]
	fmt.Println(strings.SplitN(str, "=", 5))  // [a b c d e]
	fmt.Println(strings.SplitN(str, "=", 6))  // [a b c d e]
	fmt.Println(strings.SplitN(str, "=", 7))  // [a b c d e]

}
