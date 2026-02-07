package main

import (
	"fmt"
	"time"
)

func timeExample() {

	fmt.Println(time.Now()) 

	// Specific Time
	specificTime := time.Date(2024 , time.March , 2 , 12, 0 , 0 , 0 , time.UTC)
	fmt.Println("Specific Time : " , specificTime)

	now := time.Now()

	fmt.Println(now.Minute())
	fmt.Println(now.Hour())
	fmt.Println(now.Day())

	//Formatting
	fmt.Println(now.Format("2006-01-02"))                 
	fmt.Println(now.Format("2006-01-02 15:04:05"))        
	fmt.Println(now.Format("Jan 02, 2006"))               
	fmt.Println(now.Format("3:04 PM"))                    
	fmt.Println(now.Format(time.RFC3339))                 
	fmt.Println(now.Format(time.Kitchen))   

	// Parsing time
	layout := "2004-06-12"
	t, err := time.Parse(layout, "2024-03-15")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)
	
	// Time Arithematic
	past := now.Add(-24 * time.Hour)
	future := now.Add(24 * time.Hour)

	fmt.Println(past)
	fmt.Println(future)
}
