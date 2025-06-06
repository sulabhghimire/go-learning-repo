package main

import (
	"fmt"
	"time"
)

func main() {

	// Epoch -> Time since 00:00:00 UTC on Janaury 1, 1970 -> doesn't account leap seconds -> In seconds or milliseoncds
	// Epoch time values -> + after and -before
	// Epoch Applications --> a) Database Storage, b) System TimeStamps, c) Cross-Platform Compatibility

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current unix time:", unixTime)

	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println(t.Format("2006-01-02"))

	// Time Formatting and Parsing
	// Mon Jan 2 15:04:05 MST 2006 --> Reference for the layout
	// Mountain Standard time that has 7 hrs difference from UTC

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z" // Not mentioning timezone defaults to UTC Timezone

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	fmt.Println(t)

	str1 := "Jul 02 2024 02:18 PM"
	layout1 := "Jan 02 2006 03:04 PM"
	t, err = time.Parse(layout1, str1)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	fmt.Println(t)

}
