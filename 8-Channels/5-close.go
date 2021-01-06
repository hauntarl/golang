package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan int, 50) // making a buffered channel

	// dealing with buffered channels in goroutines
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch { // syntax a bit different from other datastructures
			fmt.Println(i)
		}
		// keeps on listening to data coming from channel infinitely
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 10
		ch <- 20
		close(ch)
		// Deadlock will occurr if we do not close the channel, why?
		// our receiver function is listening to data infinitely, the code will
		// never exit out of the for loop. When we close a channel, it sends a
		// singal notifying that no more data will come

		// ch <- 20 // not allowed to send data after closing, causes panic
		wg.Done()
	}(ch)
	wg.Wait()
}
