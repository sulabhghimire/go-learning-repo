package main

import (
	"embed" // Using embed package for side effects i.e using inderictly --> called a blank import
	"fmt"
	"io/fs"
)

/*
The below comment will not be discarded by go compiler. The go compiler will notice this and execute this.
Syntax looks like go:embed <file/directory>
Need to import "embed" package
In our case the compiler is told it that it needs to include example.txt in the final exacutable
We need to use directive before var decleration
*/
//go:embed example.txt
var content string

//go:embed basics
var basicsFolder embed.FS

func main() {

	fmt.Println("Embedded content:", content)

	content, err := basicsFolder.ReadFile("basics/text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content of text.txt:", string(content))

	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(path)

		return nil
	})
	if err != nil {
		panic(err)
	}

}
