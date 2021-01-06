package main

import (
	"fmt"
	"time"
)

/*
	Threads:
	- most programing languages uses OS threads, that means they got individual
	  function call stack dedicated to the execution of whatever code is handed
	  to that thread
	- traditionally these tend to be very very large, about 1MB of RAM, they
	  take quite a bit of time for the application to set up, so you have to be
	  very conservative about how you use the thread
	- that is why you need thread pooling and stuff like that, because the
	  creation and destruction of threads is very very expensive
	- now go uses green threads, instead of creating these very massive heavy
	  overhead of threads, we are going to create an abstraction of a thread we
	  call goroutine
	- inside of the go runtime we have a scheduler that's going to map these
	  goroutines onto these OS threads
	- scheduler takes turns with different OS threads available and assign a
	  certain amount of processing time on those threads and we don't have to
	  interact with those low-level threads directly
	- advantage: goroutine can start with very small stack spaces, because they
	  can be reallocated very quickly, so they're very cheap to create and
	  destroy
	- common to see 1,000s and 10,000s of goroutines running at the same time
	- not possible with languages that rely on OS threads
*/

func main() {
	go sayHello()                      // the go keyword spawns a new goroutine
	time.Sleep(100 * time.Millisecond) // remove this line and see what happens
	/*
		go scheduler is not going to interrupt the main thread until it hits
		Sleep call. Even though we create a goroutine, it still executes the
		main function, this is actually creating a race condition and it is a
		bad thing and should be avoided
	*/
}

func sayHello() { fmt.Println("Hello") }
