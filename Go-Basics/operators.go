package main

import "fmt"

func operators() {
	a := 10
	b := 3
	fmt.Println("a + b =", a+b) // Addition
	fmt.Println("a - b =", a-b) // Subtraction
	fmt.Println("a * b =", a*b) // Multiplication
	fmt.Println("a / b =", a/b) // Division
	fmt.Println("a % b =", a%b) // Modulus	
	// Increment and Decrement
	a++
	fmt.Println("a after increment =", a)
	b--
	fmt.Println("b after decrement =", b)
	// Comparison Operators
	fmt.Println("a == b:", a == b) // Equal to
	fmt.Println("a != b:", a != b) // Not equal to
	fmt.Println("a > b:", a > b)   // Greater than
	fmt.Println("a < b:", a < b)   // Less than
	fmt.Println("a >= b:", a >= b) // Greater than or equal to
	fmt.Println("a <= b:", a <= b)	 // Less than or equal to
	// Logical Operators
	x := true
	y := false
	fmt.Println("x && y:", x && y) // Logical AND
	fmt.Println("x || y:", x || y) // Logical OR
	fmt.Println("!x:", !x)         // Logical NOT

	// Bitwise Operators
	c := 5  // 0101 in binary
	d := 3 // 0011 in binary
	fmt.Println("c | d =", c|d)   // Bitwise OR
	fmt.Println("c ^ d =", c^d)   // Bitwise XOR	
	fmt.Println("c & d =", c&d)   // Bitwise AND
	fmt.Println("c &^ d =", c&^d) // Bitwise AND NOT

	// Shift Operators
	e := 8 // 1000 in binary
	fmt.Println("e << 1 =", e<<1) // Left shift
	fmt.Println("e >> 1 =", e>>1) // Right shift
}