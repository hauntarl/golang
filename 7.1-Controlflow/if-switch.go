package main

import (
	"fmt"
	"math"
)

func main() {
	// most basic if syntax
	fmt.Println("\nsimple if block:")
	if true {
		fmt.Println("Is true")
	}

	// if initializer syntax
	fmt.Println("\nif initializer block")
	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	// if <initialization>; <condition> {<block>}
	if data, ok := statePopulation["Texas"]; ok {
		// scope of variable "data" is limited to this if block only
		fmt.Println("Texas population: ", data)
	}

	/*
		Comparison operators:
			- <, >, <=, >=, ==, !=
			- these work with all numeric types
			- they don't work with string types
			- for string type and reference types only == and != works

		Logical operators:
			- ||, && also known as short-circuit operators
			- ! operator (takes a boolean and flips it to the other side)
	*/
	fmt.Println("\nsimple if-else block")
	number := 50
	guess := 30
	if guess < 1 || guess > 100 {
		fmt.Println("guess should be in between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("You got it")
		}
	}

	// Dealing with floats in equality tests
	fmt.Println("\nequality of floats\ntest 1")
	num := 0.123
	if num == math.Pow(math.Sqrt(num), 2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	// go deals with approximate values when it comes to float
	// so we need to be careful while performing float operations in equality

	fmt.Println("test 2")
	if math.Abs(num/math.Pow(math.Sqrt(num), 2)-1) < 0.001 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	// above example takes the anomalies into consideration
	// by using error-parameter as a difference of one-tenth of a percentage

	/*
		Switch Statements:
			- go allows multiple tests in a single case
			- tests in each case should be unique
			- similar to if statements, we can also use the initializer syntax for tag
			- go doesn't provide falling through mechanism in switch by default,
				- that means break statement is implicit and if you don't want this
			  	  behavior then use "fallthrough" keyword
				- fallthrough is logic-less, irrespective of whether next case passes
					or fails, go will execute the expressions inside the next case
			- there might be some cases where depending on the situation you might want
			  to leave the case early, you can use the break statement for it
	*/
	fmt.Println("\nsimple switch statement:")
	fmt.Println("simple tag syntax")
	switch 5 {
	case 1:
		fmt.Println("\tone")
	case 2:
		fmt.Println("\ttwo")
	case 3, 4, 5:
		fmt.Println("\tthree, four or five")
	default:
		fmt.Println("\tunexpected value")
	}

	fmt.Println("initializer tag syntax")
	switch i := 2 + 3; i { // scope of variable i limited to the switch block
	case 1:
		fmt.Println("\tone", i)
	case 2:
		fmt.Println("\ttwo", i)
	case 3, 4, 5:
		fmt.Println("\tthree, four or five", i)
	default:
		fmt.Println("\tunexpected value", i)
	}

	fmt.Println("tag list syntax")
	i := 15
	// allows overlapping of cases, executes the first case which satisfies the condition
	switch {
	case i <= 10:
		fmt.Println("\tless than or equal to 10")
	case i <= 20:
		fmt.Println("\tless than or equal to 20")
		fallthrough
	default:
		fmt.Println("\tgreater than 20")
	}

	fmt.Println("\ntype switch")
	var j interface{} = 10 // variable j is of type interface, which can take any datatype in go
	switch j.(type) {      // <var>.(type) can only be used in type switch, returns type of data
	case int:
		fmt.Println("\tint")
	case float64:
		fmt.Println("\tfloat64")
	case string:
		fmt.Println("\nstring")
	default:
		fmt.Println("\tanother type")
	}
}
