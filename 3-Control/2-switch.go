package main

import "fmt"

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
func main() {
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
	// an alternative to if-else syntax
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
	var j interface{} = 10 // variable j is an empty interface, which can take any datatype in go
	switch j.(type) {      // <var>.(type) can only be used in type switch, returns type of data
	case int:
		fmt.Println("\tint")
	case float64:
		fmt.Println("\tfloat64")
	case string:
		fmt.Println("\tstring")
	default:
		fmt.Println("\tanother type")
	}
}
