package main

import "fmt"

/*
	Embedding:
		- go doesn't support object oriented principles like inheritance(IS-A)
		- go uses a model similar to inheritance called composition(HAS)
		- we embed one structure into another, we say that structure A has
		  features of structure B
		- structure A does not become a type of structure B instead it is an
		  independent struct which has no relationship with struture B other
		  than it embeds B
		- A field declared with a type but no explicit field name is called
		  an embedded field.
		- An embedded field must be specified as a type name T or as a pointer
		- to a non-interface type name *T, T itself may not be a pointer type.
		- The unqualified type name acts as the field name.

		struct {
			T1        // field name is T1
			*T2       // field name is T2
			P.T3      // field name is T3
			*P.T4     // field name is T4
			x, y int  // field names are x and y
		}

	INVALID:
		struct {
			T     // conflicts with embedded field *T and *P.T
			*T    // conflicts with embedded field T and *P.T
			*P.T  // conflicts with embedded field T and *T
		}

	NOTE:
		- should not be used for modelling behaviors/methods, it does allow
		  behaviors/methods to carry through, however we can't use them
		  interchangeably is a severe limitation
		- better use interfaces
		- use embedding when you are describing common behavior, not talking about
		  polymorphism or ability to interchangeably use the objects, only to have
		  some common behavior carried forward
*/
func main() {
	fmt.Println("\nEMBEDDING:")
	type animal struct {
		name   string
		origin string
	}
	type bird struct {
		animal
		speed  float32
		canFly bool
	}
	emu := bird{}
	emu.animal = animal{name: "Emu"} // set fields of animal by initiazing animal
	emu.origin = "Australia"         // or directly access fields of animal as bird's fields
	emu.speed = 48                   // this is syntactical sugar as behind the scenes
	emu.canFly = false               // go handles the delegation of the request
	fmt.Println(emu)

	ostrich := bird{
		animal: animal{
			name:   "Ostrich", // if we are using this method to initialize the bird structure
			origin: "Africa",  // we must know the internal structure of bird
		},
		speed:  70,
		canFly: false,
	}
	fmt.Println(ostrich)
}
