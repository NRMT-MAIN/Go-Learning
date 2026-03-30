package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumberss() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetterss() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("%c\n", ch)
		time.Sleep(700 * time.Millisecond)
	}
}

func heavyComputation() {
	sum := 0
	for i := 0; i < 1e7; i++ {
		sum += i
	}
	fmt.Println("Heavy computation result:", sum)
}

func worker_pvc(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second) // Simulate work
        results <- job * 2
    }
}


func cvpExample() {
	// go printNumberss()
	// go printLetterss()
	// go heavyComputation()

	// time.Sleep(5 * time.Second) // Wait for goroutines to finish

	const numWorkers = 3
    const numJobs = 9
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    var wg sync.WaitGroup
    
    // Start workers (concurrent)
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker_pvc(w, jobs, results, &wg)
    }
    
    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Close results when all workers done
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Println("Result:", result)
    }
}