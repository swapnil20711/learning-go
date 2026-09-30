package main

import (
	"fmt"
	"time"
)

func main() {
	// "dd-mm-yyyy"
	// "yyyy/mm/dd"

	currentTime := time.Now()

	fmt.Println("Current time:", currentTime)
	fmt.Printf("Type of current time: %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006, 3:04 PM")
	fmt.Println("Formatted time:", formatted)

	layoutStr := "02/01/2006"
	dateStr := "25/11/2030"

	formattedTime, _ := time.Parse(layoutStr, dateStr)

	fmt.Println("Formatted time: ", formattedTime)

	// add one more day in current time
	newDate := currentTime.Add(48 * time.Hour)

	formattedNewDate := newDate.Format("2006/01/02 Monday")

	fmt.Println("New date time:", newDate)
	fmt.Println("Formatted New date time:", formattedNewDate)
}
