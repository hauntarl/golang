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
	if data, ok := statePopulation["Texas"]; ok {
		// if <initialization>; <condition> {<block>}
		// scope of variable "data" is limited to this if block only
		// and also in following if-else/else block
		fmt.Println("Texas population: ", data)
	}

	/*
		Comparison operators:
			- <, >, <=, >=, ==, !=
			- these work with all numeric types
			- they don't work with string types
			- for string and reference types, only == and != works

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
	sqrt := math.Sqrt(num)
	pow2 := math.Pow(sqrt, 2)
	if math.Abs(num/pow2-1) < 0.001 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	// above example takes the anomalies into consideration
	// by using error-parameter as a difference of one-tenth of a percentage
}
