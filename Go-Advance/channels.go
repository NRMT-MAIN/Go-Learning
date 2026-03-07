package main

import (
	"fmt"
)

func channelsExample() {
	// greetings := make(chan string)
	// greetString := "Hello, Channel!"

	// go func() {
	// 	greetings <- greetString
	// 	greetings <- "Welcome to Go programming!"
	// }()

	// time.Sleep(5 * time.Second) // Unbuffer channel will block until the value is received
	// go func() {
	// 	reciever := <-greetings
	// 	println(reciever)
	// 	reciever = <-greetings
	// 	println(reciever)
	// }()

	// time.Sleep(1 * time.Second)
	// fmt.Println("End of the Program")

	// ch := make(chan int)
	// go func() {
	// 	ch <- 42
	// 	time.Sleep(2 * time.Second) // Simulate some work before sending the value
	// 	fmt.Println("2 second Goroutine finished")
	// }()

	// go func ()  {
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("3 second Goroutine finished")
	// }()

	// reciever := <-ch 
	// fmt.Println("Received value:", reciever)
	// fmt.Println("End of the Program")

	// Unbuffered channel example
	// ch := make(chan int , 3) // Buffered channel with capacity of 3

	// go func() {
	// 	ch <- 1
	// 	fmt.Println("Sent 1")
	// 	ch <- 2
	// 	fmt.Println("Sent 2")
	// 	ch <- 3
	// 	fmt.Println("Sent 3")
	// }()

	// time.Sleep(1 * time.Second) // Simulate some work before receiving values
	// //ch <- 4 // This will block until there is space in the channel
	// //fmt.Println("Sent 4")
	// fmt.Println("Value: " , <-ch)
	// fmt.Println("Value: " , <-ch)
	// ch <- 4 // This will block until there is space in the channel
	// fmt.Println("Sent 4")
	// fmt.Println("Value: " , <-ch)


	// time.Sleep(1 * time.Second) // Wait for all goroutines to finish
	// fmt.Println("End of the Program")
	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Received:", <- ch)
	// }()
	// fmt.Println("Blocking starts!")
	// ch <- 3
	// fmt.Println("Blocking ends!")
	// fmt.Println("Received:", <- ch)
	// fmt.Println("Received:", <- ch)

	// fmt.Println("Buffered Channels")

	// Blocking on recieve only if the buffer is empty
	// ch := make(chan int, 2)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- 1
	// 	ch <- 2
	// }()

	// fmt.Println("Value : ", <-ch)
	// fmt.Println("Value : ", <-ch)
	// fmt.Println("End of the program")

	// ch1 := make(chan string)
    // ch2 := make(chan string)
    
    // go func() {
    //     time.Sleep(100 * time.Millisecond)
    //     ch1 <- "one"
    // }()
    
    // go func() {
    //     time.Sleep(50 * time.Millisecond)
    //     ch2 <- "two"
    // }()
    
    // // Select waits for first ready channel
    // select {
	//     case msg1 := <-ch1:
	//         fmt.Println("Received:", msg1)
	//     case msg2 := <-ch2:
	//         fmt.Println("Received:", msg2)
	// 	default:
	// 		fmt.Println("No channel is ready")
    // }

	//ch := make(chan int)
    
    // select {
	// 	case value := <-ch:
	// 		fmt.Println("Received:", value)
	// 	default:
	// 		fmt.Println("No value ready, not blocking")
    // }
    
    // Non-blocking send
    // select {
	// 	case ch <- 42:
	// 		fmt.Println("Sent 42")
	// 	default:
	// 		fmt.Println("Channel not ready, not blocking")
    // }

	// Timeout
	// go func() {
    //     time.Sleep(2 * time.Second)
    //     ch <- 42
	// 	close(ch) // Close the channel after sending the value
    // }()
    
    // select {
	// 	case value := <-ch:
	// 		fmt.Println("Received:", value)
	// 	case <-time.After(3 * time.Second):
	// 		fmt.Println("Timeout!")
    // }

	ch1 := make(chan int)
    ch2 := make(chan int)
    done := make(chan bool)
    
    go func() {
        ch1 <- 1
    }()
    
    go func() {
        ch2 <- 2
    }()
    
    go func() {
        done <- true
    }()
    
    for i := 0; i < 3; i++ {
        select {
			case v1 := <-ch1:
				fmt.Println("From ch1:", v1)
			case v2 := <-ch2:
				fmt.Println("From ch2:", v2)
			case <-done:
				fmt.Println("Done signal")
        }
    }
}