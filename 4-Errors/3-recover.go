package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		Recover:
			- when you have a panicking situation you believe you can recover from
			- only useful inside of a deferred functions
			- when application panics, it won't execute rest of the statements
			  but will execute deferred functions
			- if you do recover, the higher functions in call stack will continue execution
			- if you don't want that behavior, then you can re-panick in recover()
	*/
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("I am about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	panic("Something bad happened!")
	fmt.Println("I did not panic")
}
