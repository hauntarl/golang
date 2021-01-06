package main

import (
	"fmt"
	"sync"
)

// WaitGroup - synchronizes goroutines, joins them
var (
	wg      = sync.WaitGroup{}
	counter = 0
)

func main() {
	msg := "Hello again"
	wg.Add(1) // adds number of goroutines we want our application to wait for
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // decrements the goroutine count by 1
	}(msg)
	msg = "Goodbye"
	wg.Wait() // waits until number is 0

	fmt.Println("\nIssue with only using WaitGroup\nSpawning 10 go routines:")
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHi()
		go increment()
	}
	wg.Wait() // goroutines are racing against each other, no synchronization
}

func sayHi() {
	fmt.Println("Hi", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
