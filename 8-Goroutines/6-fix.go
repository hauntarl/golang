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
	fmt.Println("\nSpawning 10 go routines using mutexes, synchronizing them")
	runtime.GOMAXPROCS(100) // refer next section for explanation of this
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock() // applying read lock
		go sayHi()
		m.Lock() // applying write lock
		go increment()
	}
	wg.Wait()
	// problem with this is, we have completely destroyed concurrency and
	// parallelism, this application will probably perform worse than without
	// goroutines. If you are doing operations like this, you're better off
	// without goroutines

	fmt.Println("\nThreads:", runtime.GOMAXPROCS(-1))
}

func sayHi() {
	fmt.Println("Hi", counter) // reading
	m.RUnlock()                // releasing shared resource
	wg.Done()
}

func increment() {
	counter++  // writing
	m.Unlock() // releasing shared resource
	wg.Done()
}
