package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan int)

	wg.Add(2)
	go func() {
		fmt.Println(<-ch) // 2. stops execution, until it reads some data
		ch <- 20          // 3. writing data on the channel
		wg.Done()
	}()
	go func() {
		ch <- 10          // 1. writing data on the channel
		fmt.Println(<-ch) // 4. stops execution until it reads some data
		wg.Done()
	}()
	wg.Wait()
	// both goroutines are acting as readers and writers, but very often you
	// want to dedicate a goroutine to either reading-from or writing-to a
	// channel
}
