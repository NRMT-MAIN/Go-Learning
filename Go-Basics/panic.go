package main

import "fmt"

func panicExample() {
	fmt.Println("Start of panicExample function")
	defer fmt.Println("Deferred: This will run at the end of panicExample function")

	fmt.Println("About to panic!")
	panic("This is a panic situation!")
	fmt.Println("This line will not be executed")
}