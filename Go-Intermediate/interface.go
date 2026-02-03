package main

import (
	"fmt"
	"math"
)


type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width) 
}
//If we comment one of the method
//.\interface.go:48:16: r.perim undefined (type rect has no field or method perim)

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func interfaceExample() {
	r := rect{
		width: 3,
		height: 4,
	}

	c := circle {
		radius: 4,
	}

	fmt.Println(c.perim())  
	fmt.Println(r.perim())
	fmt.Println(c.area())
	fmt.Println(r.area())
}