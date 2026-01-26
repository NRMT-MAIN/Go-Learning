package main

import "fmt"

func deferExample() {
	fmt.Println("Start of deferExample function")
	defer fmt.Println("Deferred: This will run at the end of deferExample function")

	fmt.Println("Doing some work in deferExample function")
	for i := 1; i <= 3; i++ {
		defer fmt.Println("Deferred in loop:", i)
	}
	fmt.Println("End of deferExample function")

	process(5)
}

func process(i int) {
	defer fmt.Println("Deferred in process function:", i)
	i++
	fmt.Println("Processing:", i)
}
