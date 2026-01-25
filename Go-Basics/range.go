package main

import "fmt"

func rangeExample() {
	message := "Hello, Go!" 

	// Iterate over each character in the string
	for index, char := range message {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// Iterate over a slice
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterate over a map
	capitals := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
	}
	for country, capital := range capitals {
		fmt.Printf("Country: %s, Capital: %s\n", country, capital)
	}
}