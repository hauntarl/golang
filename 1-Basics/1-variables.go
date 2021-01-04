package main

import (
	"fmt"
	"strconv"
)

/*
	VARIABLE DECLARATION:
		1. var i int (defaults to 0)
		2. var i int = 0
		3. var i = 0
		4. i := 0 (type inference)(not allowed when declaring global variables)

	VISIBILTY:
		1. no private scope
		2. local variables are block scoped
		3. lower case global variables have package scope
		4. upper case variables are exported

	NAMING CONVENTION:
		1. Pascal for Exporting variables
		2. camelCase for rest
		3. Capitalize acronyms (theHTTP, theURL)
*/

// K is exported (good practice to add comment before exported variable)
var K int // global variable, scope: exported outside package

var k int // global variable, scope: current package

var ( // declaring mutliple global variables at once
	author string = "hauntarl"
	age    int    = 22
	theDOB        = 1997 // acronyms in capital
)

func main() {
	// using global variables
	fmt.Printf(
		"author: %v, age: %v, dob: %v, exported K: %v\n",
		author, age, theDOB, K,
	)

	i := 420 // local variable, scope: block
	cout(i)

	var j float32  // if declared, local variable must be used.
	j = float32(i) // requires explicit conversion to make sure no data is lost
	cout(j)

	/*
		TYPE CONVERSION:
			1. destinationType(variable)
			2. use strconv package to convert numbers to string
			3. alternatively, use fmt.Sprintf() to convert variables to string
			NOTE: no implicit type conversion
	*/
	k := string(i) // will convert number to its respective unicode character
	cout(k)

	l := strconv.Itoa(i) // import strconv package to convert number to string
	cout(l)

	m := fmt.Sprintf("%v", true)
	cout(m)
}

func cout(val interface{}) { fmt.Printf("%v, %T\n", val, val) }
