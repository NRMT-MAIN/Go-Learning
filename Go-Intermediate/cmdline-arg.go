package main

import (
	"flag"
	"fmt"
	"os"
)

func cmdArgExample() {
	fmt.Printf("Program name: %s\n", os.Args[0])
    
    fmt.Printf("Number of arguments: %d\n", len(os.Args)-1)
    
    for i, arg := range os.Args[1:] {
        fmt.Printf("Argument %d: %s\n", i+1, arg)
    }

	// Define flags 
	var name string 
	var age int 
	var male bool

	flag.StringVar(&name , "name" , "John" , "Name of the user")
	flag.IntVar(&age , "age" , 12 , "Age of user")
	flag.BoolVar(&male , "male" , true , "Gender of user")

	flag.Parse()

	fmt.Println("Name : " , name)
	fmt.Println("Age : " , age)
	fmt.Println("Male : " , male)
}