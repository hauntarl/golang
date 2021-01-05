package main

import "fmt"

/*
	Creating Pointers:
		1. var <var2> *<datatype> = &<var1>
			- * declaring a pointer to data of that type
			- & is addressof operator
		2. <var2> := &<var1>

	Dereferencing Pointers:
		- *<var2> we ask go to look into the value stored in that address
		- complex types (eg. structs) are automatically dereferenced

	The new function:
		- many cases you want a variable to point at a memory location
		- new(<datatype>) creates a new object in memory, initializes it with
		  default value and returns the address of that memory

	Types with internal pointers:
		- all assignment operations in go are copy operations
		- slices, maps (shallow copy, as they have internal pointers)
		- primitives, arrays and structs (deep copy)
*/

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Printf("a = %v, *b = %v\t &a = %v, b = %v\n", a, *b, &a, b) // * dereferencing operator
	a = 27
	fmt.Printf("a = %v, *b = %v\t &a = %v, b = %v\n", a, *b, &a, b)

	fmt.Println()
	var ms1 *myStruct        // example of creating some value in memory
	ms1 = &myStruct{foo: 10} // and making a pointer point to it
	fmt.Println(ms1)         // without having to define an intermediate variable

	fmt.Println()
	var ms2 *myStruct   // performing above operation
	ms2 = new(myStruct) // using the new keyword
	fmt.Println(ms2)    // only initializes the object with default values
	ms2.foo = 20        // (*ms2).foo = 20, as go doesn't allow arithmetic operations on
	fmt.Println(ms2)    // pointers, compiler understands that it needs to dereference and perform

	fmt.Println("\nnil:") // as go initializes every variable with some default
	var ms3 *myStruct     // a pointer which you don't initialize has <nil> value as default
	fmt.Println(ms3)      // we must make sure before using the pointer that it is not nil
}
