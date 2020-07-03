package main

import (
	"errors"
	"fmt"
)

func sayHello() {
	fmt.Println("Hello World!")
}

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

func main() {
	/*
		Basic Syntax:
			- func <name>() {<body...>}
			- <name> - Pascal or camelCase, depending on whether you want to export or not
	*/
	sayHello()

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
	fmt.Println("\nBasic syntax:")
	for i := 1; i <= 2; i++ {
		sayMessage("Hello", i)
	}

	sum(1, 2, 3, 4, 5) // variatic paramters

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

	result, err := div(5, 0) // put second parameter as 0 to test
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
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
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
