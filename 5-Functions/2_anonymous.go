package main

import "fmt"

/*
	Anonymous functions:
		1. func() {}() - known as IIFE in Js - Immediately Invoked Function
		   Expression
		2. Go lets you define function as a variable then you can pass around
		   the function, they are treated as first class citizens
*/
func main() {
	fmt.Println("\nAnonymous function:")
	func() { fmt.Println("I am Anonymous") }()

	var f func() = func() { fmt.Println("I am a first class citizen") }
	caller(f) // treating a function as a normal variable and passing it around

	foo := func() { fmt.Println("I am assigned using a short-hand operator") }
	fmt.Println(foo) // returns the address in memory for the function
	foo()

	// declaring a variable divide of type function
	var divide func(float64, float64) (float64, error) // defaults to nil
	fmt.Println(divide)

	// initializing it
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}

	result, err := divide(5, 0) // put second parameter as 0 to test
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func caller(f func()) { f() }
