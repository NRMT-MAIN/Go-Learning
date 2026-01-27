package main

import "fmt"

func init() {
	fmt.Println("Panic Example Initialized")
}

func init() {
	fmt.Println("Another init function for Panic Example")
}

func init() {
	fmt.Println("Yet another init function for Panic Example")
}
func panicExample() {
	fmt.Println("Start of panicExample function")
	defer fmt.Println("Deferred: This will run at the end of panicExample function")

	fmt.Println("About to panic!")
	panic("This is a panic situation!")
	fmt.Println("This line will not be executed")
}

func process2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panicExample()
	fmt.Println("End of process2 function")
}