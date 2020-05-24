package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	/*
		Threads:
			- most programing languages uses OS threads, that means they got individual function
			  call stack dedicated to the execution of whatever code is handed to that thread
			- traditionally these tend to be very very large, about 1MB of RAM, they take quite a
			  bit of time for the application to set up, so you have to be very conservative about
			  how you use the thread
			- that is why you need thread pooling and stuff like that, bacause the creation and
			  destruction of threads is very very expensive
			- now go uses green threads, instead of creating these very massive heavy overhead of
			  threads, we are going to create an abstraction of a thread we call goroutine
			- inside of the go runtime we have a scheduler that's going to map these goroutines
			  onto these OS threads
			- scheduler takes turns with different OS threads available and assign a certain amount
			  of processing time on those threads and we don't have to interact with those low-level
			  threads directly
			- advantage: goroutine can start with very small stack spaces, because they can be
			  reallocated very quickly, so they're very cheap to create and destroy
			- common to see 1,000s and 10,000s of goroutines running at the same time
			- not possible with languages that rely on OS threads
	*/
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	msg := "Hello again"
	go func() {
		fmt.Println(msg) // go has closures
		// function can access variables from its outer scope even if we are creating
		// a different execution stack, runtime knows where to look up the variable
		// so any changes made to the msg varible before goroutine is executed, will be
		// reflected at runtime
	}()
	msg = "Goodbye"
	// go scheduler is not going to interrupt the main thread until it hits Sleep call
	// even though we create a goroutine, it still executes the main function
	// this is actually creating a race condition and it is a bad thing and should be avoided
	time.Sleep(100 * time.Millisecond)

	msg = "Hello again"
	go func(msg string) {
		fmt.Println(msg)
		// providing a variable in local context for execution of this function stack
		// we eliminate the issue we faced earlier, using pass by value (only)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
	// we should avoid using Sleep calls in production as they can be very unreliable,
	// as you need your application to sleep relative to the average performance of your
	// application and not the system clock

	/*
		WaitGroup:
			- synchronizes goroutines, joins them
	*/
	msg = "Hello again"
	wg.Add(1) // adds number of goroutines we want our applicatio to wait for
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // decrements the number of goroutines by 1
	}(msg)
	msg = "Goodbye"
	wg.Wait() // waits until number is 0

	// spawning 20 go routines
	fmt.Println("\nSpawning 10 go routines")
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHi()
		go increment()
	}
	wg.Wait() // goroutines are racing against each other, no sync

	/*
		Mutexes:
			- mutual exclusion
			- RWMutex: many can read, but only one can write and while reading
			  no one can write
	*/
	fmt.Println("\nSpawning 10 go routines using mutexes")
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go mutexSayHi()
		go mutexIncrement()
	}
	wg.Wait()
	// this only fixes a part of our problem, as we still don't have control over
	// synchronizing our goroutines, though we are reading in sequence

	fmt.Println("\nSpawning 10 go routines using mutexes and synchronizing them")
	runtime.GOMAXPROCS(100) // refer end of this function, comment this to print total OS threads
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock() // applying read lock
		go mutexUnlockOnlySayHi()
		m.Lock() // applying write lock
		go mutexUnlockOnlyIncrement()
	}
	wg.Wait()
	// problem with this is, we have completely destroyed concurrency and parallelism
	// this application will probably perform worse than without goroutines
	// if you are doing operations like this, you're better off without goroutines

	fmt.Println("\nThreads:", runtime.GOMAXPROCS(-1))
	// runtime.GOMAXPROCS(-1) gives total number of threads you can work with
	// default is the total OS threads you have
	// whenever you invoke the function, it returns previously set threads, if you do
	// not wish to change the number of threads assigned, put negative number

	// you can restrict the number of threads to 1, we can achieve synchronization for
	// above problem doing this, what it does is: enable concurrency but restrict parallelism
	fmt.Println("\nRestricting number of threads to 1")
	runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHi()
		go increment()
	}
	wg.Wait()
	// runtime.GOMAXPROCS(<int>) is a tuning tool
	// it is recommended to have minimun of total OS threads
	// if you increase the number of threads, the performance will increase
	// but you shouldn't go too crazy with it because you can run into other problems
	// such as: memory overhead for maintaining all those threads, scheduler has to work harder
	// fine tune your application by testing it using various number of threads

	/*
		Best Practices:
			- let consumer control concurrency, avoid creating go routines in library
			- when you create goroutine, know how it will end, avoid subtle memory leaks
			- check for race conditions at compile time, use: go run -race <file>.go
	*/
}

var wg = sync.WaitGroup{}

var counter = 0

func sayHi() {
	fmt.Println("Hi", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

var m = sync.RWMutex{}

func mutexSayHi() {
	m.RLock()                        // applying read lock
	fmt.Println("Mutex Hi", counter) // reading
	m.RUnlock()                      // releasing shared resource
	wg.Done()
}

func mutexIncrement() {
	m.Lock()   // applying write lock
	counter++  // writing
	m.Unlock() // releasing shared resource
	wg.Done()
}

func mutexUnlockOnlySayHi() {
	fmt.Println("Mutex Hi", counter) // reading
	m.RUnlock()                      // releasing shared resource
	wg.Done()
}

func mutexUnlockOnlyIncrement() {
	counter++  // writing
	m.Unlock() // releasing shared resource
	wg.Done()
}
