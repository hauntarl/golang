package main

import "fmt"

/*
	Structures:
		- unlike collections like arrays, slices and maps
		  structures are flexible
		- to access data from structure we use "." syntax
		- structures are pass by value like arrays (deep copy)
		- just like arrays we can use "&" operator to make variable point
		  to same location in memory (shallow copy)

		An empty struct.
		struct {}

		A struct with 6 fields.
		struct {
			x, y int
			u float32
			_ float32  // padding
			A *[]int
			F func()
		}
*/

// Doctor is exported
type Doctor struct {
	number     int      // if struct is exported then
	name       string   // fields starting with lower-case are not visible
	episodes   []string // inorder to export fields, use Pascal case
	Companions []string // only Companions field is visible to other packages
}

func main() {
	fmt.Println("\nSTRUCTURES:")
	aDoctor := Doctor{
		number: 3,
		name:   "John Pertwee",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println("aDoctor", aDoctor)
	/*
		Using Structure:
		1. named arguments
			- explicitly specifying data for each field with field-name
			- code becomes more robust and safe as changes to underlying
			  structure won't affect previous implementation
		2. positional arguments
			- entering values in the sequence in which structure is defined
			- not recommended because it can become a maintainance problem
			- will create conflicts if structure definition changes
	*/

	bDoctor := Doctor{}
	bDoctor.number = 69
	bDoctor.name = "hauntarl"
	bDoctor.Companions = []string{"private"}
	fmt.Println("bDoctor", bDoctor)

	/*
		Anonymous Structures:
			- anonymousStruct := struct{<structure-definition>}{<structure-initialiser>}
			- very few use cases, can be used to format responses
	*/
	fmt.Println("\nANONYMOUS STRUCTURES:")
	anonymousStruct := struct{ name string }{name: "Anonymous Structure"}
	fmt.Printf("anonymousStruct: %v, type: %T\n", anonymousStruct, anonymousStruct)
}
