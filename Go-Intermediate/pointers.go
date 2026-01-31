package main

import "fmt"

func pointerExample() {
	x := 10
	p := &x // p is a pointer to x

	fmt.Println("Pointer to x : " , p)

	*p = 20                   // changing the value at the address p points to
	println("Value of x:", x) // Output: Value of x: 20

	var ptr *int // declaring a pointer variable
	fmt.Println("Value of ptr:", ptr) // Output: Value of ptr: <nil> a nil pointer
	var a int = 10
	ptr = &a // assigning address of a to ptr
	fmt.Println("Address of a:", ptr)
	fmt.Println("Value at ptr:", *ptr) // dereferencing ptr to get value of a

	//Modifying Pointer
	modifyValue(ptr)
	fmt.Println(a) 
}

func modifyValue(ptr *int) {
	*ptr++
}
