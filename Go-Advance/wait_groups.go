package main

import (
	"sync"
	"time"
)

type Worker struct {
	id int
	task string
}

// Perform a task and signal completion using WaitGroup
func (w *Worker) doWork(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Worker", w.id, "is performing task:", w.task)
	time.Sleep(2 * time.Second)
	println("Worker", w.id, "has completed task:", w.task)
}

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() 
// 	println("Worker", id, "is starting work...")
// 	time.Sleep(2 * time.Second) 
// 	println("Worker", id, "has finished work.")
// }

func worker(id int , task <- chan int, results chan<- int , wg *sync.WaitGroup) {
	defer wg.Done()
	println("Worker", id, "is starting work...")
	time.Sleep(time.Second) // Simulate work by sleeping for 2 seconds
	for task := range task {
		results <- task * 2
	}
	println("Worker", id, "has finished work.")
}

func waitGroupsExample() {
	// var wg sync.WaitGroup

	
	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1) 
	// 	go worker(i, &wg)
	// }

	// wg.Wait() 
	// println("All workers have finished.")

	// var wg sync.WaitGroup
	// numWorkers := 3
	// numJobs := 5
	// results := make(chan int, numJobs)
	// tasks := make(chan int, numJobs)

	// for i := 1; i <= numJobs; i++ {
	// 	tasks <- i 
	// }
	// close(tasks)

	// wg.Add(numWorkers)

	// for i := 1; i <= numWorkers; i++ {
	// 	go worker(i + 1, tasks, results, &wg)
	// }

	// go func() {
	// 	wg.Wait()
	// 	close(results)
	// }()

	// for result := range results {
	// 	println("Received result:", result)
	// }

	var wg sync.WaitGroup
	tasks := []string{"digging", "watering", "harvesting" , "planting" , "weeding"}

	for i , task := range tasks  {
		worker := Worker{id: i + 1, task: task}
		wg.Add(1)
		go worker.doWork(&wg)
	}
	wg.Wait()

	println("Worker finished with work!")
}