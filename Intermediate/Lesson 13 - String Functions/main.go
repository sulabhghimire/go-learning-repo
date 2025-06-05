package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	str := "Hello Go"
	fmt.Println("Length of str is", len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])   // ASCII value of the alphabet
	fmt.Println(str[1:5]) // Extract ello from Hello World

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println("Length of", str3, "is", len(str3))

	// String splitting
	fruits := "apple, orange, banana"
	parts := strings.Split(fruits, ", ")
	fmt.Println("Split of", fruits, "is", parts)

	// Joining slice of string to single string
	countries := []string{"Nepal", "USA", "Canada"}
	joined := strings.Join(countries, ", ")
	fmt.Println("Joined from list", countries, "is", joined)

	// Checking for subset
	fmt.Println(strings.Contains(str, "Go"))
	fmt.Println(strings.Contains(str, "Go?"))

	// Replacing an occurance of sub string within a string with another string value
	replaced := strings.Replace(str, "Go", "Universe", 1)
	fmt.Println("Go replaced with Universe from", str, "resulting", replaced)
	// replaceAll to replace all the occurances

	// Removing leading and trailing white spaces
	strwspace := " Hello Everyone! "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	// Case conversion
	fmt.Println(strings.ToLower(strwspace), strings.ToUpper(strwspace))

	// Repeat function that repeats something a fixed number of times
	fmt.Println(strings.Repeat("foo", 3))

	// Counting the occurances of a sub string within a string
	fmt.Println(strings.Count("Helllo", "l"))

	// Can check for prefix and suffix in a string
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	// Regular Expression
	str5 := "Hel1lo, 123 Go! 11"
	// Patterns must be inside backticks
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1) // -1 indicates  we are looking for all the matches
	fmt.Println(matches)

	// Strings in go are immutable and for that we have strings.Builder for efficient concatanation in peformance critical
	// scenario
	// strings.Builder is a type in go standard library in strings package
	// Helps to build incrementaly

	// String BUILDER
	var builder strings.Builder

	// Wrtie some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// Convert the builder to a string
	resultStr := builder.String()
	fmt.Println(resultStr)

	// Using WriteRune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you?")

	resultStr = builder.String()
	fmt.Println(resultStr)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh")
	resultStr = builder.String()
	fmt.Println(resultStr)

}
