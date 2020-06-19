package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	/*
		Channels:
			- used to pass data between the goroutines in a way that is safe,
			  prevents issues such as race conditions and memory sharing problems
			- we will be mostly using channels in the context of goroutines

		Creating:
			- only way is to use the make() function
			- make(chan <datatype>)
			- strongly typed
	*/
	var ch = make(chan int)
	wg.Add(2)
	go func() {
		var i = <-ch // receive data from the channel
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		var i = 10
		ch <- i // send message into the channel "<-" arrow operator
		i = 20  // channel uses pass bt value
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("\nSpawning 5 go routines")
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			var i = <-ch // receive data from the channel
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			var i = 20
			ch <- i
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("\nread write go routines")
	wg.Add(2)
	go func() {
		fmt.Println(<-ch) // 2. stops execution of function until it reads some data
		ch <- 20          // 3. writing data on the channel
		wg.Done()
	}()
	go func() {
		ch <- 10          // 1. writing data on the channel
		fmt.Println(<-ch) // 4. stops execution of function until it reads some data
		wg.Done()
	}()
	wg.Wait()
	// both goroutines are acting as readers and writers, but very often you want to dedicate
	// a goroutine to either reading-from or writing-to a channel

	fmt.Println("\nrestricting to read/write only go routines")
	wg.Add(2)
	go func(ch <-chan int) { // receive only channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch) 
	// runtime casts the bi-directional channel into uni-directional channel
	// something specific to channels only
	go func(ch chan<- int) { // send only channel
		ch <- 10
		wg.Done()
	}(ch)
	wg.Wait()

	fmt.Println("\nbuffered channels go routines")
	var bufch = make(chan int, 50) // making a buffered channel
	wg.Add(2)
	go func(bufch <-chan int) {
		fmt.Println(<-bufch)
		fmt.Println(<-bufch) // this isn't the ideal way of dealing with buffered channels
		wg.Done()
	}(bufch)
	go func(bufch chan<- int) {
		bufch <- 10
		bufch <- 20
		wg.Done()
	}(bufch)
	wg.Wait()
	// when sender and receiver operate at a different frequency

	fmt.Println("\ndealing with buffered channels go routines")
	wg.Add(2)
	go func(bufch <-chan int) {
		for i := range bufch { // syntax a bit different from other datastructures
			fmt.Println(i)
		} // keeps on listening to data coming from channel infinitly
		wg.Done()
	}(bufch)
	go func(bufch chan<- int) {
		bufch <- 10
		bufch <- 20
		close(bufch) // to avoid getting into a deadlock, close the channel after use
		// bufch <- 20 // not allowed to send data after closing, causes panic
		wg.Done()
	}(bufch)
	wg.Wait()

	fmt.Println("\nchecking whether channel is closed, go routines")
	wg.Add(2)
	bufch = make(chan int, 50)
	go func(bufch <-chan int) {
		for {
			if i, ok := <-bufch; ok {
				fmt.Println(i, ok)
			} else {
				fmt.Println(i, ok)
				break
			}
		}
		wg.Done()
	}(bufch)
	go func(bufch chan<- int) {
		bufch <- 10
		bufch <- 20
		close(bufch)
		wg.Done()
	}(bufch)
	wg.Wait()

	// logger using defer
	fmt.Println("\nlogger, go routines")
	go logger()
	defer func() {
		fmt.Println("\nClosing logCh")
		close(logCh) // one way of gracefully closing the channel
	}()

	logCh <- logEntry{time: time.Now(), severity: logInfo, message: "App is starting"}
	time.Sleep(time.Second)

	logCh <- logEntry{time: time.Now(), severity: logWarning, message: "Something about to go down"}
	time.Sleep(time.Second)

	logCh <- logEntry{time: time.Now(), severity: logError, message: "Something went down"}
	time.Sleep(time.Second)

	// logger using signal channel
	fmt.Println("\nUsing signal channel")
	go selectLogger()

	logChWithSelect <- logEntry{time: time.Now(), severity: logInfo, message: "App is starting"}
	time.Sleep(time.Second)

	logChWithSelect <- logEntry{time: time.Now(), severity: logWarning, message: "Something about to go down"}
	time.Sleep(time.Second)

	logChWithSelect <- logEntry{time: time.Now(), severity: logError, message: "Something went down"}
	time.Sleep(time.Second)

	doneCh <- struct{}{}
	time.Sleep(1 * time.Second)
	// you can try sending data on closed channel, go-runtime will panic
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

func logger() {
	for entry := range logCh {
		fmt.Printf(
			"%v - [%v]%v\n",
			entry.time.Format("2006-01-02T15:04:05"),
			entry.severity,
			entry.message,
		)
	}
}

var logChWithSelect = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // acts as a signal only channel, empty struct = 0 mem allocation

func selectLogger() {
	defer close(logChWithSelect)
	defer fmt.Println("Closing logChWithSelect")
	defer close(doneCh)
	defer fmt.Println("Closing doneCh")
	for {
		select {
		case entry, ok := <-logChWithSelect:
			fmt.Printf(
				"%v\t%v - [%v]%v\n",
				ok,
				entry.time.Format("2006-01-02T15:04:05"),
				entry.severity,
				entry.message,
			)
		case _, ok := <-doneCh:
			fmt.Println("Signal to wrap things up:", ok)
			return
		}
		// entire select statement is blocked until a message is received on one of the case
	}
}
