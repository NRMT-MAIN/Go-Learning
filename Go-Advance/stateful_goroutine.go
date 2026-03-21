package main

import "fmt"

type StatefulGoroutine struct {
	count int
	ch    chan int
}

func (w *StatefulGoroutine) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current count:" , w.count)
			}
		}
	}()
}

func (w *StatefulGoroutine) Send(value int) {
	w.ch <- value
}

func statefulGoroutine() {
	w := &StatefulGoroutine{
		count: 0,
		ch:    make(chan int),
	}
	w.Start()
	w.Send(5)
	w.Send(10)
}