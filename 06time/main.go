package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in golang")

	// current time
	currentTime := time.Now()
	fmt.Println(currentTime)

	// only date
	fmt.Println(currentTime.Format("01-02-2006"))

	// only time
	fmt.Println(currentTime.Format("15:04:05"))

	// date and time
	fmt.Println(currentTime.Format("01-02-2006 15:04:05"))

	// date and day
	fmt.Println(currentTime.Format("01-02-2006 Monday"))

	// date, day and time
	fmt.Println(currentTime.Format("01-02-2006 Monday 15:04:05"))

	// in different order
	fmt.Println(currentTime.Format("Monday 01-02-2006 15:04:05"))

	// create a date
	createdDate := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

// Read more about time here: https://golang.org/pkg/time/
// "01-02-2006 15:04:05 Monday" is a format string which is used to format the time
// we use particular these time, date and day formats to get the output we want
// 	createdDate := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC) here parameter is year, month, day, hour, minute, second, nanosecond, location
// type of year is int
// type of month is time.Month
// type of day is int
// type of hour is int
// type of minute is int
// type of second is int
// type of nanosecond is int
// type of location is *time.Location
