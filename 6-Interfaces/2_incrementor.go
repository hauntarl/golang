package main

import "fmt"

// incrementer provides the definition of increment method
type incrementer interface{ increment() int }

// user-type intCounter is a new datatype whose underlying type is int
type intCounter int

// intCounter implicitly implements the incrementer interface
func (ic *intCounter) increment() int {
	*ic++
	// explicity type casting from intCounter to int, as Go does not perform
	// implicit type casting
	return int(*ic)
}

func main() {
	fmt.Println(*new(intCounter))
	var myIncrementer incrementer = new(intCounter)
	for i := 0; i < 5; i++ {
		fmt.Print(myIncrementer.increment(), " ")
	}
}
