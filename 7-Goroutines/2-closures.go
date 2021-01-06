package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Hello again"
	go func() {
		fmt.Println(msg) // go has closures
		// function can access variables from its outer scope even if we are
		// creating a different execution stack, runtime knows where to look up
		// the variable so any changes made to the msg variable before goroutine
		// is executed, will be reflected at runtime
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
	msg = "Hi"
}
