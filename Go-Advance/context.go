package main

import (
	"context"
	"fmt"
	"time"
)

func workers(ctx context.Context , id int) {
	for {
		select {
		case <-ctx.Done() :
			fmt.Printf("Worker %d: canceled (%v)\n", id, ctx.Err())
            return
		default : 
			fmt.Printf("Worker %d: working...\n", id)
            time.Sleep(500 * time.Millisecond)
		}
	}
}

func operation(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed")
		return nil
	case <-ctx.Done():
		fmt.Println("Operation canceled:", ctx.Err())
		return ctx.Err()
	}
}

func contextExample() {

	// todoContext := context.TODO()
	// println("Todo Context:", todoContext)

	// ctxBkg := context.Background()

	// ctx := context.WithValue(todoContext , "name" , "john")
	// fmt.Println(ctx)
	// fmt.Println(ctx.Value("name"))

	// ctx1 := context.WithValue(ctxBkg , "city" , "New Delhi")
	// fmt.Println(ctx1)
	// fmt.Println(ctx1.Value("city"))


	// ctx , cancel := context.WithCancel(context.Background())

	// for i := 1 ; i <= 3 ; i++ {
	// 	go worker(ctx , i)
	// }
	// time.Sleep(2 * time.Second)
	

	// fmt.Println("Cancelling workers...")
	// cancel()

	// time.Sleep(5 * time.Second)

	ctx , cancel := context.WithTimeout(context.Background() , 2 * time.Second)
	defer cancel()

	err := operation(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
}