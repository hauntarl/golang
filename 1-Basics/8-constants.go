package main

import "fmt"

/*
	Constants:
		1. Typed constants
			- const myConst int = 10
			- explicity defining the datatype
			- can interoperate only with same type
			- work like immutable variables
		2. Type Infered constants
			- const myConst = 20
			- infers the datatype from value
			- can interoperate with similar types
			- work like literals
	NOTE:
		- immutable, but can be shadowed
		- value must be calculable at compile time
		- unlike variables constants can be initialised and never be used in a block
*/
func main() {
	fmt.Println("\nCONSTANTS")
	const typedConstant int = 10
	var x int16 = 30
	cout(typedConstant)
	//fmt.Printf("%v, %T", typedConstant+x, typedConstant+x) // compiler error

	const inferredConstant = 20
	cout(inferredConstant + x) // returns variables datatype
}

func cout(val interface{}) { fmt.Printf("%v, %T\n", val, val) }
