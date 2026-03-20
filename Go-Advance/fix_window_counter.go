package main

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindowCounter struct {
	mu          sync.Mutex
	limit       int
	window      time.Duration
	counter     int
	windowStart time.Time
}

func NewFixedWindowCounter(limit int, window time.Duration) *FixedWindowCounter {
	return &FixedWindowCounter{
		limit:       limit,
		window:      window,
		windowStart: time.Now(),
	}
}

func (fw *FixedWindowCounter) Allow() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	now := time.Now()
	if now.Sub(fw.windowStart) >= fw.window {
		fw.counter = 0
		fw.windowStart = now
	}
	if fw.counter < fw.limit {
		fw.counter++
		return true
	}
	return false
}


func fixWindowCounter() {
	limiter := NewFixedWindowCounter(5 , time.Second)

	for i := 0; i < 20; i++  {
		if limiter.Allow() {
			fmt.Printf("Request %d : ALLOWED at %s\n" , i + 1 , time.Now().Format("15:04:05:0000"))
		} else {
			fmt.Printf("Request %d : REJECTED at %s\n" , i + 1 , time.Now().Format("15:04:05:0000"))
		}
		time.Sleep(150 * time.Millisecond)
	}

}
