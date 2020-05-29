package main

import "fmt"

func main() {
	/*
		Boolean:
			- not an alias for another datatype like integer
			- cannot be converted back and forth with other datatypes
	*/
	fmt.Println("\nBOOLEAN:")
	var boolean bool // defaults to false
	fmt.Printf("%v, %T\n", boolean, boolean)

	/*
		Signed Integer:
			1. int - unspecified size (atleast 32 bits)
			2. int8
			3. int16
			4. int32
			5. int64

		Unsigned Integer:
			1. uint - 32 bits
			2. uint8
			3. uint16
			4. uint32

		Byte:
			- byte: uint8

		NOTE:
			- cannot mix types even in same family (datatype must be same)
			- need to perform explicit type conversion (risk of data loss)
	*/
	fmt.Println("\nINTEGERS:")
	var signedInteger int8
	fmt.Printf("%v, %T\n", signedInteger, signedInteger)
	var unsignedInteger uint16
	fmt.Printf("%v, %T\n", unsignedInteger, unsignedInteger)
	var byteVar byte
	fmt.Printf("Byte %v, %T\n", byteVar, byteVar)

	fmt.Println("\nBIWISE OPERATIONS:")
	a := 10 // defaults to int
	b := 3
	fmt.Printf("%v & %v = %v\t%b & %b = %b\n", a, b, a&b, a, b, a&b)     // bitwise AND
	fmt.Printf("%v | %v = %v\t%b | %b = %b\n", a, b, a|b, a, b, a|b)     // bitwise OR
	fmt.Printf("%v & %v = %v\t%b & %b = %b\n", a, b, a^b, a, b, a^b)     // bitwise XOR
	fmt.Printf("%v &^ %v = %v\t%b &^ %b = %b\n", a, b, a&^b, a, b, a&^b) // bitwise AND-NOT (Bit-clear)
	fmt.Printf("%v << 3 = %v\t%b << 11 = %b\n", a, a<<3, a, a<<3)        // bitwise left-shift
	fmt.Printf("%v >> 3 = %v\t%b >> 11 = %b\n", a, a>>3, a, a>>3)        // bitwise right-shift

	/*
		Float:
			1. float32
			2. float64

		NOTE:
			- if mantissa exceeds the negative limit then value is set to zero,
			  for positive limit, compiler throws an error
			- modulus operation not allowed on floating point numbers
	*/
	fmt.Println("\nFLOAT:")
	c := 3.14                // defaults to float64
	var d float32 = 13.7E-72 // explicity defining a float 32
	e := 13.7e72
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)

	/*
		Complex:
			1. complex64 (real - float32, imaginary - float32)
			2. complex128 (real - float64, imaginary - float64)

		NOTE:
			- to generate use complex(real, imag)
			- to decompose use real(complex), imag(complex)
			- same arithmetic operations as float
	*/
	fmt.Println("\nCOMPLEX:")
	var f complex64 // defaults to 0 + 0i
	fmt.Printf("%v, %T\n", f, f)
	g := complex(4, 5) // defaults to complex128
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", real(g), real(g))
	fmt.Printf("%v, %T\n", imag(g), imag(g))

	/*
		Texts:
			1. string - UTF8
				- str := "string value"
				- immutable
				- can be treated as an array of bytes
				- can be converted back and forth with byte
				- []byte(string) - byte slicing to get array of UTF8 codes
				- + operation for concatenation
			2. rune - UTF32
				- rn := 'rune value'
				- type alias for int32
				- special methods required to process
	*/
	fmt.Println("\nTEXTS:")
	h := "hello world"
	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\tbyte at index: 6\n", h[6], h[6]) // cannot modify the value
	fmt.Printf("%v, %T\tbyte coverted to string\n", string(h[6]), string(h[6]))
	i := []byte(h) // byte slicing
	fmt.Printf("byte array: %v, %T\n", i, i)
	j := 'a' // rune datatype
	fmt.Printf("Rune '%v' - %v, %T\n", string(j), j, j)
}
