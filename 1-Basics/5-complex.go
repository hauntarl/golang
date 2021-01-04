package main

import "fmt"

/*
	Complex:
		1. complex64 (real - float32, imaginary - float32)
		2. complex128 (real - float64, imaginary - float64)

	NOTE:
		- to generate use complex(real, imag)
		- to decompose use real(complex), imag(complex)
		- same arithmetic operations as float
*/
func main() {
	fmt.Println("\nCOMPLEX:")
	var f complex64 // defaults to 0 + 0i
	cout(f)
	g := complex(4, 5) // defaults to complex128
	cout(g)
	cout(real(g))
	cout(imag(g))
}

func cout(val interface{}) { fmt.Printf("%v, %T\n", val, val) }
