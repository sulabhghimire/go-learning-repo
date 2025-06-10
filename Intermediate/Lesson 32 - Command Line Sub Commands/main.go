package main

import (
	"flag"
	"fmt"
	"os"
)

// While using subCommands and multiple flags use -processing=true -bytes=113 else only first flag will be evaluated
func main() {

	// stringFlag := flag.String("user", "Guest", "Name of the user")
	// flag.Parse()
	// fmt.Println(stringFlag)

	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of response")

	flagsc2 := subCommand2.String("language", "GO", "Enter your desired langauge")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language:", *flagsc2)
	default:
		fmt.Println("couldn't understand subcommnad:", os.Args[1])
		os.Exit(1)
	}

}
