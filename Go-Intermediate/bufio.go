package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func bufioExample() {

	// reader := bufio.NewReader(strings.NewReader("Hello , bufio package! \n How are you ? ")) 
	// // Accepts another reader then it buffer that existing reader.
	
	// data := make([]byte , 20)

	// n , err := reader.Read(data)

	// if err != nil {
	// 	fmt.Println("Error reading : " , err)
	// 	return
	// }

	// fmt.Println("Read bytes : " , n)
	// fmt.Println(data) 

	// line , err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading string: " , err)
	// 	return
	// }

	// fmt.Println("Read String : " , line)

	fmt.Println()
	writer := bufio.NewWriter(os.Stdout)
	
	data := []byte("Hello , bufio package \n")

	n , err := writer.Write(data)

	if err != nil {
		fmt.Println("Error Writing :" , err)
	}
	fmt.Printf("Wrote %d bytes \n" , n)

	writer.Flush()

	// Writing string
	str := "This is a string.\n"

	n , err = writer.WriteString(str)

	if err != nil {
		fmt.Println("Error writing string:" , err)
		return
	}

	fmt.Printf("Wrote %d bytes. \n" , n)

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer: " , err)
		return
	}

}