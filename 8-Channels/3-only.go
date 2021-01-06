package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan int)

	wg.Add(2)
	go func(ch <-chan int) { // receive only channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // send only channel
		ch <- 10
		wg.Done()
	}(ch)
	wg.Wait()
	// runtime casts the bi-directional channel into uni-directional channel,
	// something specific to channels only
}
