package main

import (
	"fmt"
	"strconv"
)

func numberParsingExample() {
	numStr := "12345"
	num , err := strconv.Atoi(numStr)

	if err != nil {
		fmt.Println("Error parsing the value : " , err)
	}

	fmt.Println("Parsing Integer : " , num + 1)

	numisstr , err := strconv.ParseInt(numStr , 10 , 32)
	if err != nil {
		fmt.Println("Error parsing the value : " , err)
	}

	fmt.Println("Parsed Integer : " , numisstr)

	floatstr := "3.14"
	floatval, err := strconv. ParseFloat (floatstr, 64)
	if err != nil {
		fmt.Println("Error parsingIvalue:", err)
	}

	fmt.Printf("Parsed float: %.2f\n", floatval)

	binaryStr := "1010"

	decimal , err := strconv.ParseInt(binaryStr , 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value : " , err)
	}

	fmt.Println("Parsed Binary to decimal" , decimal)
}