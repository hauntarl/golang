package main

import "fmt"

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
		- if you want to have a different starting point, we can do all the
		  operations that are allowed on integers, like: arithmetic, bitwise,
		  bit-shifting
*/

const a1 = iota // assigned value 0

const (
	// compiler understands the pattern from value assigned to the first variable
	a2 = iota - 5 // assigned value -5
	b2            // assigned value -4
	c2            // assigned value -3
)

const (
	a3 = iota + 5 // assigned value 5
	b3            // assigned value 6
)

func main() {
	fmt.Println("\nENUMERATED CONSTANTS")
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("c2: %v\n", c2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("b3: %v\n", b3)

	// different applications for enumerated constants
	specialist()
	fileSize()
	userRoles()
}

func specialist() {
	const (
		// '_' is a write only variable, used when you don't care about the value assigned
		// compiler doesn't assign any memory to it
		_               = iota // specialistType will point to 0 if not initialised
		catSpecialist          // if you want to avoid variable pointing at any
		dogSpecialist          // particular field use a dummy variable to initiate
		snakeSpecialist        // iota, now the actual values start from one
	)
	var specialistType int // if not initialised, value is 0
	fmt.Printf("\nspecialistType == catSpecialist? %v\n", specialistType == catSpecialist)
}

// example of iota with bitshift - file size manipulation
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

func fileSize() {
	fmt.Println("\nFile Size Manipulation")
	fileSize := 4000000000.
	fmt.Printf("file size in B: %.2f\n", fileSize)
	fmt.Printf("file size in KB: %.2f\n", fileSize/KB)
	fmt.Printf("file size in MB: %.2f\n", fileSize/MB)
	fmt.Printf("file size in GB: %.2f\n", fileSize/GB)
}

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

func userRoles() {
	fmt.Println("\nUser roles example")
	var userByte byte = canSeeFinancials | canSeeAfrica | canSeeAsia
	fmt.Printf("Can user see the financials? %v\n", userByte&canSeeFinancials == canSeeFinancials)
	fmt.Printf("Is the user an admin? %v\n", userByte&isAdmin == isAdmin)
	var accessByte byte = canSeeAfrica | canSeeAsia
	fmt.Printf("Can user see Africa and Asia? %v\n", userByte&accessByte == accessByte)
}
