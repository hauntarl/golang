package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg      = sync.WaitGroup{}
	m       = sync.RWMutex{}
	counter = 0
)

func main() {
	fmt.Println("\nThreads:", runtime.GOMAXPROCS(-1))
	// runtime.GOMAXPROCS(-1) gives total number of threads you can work with,
	// default is the total OS threads you have at the moment you invoke the
	// function, it essentially returns previously set threads.
	// If you do not wish to change the number of threads assigned, put
	// negative number

	// you can restrict the number of threads to 1, we can achieve
	// synchronization for previous problem.
	// what it does is: enable concurrency but restrict parallelism
	fmt.Println("\nRestricting number of threads to 1")
	runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHi()
		go increment()
	}
	wg.Wait()
	/*	- runtime.GOMAXPROCS(<int>) is a tuning tool, it is recommended to have
		  minimun of total OS threads.
		- If you increase the number of threads, the performance will increase
		  but you shouldn't go too crazy with it because you can run into other
		  problems such as: memory overhead for maintaining all those threads,
		  scheduler has to work harder.
		- Fine tune your application by testing it using various number of
		  threads

		Best Practices:
		- let consumer control concurrency, avoid creating go routines in
		  library
		- when you create goroutine, know how it will end, avoid subtle memory
		  leaks
		- check for race conditions at compile time, use: go run -race <file>.go
	*/
}

func sayHi() {
	fmt.Println("Hi", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
