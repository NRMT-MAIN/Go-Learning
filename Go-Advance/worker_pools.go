package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personId   int
	numTickets int
	cost       int
}

// simulate a worker pool for processing ticket requests
func ticketProcessor(id int , request <-chan ticketRequest, results chan<- int) {
	for req := range request {
		fmt.Printf("Processing ticket request from person %d for %d tickets from  worker %d\n", req.personId, req.numTickets, id)
		time.Sleep(500 * time.Millisecond) // Simulate processing time
		results <- req.numTickets * 100 // Simulate cost calculation
	}
}

// func worker(id int, tasks <-chan int, result chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		time.Sleep(time.Millisecond * 500)
// 		result <- task * 2 // Simulate work by doubling the task value
// 	}
// }

func workerPoolExample() {
	// numWorkers := 3
	// numJobs := 10

	// tasks := make(chan int, numJobs)
	// results := make(chan int, numJobs)

	// // Start worker goroutines
	// for i := 1; i <= numWorkers; i++ {
	// 	go worker(i, tasks, results)
	// }

	// // Send jobs to the tasks channel
	// for j := 1; j <= numJobs; j++ {
	// 	tasks <- j
	// }
	// close(tasks)

	// // Collect results
	// for i := 0; i < numJobs; i++ {
	// 	result := <-results
	// 	fmt.Printf("Result: %d\n", result)
	// }
	// fmt.Println("All tasks processed.")

	numRequests := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	results := make(chan int, numRequests)

	// Start worker goroutines
	for i := 1; i <= 3; i++ {
		go ticketProcessor(i, ticketRequests, results)
	}

	// Send ticket requests
	for i := 1; i <= numRequests; i++ {
		ticketRequests <- ticketRequest{
			personId:   i,
			numTickets: i * 2,
		}
	}
	close(ticketRequests)

	// Collect results
	for i := 0; i < numRequests; i++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println("All ticket requests processed.")
}

