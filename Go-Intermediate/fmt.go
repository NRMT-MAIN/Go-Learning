package main

import "fmt"

func fmtPackageExample() {
	// The fmt package in Go provides I/O formatting functions.
	// Here are some commonly used functions from the fmt package:
	// Println - prints to standard output with a newline
	// Printf - prints to standard output with formatted string
	// Sprintf - returns a formatted string
	// Scanln - reads input from standard input

	// Example usage:
	name := "Bob"
	age := 25
	fmt.Println("Hello,", name) // Hello, Bob
	fmt.Printf("%s is %d years old.\n", name, age)
	greeting := fmt.Sprintf("Welcome, %s!", name)
	fmt.Println(greeting)

	var inputName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&inputName)
	fmt.Println("You entered:", inputName)

	//Formatting Functions
	s := fmt.Sprintf("Name: %s, Age: %d", name, age)
	s = fmt.Sprint("Hello World , " , "from Go : " , 123)
	// Can be used to format various data types into a string
	// Useful for logging, displaying messages, etc.
	// Use for concatenation of different data types into a single string
	fmt.Println(s)

	// Scan vs Scanf vs Scanln
	// Scan - reads space-separated values from standard input
	// Scanf - reads formatted input based on a format string
	// Scanln - reads input until a newline is encountered
	var a int
	var b string
	fmt.Print("Enter an integer and a string (space-separated): ")
	fmt.Scan(&a, &b)
	//If not used &, it will not throw an error
	//It will just not store the input values into the variables 
	//just create new local variables inside the function scope
	fmt.Printf("You entered: %d and %s\n", a, b)

	fmt.Println("Enter an integer and a string (space-separated, ends with newline): ")
	var c int
	var d string
	fmt.Scanln(&c, &d) // Reads until newline
	//It discards any extra input after the newline
	//When I enter 1 Sahu from above scan and then press enter 
	//It will read 1 and Sahu and discard anything after that because of newline
	//Scanln doesn't take input after scan is used once
	fmt.Printf("You entered (Scanln): %d and %s\n", c, d)
	
	// Error Formatting Functions
	errMsg := fmt.Errorf("An error occurred: %s", "sample error")
	fmt.Println(errMsg) // An error occurred: sample error

	err := checkAge(16)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid.")
	}
	// Output: Error: Age 16 is below 18

}

func checkAge(age int) error{
	if age < 18 {
		return fmt.Errorf("Age %d is below 18", age)
	} 
	return nil
}