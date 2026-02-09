package main

import (
	"fmt"
	"math/rand"
)

func randExample() {
	fmt.Println(rand.Intn(101)) 

	fmt.Println(rand.Intn(6) + 5)

	val := rand.New(rand.NewSource(42))

	fmt.Println(val.Intn(101))
}