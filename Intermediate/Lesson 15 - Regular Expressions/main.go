package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said \"I'am great!\"")

	// Compile a regular expression to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Test Strings
	email1 := "user@email.com"
	email2 := "invalid_email"

	// Match
	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))

	// Captring Groups
	// Compile a regex to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test string
	date := "2035-03-30"

	// Find all the matches
	subMatches := re.FindStringSubmatch(date)
	fmt.Println(subMatches)

	// Replacing charcater in target string
	// Source String
	str := "Hello World"
	re = regexp.MustCompile(`[aeiou]`)
	result := re.ReplaceAllString(str, "*")
	fmt.Println("Replaced String", result)

	// Flags And Options
	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)

	testString := "Golang is going great"
	// Match
	fmt.Println("Match:", re.MatchString(testString))

}
