package main

import (
	"fmt"
	"math"
)

func types() {
	var age int // Declares 'age' as an int with value 0.
	fmt.Println(age)

	var name string = "John" // Explicit type
	var name1 = "Jane"       // Type inferred as string
	//The type is optional if initializing as it is inferred from the value
	fmt.Println(name)
	fmt.Println(name1)

	//Short Variable Declaration := to declare and initialize in one step. Type is always inferred
	//Can only be used within function

	count := 10 // Declares and initializes an int variable 'count'
	fmt.Println(count)

	var isActive bool = true
	fmt.Println(isActive)

	//Multiple variables can be declared in a single statement.

	//a. On one line
	var price, quantity int = 5000, 100
	fmt.Println(price, quantity) // Output: 5000 100

	//b. In a block (useful for grouping)
	var (
		name2   = "Naveen" // type inferred
		age2    = 38       // type inferred
		height int        // zero value (0)
	)

	//c. Using shorthand (:=)
	a, b := 20, 30
	fmt.Println(a, b) // Output: 20 30
	//Rule for :=: At least one variable on the left-hand side must be new.

	//a, b := 20, 30 -- ERROR: no new variables

	//Correct usage:
	b, c := 40, 50 // OK: 'b' is already declared, but 'c' is new.
	// a, b := 40, 50 // ERROR: no new variables.

	fmt.Println(name2, age2, height)
	fmt.Println(c)

	const cnst = 50 // untyped constant
	const retryLimit = 4
	const httpMethod = "GET"

	// Grouped declaration
	const (
		pi = 3.14
		e  = 2.718
		thursday int = 4
	)

	// const b = math.Sqrt(4) // ERROR: math.Sqrt(4) is not a constant
	//Must be known at compile time. Cannot be assigned a value from a function call
	var sqr = math.Sqrt(4) // This is allowed (variable, not constant)
	fmt.Println(cnst, retryLimit, httpMethod, pi, e, thursday, sqr)
	

}
