package main

import "fmt"

func formattingVerb() {
	// Formatting verbs in Go are used with the fmt package to format strings, numbers, and other data types.
	// Here are some common formatting verbs:

	// %v - the value in a default format
	// %T - a Go-syntax representation of the type of the value
	// %d - decimal integer
	// %f - floating-point number
	// %s - string
	// %q - double-quoted string
	// %p - pointer address
	// %t - boolean
	// %#v - a Go-syntax representation of the value

	// Example usage:
	name := "Alice"
	age := 30
	height := 5.7
	isStudent := false
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Student: %t\n", name, age, height, isStudent) //  Name: Alice, Age: 30, Height: 5.7, Student: false
	fmt.Printf("Value: %v, Type: %T\n", age, age) // Value: 30, Type: int
	fmt.Printf("Pointer address of age: %p\n", &age) // Pointer address of age: 0xc0000140b0
 
	fmt.Printf("Go-syntax representation of name: %#v\n", name) // Go-syntax representation of name: "Alice"
}