package main

import (
	"fmt"
	"time"
)

func channelsExample() {
	greetings := make(chan string)
	greetString := "Hello, Channel!"

	go func() {
		greetings <- greetString
		greetings <- "Welcome to Go programming!"
	}()

	go func() {
		reciever := <-greetings
		println(reciever)
		reciever = <-greetings
		println(reciever)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("End of the Program")
}