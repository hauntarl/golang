package main

import "fmt"

/*
	Boolean:
		- not an alias for another datatype, like integer
		- cannot be converted back and forth with other datatypes
*/
func main() {
	fmt.Println("\nBOOLEAN:")
	var boolean bool // defaults to false
	fmt.Printf("%v, %T\n", boolean, boolean)
}
