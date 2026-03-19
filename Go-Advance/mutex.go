package main

import (
	"sync"
	"sync/atomic"
)

type counter struct {
	mu sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.count , 1) ; 
}

func mutexExample() {
	// var wg sync.WaitGroup

	// counter := &counter{}
	// numWorkers := 5

	// for i := 0; i < numWorkers; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for j := 0; j < 1000; j++ {
	// 			counter.increment()
	// 			//counter.count++
	// 		}
	// 	}()
	// }
	// wg.Wait()
	// println("Final count:", counter.getCount())

	// var counter int
	// var wg sync.WaitGroup
	// var mu sync.Mutex

	// numGoroutines := 5
	// wg.Add(numGoroutines)

	// increment := func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 1000; i++ {
	// 		mu.Lock()
	// 		counter++
	// 		mu.Unlock()
	// 	}
	// }

	// for range numGoroutines {
	// 	go increment()
	// }

	// wg.Wait()
	// println("Final counter value:", counter)

	var wg sync.WaitGroup
	counter := &AtomicCounter{}
	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			counter.Increment()
		}
	}

	for range numGoroutines {
		go increment()
	}
	wg.Wait()
	println("Final counter value:", counter.count)
}