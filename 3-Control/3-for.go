package main

import "fmt"

/*
	For Loops:
		1. for <initializer>; <condition>; <incrementor> {}
		2. for <condition> {}
		3. for {}
		4. for <key>, <value> := range <collection> {}
			- arrays, slices, maps, strings and channels
			- in case of channels, we only get one parameter: <value>

	Exiting early?
		1. break
		2. continue
		3. labels
*/
func main() {
	fmt.Println("\nsimple for loop:")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	// increment/decrement operator is not an expression, it is a statement in itself
	// so go does not allow it to be mixed with other statements

	// initializing multiple values at the same time
	fmt.Println("\nfor loop with multiple values")
	for i, j := 0, 5; i < 5; i, j = i+1, j-1 {
		fmt.Print("i:", i, ", j:", j, "\t")
	}

	fmt.Println("\nwhile loop in go")
	i := 0
	for i < 5 { // psych! go doesn't have one
		fmt.Print(i, " ")
		i++
	}

	fmt.Println("\ninfinite loop in go")
	for {
		if i == 10 {
			break // also go has "continue" keyword like other languages
		}
		fmt.Print(i, " ")
		i++
	}

	fmt.Println("\nlabelled loop in go")
	// why? when we are using nested for loops, the "break" keyword, exits the closest
	// loop it can find. In order to break out of outerloops we use labelled loops
Loop: // if label declared, must be utilized or compiler throws an error
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Print(i*j, " ")
			if i*j > 3 {
				break Loop
			}
		}
	}

	fmt.Println("\ncollections with for loop\nslice")
	slice := []int{1, 2, 3}
	// the for range loop, also works with strings
	for k, v := range slice {
		fmt.Print("pair: ", k, v, "\t")
	}

	fmt.Println("\nmap")
	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}
	for k, v := range statePopulation {
		fmt.Print("pair: ", k, ", ", v, "\t")
	}
}
