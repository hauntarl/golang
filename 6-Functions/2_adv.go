package main

import "fmt"

func main() {
	/*
		Anonymous functions:
			- func() {}()
			- known as IIFE in Js - Immediately Invoked Function Expression
			- lets you define function as a variable then you can pass around the function
			- functions are treated as first class citizens
	*/
	fmt.Println("\nAnonymous function:")
	func() {
		fmt.Println("I am Anonymous")
	}()

	var f func() = func() {
		fmt.Println("I am a first class citizen")
	}
	f()

	// declaring a variable divide of type function
	var divide func(float64, float64) (float64, error) // defaults to nil
	// initializing it
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}
	result, err := divide(5, 1) // put second parameter as 0 to test
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	/*
		Methods:
			- a function which is executed in a given context is called a method
			- methods define the behavior of a structure
			- func (<var> <type>) <func-name>() {}
			- (<var> <type>) - value receiver can accept both a copy or a reference
	*/
	fmt.Println("\nMethods")
	g := greeter{
		greeting: "Hello",
		name:     "hauntarl",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
