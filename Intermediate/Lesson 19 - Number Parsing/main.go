package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Number parsing from string conversion package -> strconv

	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted num:", num)
	}

	num2, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted num Parsed Int:", num2)
	}

	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted num Float:", floatVal)
	}

	binrayString := "10"
	decimal, err := strconv.ParseInt(binrayString, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted binray", binrayString, "to decimal is", decimal)
	}

	hexString := "FF"
	decimal, err = strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted hexString", hexString, "to decimal is", decimal)
	}
}
