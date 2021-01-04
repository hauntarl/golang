package main

import "fmt"

/*
	Slices:
		- []<datatype>{<values>}
		- built-in function len(slice) to find the length
		- built-in function cap(slice) actual length of underlying array
		- value of an uninitialized slice is nil

	Slice Copy:
		- slice2 := slice1 (shallow copy)
		- slice2 := slice1[:] (slice of all elements)
		- slice2 := slice1[3:] (slice from 4th element to end)
		- slice2 := slice1[:6] (slice of first 6 elements)
		- slice2 := slice1[3:6] (slice of elements between 4 to 6)
		- NOTE: <inclusive>:<exclusive>

	Built-in functions:
		1. slice := make([]<datatype>, <length>, <capacity>)
			- Unlike array, slice doesn't have fixed size
			- slices have the capability to add and remove elements from them
			- make([]int, 50, 100) is equivalent to new([100]int)[0:50]
		2. slice = append(<source>, <elements...>) - vargs
			- if underlying array does not have the capacity to store the new element
			- it creates a new array of larger size
			- copies elements from previous array
			- then appends the new element
			- this operation can be very expensive
*/
func main() {
	fmt.Println("\nSLICES:")
	d := []int{1, 2, 3, 4, 5}
	e := d
	e[1] = 6
	fmt.Printf("d: %v\ne: %v\n", d, e)
	fmt.Printf("Length of d: %v\n", len(d))
	fmt.Printf("Capacity of d: %v\n", cap(d))

	fmt.Println("\nRange slicing:")
	f := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	g := f[:] // slicing operations can have their source as an array or slice
	h := f[3:]
	i := f[:6]
	j := f[3:6]
	f[5] = 10 // change should be reflected on all of them
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println("\nBUILT-IN FUNCTIONS:")
	fmt.Println("Make Slice:")
	slice := make([]int, 0, 10)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("Length of slice: %v\n", len(slice))
	fmt.Printf("Capacity of slice: %v\n", cap(slice))

	fmt.Println("\nAppend:\nInitial slice")
	k := []int{} // slice and underlying array both of size 0
	fmt.Printf("k: %v\n", k)
	fmt.Printf("Length of k: %v\n", len(k))
	fmt.Printf("Capacity of k: %v\n", cap(k))

	fmt.Println("\nAfter append")
	// k = append(k, 1, 2, 3, 4) // append function accepts variable arguments
	k = append(k, []int{1, 2, 3, 4, 5, 6}...) // use the (...) operator to spread slice
	fmt.Printf("k: %v\n", k)                  // also known as spread operator
	fmt.Printf("Length of k: %v\n", len(k))
	fmt.Printf("Capacity of k: %v\n", cap(k))

	fmt.Println("\nOperations on slice")
	slice = append(slice, k...)
	fmt.Printf("Initial slice: %v\n", slice)
	slice = slice[1:] // removing first element
	fmt.Printf("Remove first: %v\n", slice)
	slice = slice[:len(slice)-1] // removing last element
	fmt.Printf("Remove last: %v\n", slice)
	slice = append(slice[:len(slice)/2], slice[len(slice)/2+1:]...) // removing an element from middle
	fmt.Printf("Remove middle: %v\n\n", slice)

	// how the underlying length and capacity changes
	slice = make([]int, 5, 10)
	cout(slice)
	slice1 := slice // remains same
	cout(slice1)
	slice2 := slice[:] // remains same
	cout(slice2)
	slice4 := slice[:3] // underlying length changes
	cout(slice4)
	slice3 := slice[3:5] // along with the length, underlying capacity changes
	cout(slice3)

	arr := [...]int{1, 2, 3, 4, 5}
	s := arr[:3]
	cout(s)
	// for slices, the upper index bound is the slice capacity cap(a) rather than the length
	s1 := s[:cap(s)]
	cout(s1)

	// full slice expression a[low:high:max]
	s2 := arr[:2:3]
	cout(s2)
}

func cout(slice []int) {
	fmt.Printf("slice: %v,\nlen: %v, cap: %v\n\n", slice, len(slice), cap(slice))
}
