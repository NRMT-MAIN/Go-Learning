package main

import (
	"fmt"
	"regexp"
)

func regexExample() {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	email1 := "nirmit@example.com"
	email2 := "invalid_email" 

	if re.MatchString(email1) {
		fmt.Println("Valid Email")
	} else {
		fmt.Println("Invalid Email")
	}

	if re.MatchString(email2) {
		fmt.Println("Valid Email")
	} else {
		fmt.Println("Invalid Email")
	}

	// Capturing Group
	re2 := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
	match := re2.FindStringSubmatch("Date: 2024-03-15")

	// Get group names
	names := re2.SubexpNames()

	// Create a map of name -> value
	result := make(map[string]string)
	for i, name := range names {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	// result["year"] = "2024"
	// result["month"] = "03"
	// result["day"] = "15"

	for key , value := range result {
		fmt.Println(key + " --> " + value)
	}

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)
	// re = regexp.MustCompile(`go`)

	// Test string
	text := "Golang is great"

	// Match
	fmt.Println("Match:", re.MatchString(text)) // true
}