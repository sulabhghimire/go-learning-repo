package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// A string is a sequence of bytes -> bytes are uint8 values --> represent text
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!" // \r takes the cursor to first position in the line causing the output to be Go!sage Two: Hello,
	rewMessage := `Hello\nGo`  // means it is a raw string literal and escape sequences will not work

	fmt.Println("Message:", message)
	fmt.Println("Message One:", message1)
	fmt.Println("Message Two:", message2)
	fmt.Println("Raw Message:", rewMessage)

	// len function can be used to calculate length of string
	fmt.Println("Length of message variable is:", len(message)) // 11 the escape sequence \n is counted as one
	fmt.Println("Length of raw message variable is:", len(rewMessage))

	// we can extract any character from string by using index and doing so returns its byte value
	fmt.Println("The first character in message is", message[0]) // ASCII value of H

	// String concatanation
	greeting := "Hello"
	name := "Sulabh"
	fmt.Println(greeting + name)

	// String comparision
	str1 := "Apple"
	str2 := "Banana"
	fmt.Println("Str 1 < Str2 :", str1 < str2) // This output is given on basis of lexical graphical comparision
	// Lexicial graphical comparision is also known as lexical graphical ordering or dictionary ordering
	// Method of comparing strings based on alphabetical order of their components

	for i, char := range message {
		fmt.Printf("Character at index %d id %c\n", i, char)
		fmt.Printf("%x\n", char) // Hexa decimal value
		fmt.Printf("%v\n", char) // ASCII value of rune in uint32 or default value
	}

	// like length but counts number of runes character is string
	fmt.Println("Rune count in greetings:", utf8.RuneCountInString(greeting))

	// Strings are immutable so to do something we need to create new string
	greetingWithName := greeting + " " + name
	fmt.Println("Greeting with name", greetingWithName)

	// IN GO rune is alias to int32 and represents a unicode code point
	// Used to represent single character in string

	var ch rune = 'a'
	fmt.Printf("Rune %v: Character %c:\n", ch, ch)

	// Converts runes to string
	cstr := string(ch)
	fmt.Println(cstr)

	// We can use format verb %t to get type of anu variable
	// Type of rune character is int32
	fmt.Printf("Type of cstr is %T", cstr)

}
