package main

import (
	"fmt"
	"sync"
)

/*
	Channels:
	- used to pass data between the goroutines in a way that is safe, prevents
	  issues such as race conditions and memory sharing problems
	- we will be mostly using channels in the context of goroutines

	Creating:
	- only way is to use the make() function
	- make(chan <datatype>)
	- strongly typed
*/

var wg = sync.WaitGroup{}

func main() {
	var ch = make(chan int) // creating a channel object

	wg.Add(2)
	go receive(ch)
	go send(ch)
	wg.Wait()
	/*
		- By default, sends and receives block until the other side is ready.
		- This allows goroutines to synchronize without explicit locks or
		  condition variables.
		- if we try to run the above code snippet without using the go keywork
		  i.e. as a sequential program, we will run into a deadlock
		- when the main thread encounters receive function and starts execution,
		  it reaches the recieve from channel statement, as there is no data in
		  the channel it will block further execution until someone sends the
		  data.
		- eventually the main thread gives up its control, similar to what
		  happens when we perform Sleep operation or Wait operation
		- the go scheduler sees that their is no other goroutine present for
		  execution (no one to send the data onto the channel), the main thread
		  gets blocked indefinitely resulting in a deadlock
		- we don't need to use the WaitGroup here, but we are performing send
		  and receive both on different goroutines and not on the main thread,
		  we need something in the main thread which tells it to wait until
		  other goroutines finishes their execution
		- for more information about how go handles this situation, try running
		  the above code snippet without the go keyword
	*/

	// above example but with more than 2 goroutines
	fmt.Println("\nSpawning 10 go routines")
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go receive(ch)
		go send(ch)
	}
	wg.Wait()
}

func receive(ch chan int) {
	var i = <-ch   // receive data from the channel
	fmt.Println(i) // arrow determines the flow of data
	wg.Done()
}

func send(ch chan int) {
	var i = 10
	ch <- i // send message into the channel "<-" arrow operator
	i = 20  // channel uses pass by value
	wg.Done()
}
