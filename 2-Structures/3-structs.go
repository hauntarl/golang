package main

import (
	"fmt"
	"reflect"
)

// Doctor is exported
type Doctor struct {
	number     int      // if struct is exported then
	name       string   // fields starting with lower-case are not visible
	episodes   []string // inorder to export fields, use Pascal case
	Companions []string // only Companions field is visible to other packages
}

func main() {
	/*
		Structures:
			- unlike collections like arrays, slices and maps
			  structures are flexible
			- to access data from structure we use "." syntax
			- structures are pass by value like arrays (deep copy)
			- just like arrays we can use "&" operator to make variable point
			  to same location in memory (shallow copy)

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

	/*
		Embedding:
			- go doesn't support object oriented principles like inheritance(IS-A)
			- go uses a model similar to inheritance called composition(HAS)
			- we embed one structure into another, we say that structure A has
			  features of structure B
			- structure A does not become a type of structure B instead it is an
			  independent struct which has no relationship with struture B other
			  than it embeds B

		NOTE:
			- should not be used for modelling behaviors/methods, it does allow
			  behaviors/methods to carry through, however we can't use them
			  interchangeably is a severe limitation
			- better use interfaces
			- use embedding when you are describing common behavior, not talking about
			  polymorphism or ability to interchangeably use the objects, only to have
			  some common behavior carried forward
	*/
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

	/*
		Tags:
			- when defining a structure we can attach a tag or a simple message to each field
			- tag can be any arbitrary messsage and go does not use those tags in any way
			- to get the tag information for any particular field we need to import "reflect"
			- it can look in the structure definition and fetch tags for respective fields
			- in order to make use of tags, you need to create your own validation library
			  which will have its own logic to look for keywords in tags
	*/
	fmt.Println("\nTAGS:")
	type simple struct {
		name    string `required:"true" max:"100"`
		address string // in the above tag, linter will give a warning if
		theDOB  string // space separated key:"value" is not present
	}
	t := reflect.TypeOf(simple{})
	field, ok := t.FieldByName("name")
	fmt.Println("Field Information:", field)
	fmt.Printf("tag: %v, isPresent: %v", field.Tag, ok)
}
