package main

import "fmt"

func sayMessage(message string, index int) {
	fmt.Println(message, index)
}

// arguments with same datatype can be comma separated, as it becomes less verbose
func sayGreeting(greet, name string) {
	fmt.Println(greet, name)
}

func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, value := range values {
		result += value
	}
	fmt.Println("Sum is", result)
}

func whatAmI() *string {
	response := "I am go"
	fmt.Println(&response, response)
	return &response
}

func mul(a, b int) (result int) {
	result = a * b
	return
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type greet struct {
	greeting string
	name     string
}

func (g greet) greeter() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	/*
		Basic Syntax:
			- the main() function itself
			- func <name>() {<body...>} - basic syntax
			- <name> - Pascal or camelCase, depending on whether you want to export or not
	*/
	fmt.Println("\nBasic syntax:")
	for i := 1; i <= 2; i++ {
		sayMessage("Hello", i)
	}

	/*
		Parameters:
			- func <name>(<param1> <datatype>, ...) {<body...>}
			- positional arguments
			- pass by value
			- pass by reference (using pointers)
			- variatic parameters:
				- func <name>(<var-name> ...<datatype>) {}
				- wraps all arguments in a slice
				- you can only have one and it should be the last one
	*/
	fmt.Println("\nParameters:")
	sayGreeting("Hi", "hauntarl")
	sum(1, 2, 3, 4, 5)

	/*
		Return values:
			- func <name>() <return-type> {<body...>}
			- you can also return a pointer
			- when go realizes that you are returning addressof a local variable
			- it will automatically promote the memory from stack to heap
			- named return:
				- func <name>() (<return-var> <return-type>) {}
				- <return-var> is locally available and you can perform operations on it
				- by writing just the return keyword will return the <return-var>
				- syntactic sugar
			- multiple returns
	*/
	fmt.Println("\nReturn:")
	response := whatAmI()
	fmt.Println(response, *response)

	fmt.Println(mul(4, 5))

	result, err := div(5, 1) // put second parameter as 0 to test
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	/*
		Anonymous functions:
			- func() {}()
			- known as IIFE in Js
			- lets you define function as a variable then you can pass around the function
			- functions are treated as first class citizens
	*/
	fmt.Println("\nAnonymous function:")
	func() {
		fmt.Println("I am Anonymous")
	}()

	var f func() = func() {
		fmt.Println("I am also a first class citizen")
	}
	f()

	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}
	result, err = divide(5, 1) // put second parameter as 0 to test
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	/*
		Methods:
			- a function which is executed in a given context is called a method
			- func (<var> <type>) <func-name>() {}
			- (<var> <type>) - value receiver can accept both a copy or a reference
	*/
	fmt.Println("\nMethods")
	greetings := greet{
		greeting: "Hello",
		name:     "hauntarl",
	}
	greetings.greeter()
}
