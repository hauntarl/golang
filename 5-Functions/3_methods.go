package main

import "fmt"

/*
	Methods:
		- a function which is executed in a given context is called a method
		- methods define the behavior of a structure
		- func (<var> <type>) <func-name>() {}
		- (<var> <type>) - can be of 2 types
		- value receiver - behaves as pass by value, changes made to the object
		  does not reflect, instead a copy of that object is modified
		- pointer receiver - behaves as pass by reference, changes made to the
		  object gets reflected on the same object
*/
func main() {
	g := greeter{
		greeting: "Hello",
		name:     "hauntarl",
	}

	g.greet()

	g.changeGreet("Hey")
	g.greet()

	g = g.changeGreetings("Hey")
	g.greet()

	g.changeName("Sameer")
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

// value receiver
func (g greeter) greet() { fmt.Println(g.greeting, g.name) }

// value receiver is used to change greet message, won't work
func (g greeter) changeGreet(newGreet string) { g.greeting = newGreet }

// value receiver to change greet message, workaround
func (g greeter) changeGreetings(newGreet string) greeter {
	g.greeting = newGreet
	return g
}

// pointer receiver
func (g *greeter) changeName(newName string) { g.name = newName }
