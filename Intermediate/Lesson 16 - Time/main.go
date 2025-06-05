package main

import (
	"fmt"
	"time"
)

func main() {

	// Current local time
	fmt.Println("Local time:", time.Now())

	// Specific time
	specificTime := time.Date(2025, time.June, 6, 12, 0, 0, 0, time.UTC)
	fmt.Println("Specific time:", specificTime)

	// Parse Time
	parsedTime, _ := time.Parse("2006-01-02", "2020-06-27")          // Mon Jan 2  15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01")             // Mon Jan 2  15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-01-02 15-04", "20-05-01 18-03") // Mon Jan 2  15:04:05 MST 2006
	fmt.Println("Parsed Time:", parsedTime)
	fmt.Println("Parsed Time1:", parsedTime1)
	fmt.Println("Parsed Time1:", parsedTime2)

	// Fomatting Time
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("Mon 06-01-02 15-04-05"))

	// Manipulate time
	// Add or subtract durations
	// Round or truncate time

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println("One day later:", oneDayLater)
	fmt.Println("One day later: Weekday:", oneDayLater.Weekday())

	fmt.Println("Rounded Time:", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kathmandu")
	t = time.Date(2025, time.June, 4, 21, 40, 00, 00, time.UTC)

	// Conver this to specific time zone
	tLocal := t.In(loc)

	// Peform Rounding
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original Time UTC:", t)
	fmt.Println("Local Time:", tLocal)
	fmt.Println("Rounded Time UTC:", roundedTime)
	fmt.Println("Rounded Time Local:", roundedTimeLocal)

	fmt.Println("Truncated Time", time.Now().Truncate(time.Hour))

	fmt.Println("---------------------------------")
	loc, _ = time.LoadLocation("America/New_York")

	// Convert time
	tInNy := time.Now().In(loc)
	fmt.Println("New York Time:", tInNy)

	t1 := time.Date(2025, time.March, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, time.March, 4, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)

	// Compare time
	fmt.Println("t2 is after t1", t2.After(t1))

}
