package main

import (
	"fmt"
	"os"
)

func fileExample() {
	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("Error creating file." , file)
		return
	}

	defer file.Close()

	// write data to file
	data := []byte("Hello World \n")
	_ , err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file : " , err)
		return
	}

	fmt.Println("Data has been successfully written")

	content, err := os.ReadFile("output.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Content:", string(content))
}