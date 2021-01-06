package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan int, 50) // making a buffered channel

	// checking whether channel is closed
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i, ok)
			} else {
				fmt.Println(i, ok)
				break
			}
			// waits indefinitely until someone sends data over the channel
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 10
		ch <- 20
		close(ch)
		// When the channel is closed, it returns the zero value of data the
		// channel allows, along with a boolean value which represents whether
		// the channel is closed or not, false if closed.
		wg.Done()
	}(ch)
	wg.Wait()
}
