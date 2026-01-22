package main

import "fmt"

func loop() {

	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	//Iterate over collections

	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//If index is not needed, use _ 

	//Continue and Break
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		if i > 7 {
			break // Exit loop if i is greater than 7
		}
		fmt.Println("Odd number less than or equal to 7:", i)
	}

	//While loop using for
	count := 1
	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}

	for i := range 10 {
		fmt.Println(i)
		i++
	}

}