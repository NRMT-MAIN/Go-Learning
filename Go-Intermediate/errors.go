package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error : square root of negative number")
	}

	return x, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error : Empty data")
	}
	return nil
}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error : %s" , e.message)
}

func process2() error {
	return &myError{
		"Custom error message" , 
	}
}

func readConfig() error {
	return errors.New("Config Error")
}

func readData() error {
	err := readConfig()

	if err != nil {
		return fmt.Errorf("readData : %w" , err)
	}

	return nil
}

func errorExample() {
	// result, err := sqrt(16)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(result)

	// result1, err1 := sqrt(-16)

	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	// fmt.Println(result1)

	// data := []byte{}
	// if err := process(data); err != nil {
	// 	fmt.Println("Error : ", err)
	// 	return
	// }
	// fmt.Println("Data process successfully!")

	// err1 := process2() 
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	if err := readData() ; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Read Succesfully")
}
