package main

import "fmt"
func maps() {
	// Declare and initialize a map
	countryCapitalMap := map[string]string{
		"France":    "Paris",
		"Italy":     "Rome",
		"Japan":     "Tokyo",
		"India":     "New Delhi",
	}
	fmt.Println("Country-Capital Map:", countryCapitalMap)
	// Accessing elements
	fmt.Println("Capital of India:", countryCapitalMap["India"])	
	// Modifying elements
	countryCapitalMap["Italy"] = "Rome Updated"
	fmt.Println("Modified Map:", countryCapitalMap)

	// Deleting an element
	delete(countryCapitalMap, "Japan")
	fmt.Println("Map after deletion:", countryCapitalMap)
	// Checking if a key exists
	capital, exists := countryCapitalMap["France"]
	if exists {
		fmt.Println("Capital of France:", capital)
	} else {
		fmt.Println("France not found in the map")
	}
	// Iterating over a map
	for country, capital := range countryCapitalMap {
		fmt.Printf("Country: %s, Capital: %s\n", country, capital)
	}
	// Length of the map
	fmt.Println("Number of entries in the map:", len(countryCapitalMap))

	// Initializing an empty map
	newMap := make(map[string]int)
	newMap["One"] = 1
	newMap["Two"] = 2
	fmt.Println("New Map:", newMap)

	// key not present
	value, ok := newMap["Three"]
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key 'Three' not found in the map")
	}
	
	// Nested maps
	nestedMap := map[string]map[string]string{
		"USA": {
			"Capital": "Washington, D.C.",
			"Currency": "Dollar",
		},
		"Germany": {
			"Capital": "Berlin",
			"Currency": "Euro",
		},
	}
	fmt.Println("Nested Map:", nestedMap)
	fmt.Println("Capital of USA:", nestedMap["USA"]["Capital"])

	// Clearing a map
	for k := range countryCapitalMap {
		delete(countryCapitalMap, k)
	}
	fmt.Println("Map after clearing:", countryCapitalMap)
}