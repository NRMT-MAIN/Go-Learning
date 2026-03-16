package main

import (
	"fmt"
	"time"
)

func timerExample() {
	fmt.Println("Starting timer example...")
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for timer to expire...")
	stopped := timer.Stop()
	if !stopped {
		fmt.Println("Timer Stopped")
	}

	timer.Reset(time.Second)
	fireTimer := <-timer.C // Wait for the timer to expire
	fmt.Println("Timer expired!", fireTimer)


	time.AfterFunc(3 * time.Second , func() {
		fmt.Println("Timer expired after 3 seconds!")
	})

	fmt.Println(<-time.After(4 * time.Second))

	time.Sleep(5 * time.Second) // Wait to allow the AfterFunc to execute

	ticker := time.NewTicker(1 * time.Second)

	for tick := range ticker.C {
		fmt.Println("Tick at", tick)
		if tick.Second() % 5 == 0 {
			fmt.Println("Stopping ticker...")
			ticker.Stop()
			break
		}
	}
}