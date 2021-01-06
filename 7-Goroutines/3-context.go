package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Hello again"
	go func(msg string) {
		fmt.Println(msg)
		// providing a variable in local context for execution of this function
		// stack we eliminate the issue we faced earlier, will work only if we
		// use pass by value
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
	// we should avoid using Sleep calls in production as they can be very
	// unreliable, as you need your application to sleep relative to the
	// average performance of your application and not the system clock
}
