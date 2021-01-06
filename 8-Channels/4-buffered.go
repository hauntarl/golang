package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan int, 50) // making a buffered channel

	wg.Add(2)
	go func(ch <-chan int) {
		fmt.Println(<-ch) // Receives will block when the buffer is empty.
		fmt.Println(<-ch) // not the ideal to deal with buffered channels
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 10 // Sends to a buffered channel,
		ch <- 20 // blocks only when the buffer is full.
		wg.Done()
	}(ch)
	wg.Wait()
	// when sender and receiver operate at a different frequency, buffered
	// channels become very useful
}
