package main

import "fmt"

/*
	Arrays:
		- [<size>]<datatype>{<values>}
		- [5]int{}: create an array of specified size
		- [...]int{1, 2, 3, 4, 5}: create an array large enough to hold given data
		- default value is zero for each element
		- use built-in function len(array) to find the length
		eg. [32]byte
			[2*N] struct { x, y int32 }
			[1000]*float64
			[3][5]int
			[2][2][2]float64  // same as [2]([2]([2]float64))

	Array Copy:
		- array2 := array1 (deep copy)
		- Why? because pointers do exist
		- array2 := &array1 (shallow copy)
*/
func main() {
	fmt.Println("\nARRAY DECLARATION:")
	var grades [3]string = [3]string{"A", "B", "C"}
	scores := [3]int{95}
	marks := [...]int{85, 90}
	fmt.Printf("Grades: %v\nScores: %v\nMarks: %v\n", grades, scores, marks)
	fmt.Printf("Length of Marks array: %v\n", len(marks))

	fmt.Println("\nIDENTITY MATRIX:")
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	fmt.Println("\nARRAY COPY:")
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 4
	fmt.Printf("Deep Copy: a-%v and b-%v\n", a, b)

	c := &a
	c[1] = 4
	fmt.Printf("Shallow Copy: a-%v and c-%v\n", a, c)
}
