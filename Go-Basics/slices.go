package main

import "fmt"

func slices() {
	// Declare and initialize a slice
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", slice)

	// Accessing elements
	fmt.Println("First element:", slice[0])
	fmt.Println("Third element:", slice[2])
	// Modifying elements
	slice[1] = 25
	fmt.Println("Modified Slice:", slice)
	// Length of the slice
	fmt.Println("Length of slice:", len(slice))

	// Iterating over a slice
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Element at index %d: %d\n", i, slice[i])
	}

	// Using range to iterate
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	originalSlice := []int{1, 2, 3}
	copiedSlice := originalSlice
	copiedSlice[0] = 100
	fmt.Println("Original slice:", originalSlice)
	fmt.Println("Copied slice:", copiedSlice)

	slice2 := make([]int , 5)
	fmt.Println("Slice2:", slice2) // [0 0 0 0 0]

	slice2 = append(slice2, 10, 20)
	fmt.Println("Slice2 after append:", slice2) // [0 0 0 0 0 10 20]

	// Slicing a slice
	subSlice := slice[1:4]
	fmt.Println("Sub-slice (index 1 to 3):", subSlice) // [25 30 40]

	//Copying slices
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, 3)
	n := copy(dest, src)
	fmt.Println("Source Slice:", src)
	fmt.Println("Destination Slice after copy:", dest)
	fmt.Println("Number of elements copied:", n)

	// Multi-dimensional slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Slice (Matrix):", matrix)

	// arr := [3]int{1, 2, 3} is slice or array? 
	// It is an array. Arrays have fixed size, while slices are dynamic.
	// To convert an array to a slice, you can do the following:
	arr := [3]int{1, 2, 3}
	arrSlice := arr[:] // Convert array to slice
	fmt.Println("Array converted to Slice:", arrSlice)

}