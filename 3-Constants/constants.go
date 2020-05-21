package main

import "fmt"

const a1 = iota // assigned value 0

const (
	// compiler understands the pattern from value assigned to first variable
	a2 = iota - 5 // assigned value -5
	b2            // assigned value -4
	c2            // assigned value -3
)
const (
	a3 = iota + 5 // assigned value 5
	b3            // assigned value 6
)

// example of iota with bitshift
const (
	_  = iota             // ignored
	KB = 1 << (10 * iota) // left shift 1 by 10 bits, value assigned 1024 (bytes)
	MB                    // left shift 20 bits
	GB                    // left shift 30 bits
	TB
	PB
	EB
	ZB
	YB
)

// example of checking user roles
const (
	isAdmin            = 1 << iota // value assigned in bits: 00000001
	isHeadquarters                 // 00000010
	canSeeFinancials               // 00000100
	canSeeAfrica                   // 00001000
	canSeeAsia                     // 00010000
	canSeeNortAmerica              // 00100000
	canSeeSouthAmerica             // 01000000
	/*
		- encoding 7 access roles in a single byte of data
		- to verify access of a particular user for a particular action
		- we Bitwise-OR '|' the required roles for that action and get access byte
		- then perform Bitwise-AND '&' with user byte
		- if result equals the access byte then permission is granted
	*/
)

func main() {
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
			- value must be calculable at the compile time
			- unlike variables constants can be initialised and never be used in a block
	*/
	fmt.Println("\nCONSTANTS")
	const typedConstant int = 10
	var x int16 = 30
	fmt.Printf("%v, %T\n", typedConstant, typedConstant)
	//fmt.Printf("%v, %T", myConst + x, myConst + x) // compiler error

	const inferredConstant = 20
	fmt.Printf("%v, %T\n", inferredConstant+x, inferredConstant+x) // returns variables datatype

	/*
		Enumerated Constants:
			1. const a = iota (special keyword initiates a counter)
			2. const (
					a = iota (assigns 0)
					b = iota (assigns 1)
					c = iota (assigns 2)
			   ) the value of iota is scoped to the const block

		Use Case:
			- used as enums(from languages like Java)
			- const (
				catSpecialist = iota
				dogSpecialist
				snakeSpecialist
			  )

		NOTE:
			- iota starts with value 0 and increments by 1 upon each read
			- if you want to have a different starting point
			- we can do all the operations that are allowed on integers
			- arithmetic, bitwise, bit-shifting
	*/
	fmt.Println("\nENUMERATED CONSTANTS")
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("c2: %v\n", c2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("b3: %v\n", b3)

	var specialistType int // if not initialised, value is 0
	const (
		// '_' is a write only variable, used when you don't care about the value assigned
		// compiler doesn't assign any memory to it
		_               = iota // specialistType will point to 0 if not initialised
		catSpecialist          // if you want to avoid variable pointing at any particular field
		dogSpecialist          // use a dummy variable to initiate iota
		snakeSpecialist        // now the actual values start from one
	)
	fmt.Printf("specialistType == catSpecialist? %v\n", specialistType == catSpecialist)

	fmt.Println("\nFile Size Manipulation")
	fileSize := 4000000000.
	fmt.Printf("file size in B: %.2f\n", fileSize)
	fmt.Printf("file size in KB: %.2f\n", fileSize/KB)
	fmt.Printf("file size in MB: %.2f\n", fileSize/MB)
	fmt.Printf("file size in GB: %.2f\n", fileSize/GB)

	fmt.Println("\nUser roles example")
	var userByte byte = canSeeFinancials | canSeeAfrica | canSeeAsia
	fmt.Printf("Can user see the financials? %v\n", userByte&canSeeFinancials == canSeeFinancials)
	fmt.Printf("Is the user an admin? %v\n", userByte&isAdmin == isAdmin)
	var accessByteForAfricaAndAsia byte = canSeeAfrica | canSeeAsia
	fmt.Printf("Can user see Africa and Asia? %v\n", userByte&accessByteForAfricaAndAsia == accessByteForAfricaAndAsia)
}
