package main

import "fmt"

func funcExample() {
	// Function to add two integers
	add := func(a int, b int) int {
		return a + b
	}
	result := add(3, 5)
	fmt.Println("Sum:", result)

	//  Multiple return values
	swap := func(x, y string) (string, string) {
		return y, x
	}
	a, b := swap("hello", "world")
	fmt.Println("Swapped:", a, b)

	// Named return values
	divide := func(dividend, divisor float64) (quotient float64, remainder float64) {
		quotient = dividend / divisor
		remainder = float64(int(dividend) % int(divisor))
		return
	}
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	// Variadic function
	sumAll := func(numbers ...int) int {
		total := 0
		for _, number := range numbers {
			total += number
		}
		return total
	}

	totalSum := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Total Sum:", totalSum)

	// Function as a first-class citizen
	applyOperation := func(x, y int, operation func(int, int) int) int {
		return operation(x, y)
	}
	multiplication := func(a, b int) int {
		return a * b
	}
	resultMul := applyOperation(4, 5, multiplication)
	fmt.Println("Multiplication Result:", resultMul)

	// Anonymous function immediately invoked
	func(msg string) {
		fmt.Println("Anonymous Function says:", msg)
	}("Hello from anonymous function!")	

	var msgPrinter = func(message string) {
		fmt.Println("Message:", message)
	}
	msgPrinter("Hello, Go!")

	// Closure example
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
}