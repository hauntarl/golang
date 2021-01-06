package main

import (
	"fmt"
	"sync"
)

/*
	Mutexes:
	- mutual exclusion
	- RWMutex: many can read, but only one can write and while reading
      no one can write
*/

var (
	wg      = sync.WaitGroup{}
	m       = sync.RWMutex{}
	counter = 0
)

func main() {
	fmt.Println("\nSpawning 10 go routines using mutexes")
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHi()
		go increment()
	}
	wg.Wait()
	// this only fixes a part of our problem, as we still don't have control
	// over synchronizing our goroutines, though we are reading in sequence
}

func sayHi() {
	m.RLock()                  // applying read lock
	fmt.Println("Hi", counter) // reading
	m.RUnlock()                // releasing shared resource
	wg.Done()
}

func increment() {
	m.Lock()   // applying write lock
	counter++  // writing
	m.Unlock() // releasing shared resource
	wg.Done()
}
