package main

import "fmt"

func typeConv() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Println("Integer to Float:", f)
	var u uint = uint(f)
	fmt.Println("Float to Unsigned Integer:", u)
	var b byte = byte(u)
	fmt.Println("Unsigned Integer to Byte:", b)
	var str string = string(b)
	fmt.Println("Byte to String:", str)
	var r rune = rune(i) // rune is an alias for int32
	fmt.Println("Integer to Rune:", r)
	var f2 float32 = float32(f)
	fmt.Println("Float64 to Float32:", f2)
	var i2 int8 = int8(i)
	fmt.Println("Integer to Int8:", i2)

	a := 80         // int
	c := 91.8       // float64
	// sum := a + c // ERROR: mismatched types
	sum := a + int(c) // OK: c is truncated to 91. Value of sum is 171.

	d := 10
	var j float64 = float64(d) // Explicit conversion required for assignment

	fmt.Println("Sum:", sum)
	fmt.Println("Integer d:", d)
	fmt.Println("Float j:", j)

}