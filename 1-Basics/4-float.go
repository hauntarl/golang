package main

import "fmt"

/*
	Float:
		1. float32
		2. float64
	NOTE:
		- if mantissa exceeds the negative limit then value is set to zero,
		  for positive limit, compiler throws an error
		- modulus operation not allowed on floating point numbers

	Valid literals:
		0.
		72.40
		072.40       // == 72.40
		2.71828
		1.e+0
		6.67428e-11
		1E6
		.25
		.12345E+5
		1_5.         // == 15.0
		0.15e+0_2    // == 15.0

		0x1p-2       // == 0.25
		0x2.p10      // == 2048.0
		0x1.Fp+0     // == 1.9375
		0X.8p-0      // == 0.5
		0X_1FFFP-16  // == 0.1249847412109375
		0x15e-2      // == 0x15e - 2 (integer subtraction)

	Invalid literals:
		0x.p1        // invalid: mantissa has no digits
		1p-2         // invalid: p exponent requires hexadecimal mantissa
		0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent
		1_.5         // invalid: _ must separate successive digits
		1._5         // invalid: _ must separate successive digits
		1.5_e1       // invalid: _ must separate successive digits
		1.5e_1       // invalid: _ must separate successive digits
		1.5e1_       // invalid: _ must separate successive digits

	Refer Go language specification for more info: https://golang.org/ref/spec
*/
func main() {
	fmt.Println("\nFLOAT:")
	c := 3.14                // defaults to float64
	var d float32 = 13.7e-72 // explicity defining a float 32
	e := 13.7e72
	cout(c)
	cout(d)
	cout(e)
}

func cout(val interface{}) { fmt.Printf("%v, %T\n", val, val) }
