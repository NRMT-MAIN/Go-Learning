package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	err     error
}

// Error returns the error message. Implementing Error() method
// of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d : %s, %v", e.code, e.message , e.err)
}

func doSomething() error {
	return &customError{
		code:    500,
		message: "Something went wrong",
	}
}

func doSomething2() error {
	return errors.New("Internal Error")
}

func doSomething1() error {
	er := doSomething2()

	if er != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			err:     er,
		}
	}

	return nil
}

func customErrorExample() {
	// err := doSomething()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Operation completed successfully!")

	//Error 500 : Something went wrong

	err := doSomething1()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation completed successfully!")

	//Error 500 : Something went wrong, Internal Error
}
