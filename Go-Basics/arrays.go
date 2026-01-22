package main

import "fmt"

func arrays() {
	// Declare and initialize an array
	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)	
	// Accessing elements
	fmt.Println("First element:", arr[0])
	fmt.Println("Third element:", arr[2])

	// Modifying elements
	arr[1] = 25
	fmt.Println("Modified Array:", arr)
	// Length of the array
	fmt.Println("Length of array:", len(arr))
	// Iterating over an array
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element at index %d: %d\n", i, arr[i])
	}
	// Using range to iterate
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt. Println("Copied array:", *copiedArray)
}