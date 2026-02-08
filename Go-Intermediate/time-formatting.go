package main

import (
	"fmt"
	"time"
)

func timeFormattingExample() {
	now := time.Now()

	// Standard formats
	fmt.Println(now.Format("2006-01-02"))          // 2024-03-15
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 2024-03-15 14:30:45
	fmt.Println(now.Format("02-01-2006"))          // 15-03-2024
	fmt.Println(now.Format("01/02/2006"))          // 03/15/2024 (US format)
	fmt.Println(now.Format("02/01/2006"))          // 15/03/2024 (EU format)

	// With time
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))   // 2024-03-15 14:30:45.123
	fmt.Println(now.Format("2006-01-02T15:04:05Z07:00")) // 2024-03-15T14:30:45-07:00
	fmt.Println(now.Format("Jan 02, 2006 at 3:04 PM"))   // Mar 15, 2024 at 2:30 PM

	// Day and month names
	fmt.Println(now.Format("Monday, January 2, 2006"))  // Friday, March 15, 2024
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006")) // Fri Mar 15 14:30:45 2024
	fmt.Println(now.Format("Jan"))                      // Mar
	
	layout := "2006-01-02"
	dateStr := "2024-03-15"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(t) // 2024-03-15 00:00:00 +0000 UTC
}