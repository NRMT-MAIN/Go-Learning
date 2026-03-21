package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func sortingExample() {
	numbers := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("Before sorting:", numbers)

	sort.Ints(numbers)
	fmt.Println("After sorting:", numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println("After reverse sorting:", numbers)

	stringSlice := []string{"banana", "apple", "cherry"}
	fmt.Println("Before sorting:", stringSlice)
	sort.Strings(stringSlice)
	fmt.Println("After sorting:", stringSlice)

	num := []int{5, 2, 9, 1, 5, 6}
	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	fmt.Println("After custom sorting:", num)

	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i] > stringSlice[j]
	})
	fmt.Println("After custom sorting:", stringSlice)

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	
	sort.SliceStable(people , func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("After stable sorting:", people)
}