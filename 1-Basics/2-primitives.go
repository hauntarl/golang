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
		
		Valid literals:
			42
			4_2
			0600
			0_600
			0o600
			0O600
			0xBadFace
			0xBad_Face
			0x_67_7a_2f_cc_40_c6
			170141183460469231731687303715884105727
			170_141183_460469_231731_687303_715884_105727

		Invalid literals:
			_42         an identifier, not an integer literal
			42_         invalid: _ must separate successive digits
			4__2        invalid: only one _ at a time
			0_xBadFace  invalid: _ must separate successive digits

		Refer Go language specification for more info: https://golang.org/ref/spec
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
	fmt.Printf("%v ^ %v = %v\t%b ^ %b = %b\n", a, b, a^b, a, b, a^b)     // bitwise XOR
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
				- raw string literals - char sequence between back quotes ` `,
				  backslash \ has no special meaning in these

			Valid literals:
				`abc`                // same as "abc"
				`\n
				\n`                  // same as "\\n\n\\n"
				"\n"
				"\""                 // same as `"`
				"Hello, world!\n"
				"日本語"
				"\u65e5本\U00008a9e"
				"\xff\u00FF"

			Invalid literals:
				"\uD800"             // illegal: surrogate half
				"\U00110000"         // illegal: invalid Unicode code point

		There are four ways to represent the integer value as a numeric constant:
			- \x followed by exactly two hexadecimal digits (limit < 0x10FFFF)
			- \u followed by exactly four hexadecimal digits 
			- \U followed by exactly eight hexadecimal digits
			- \ followed by exactly three octal digits. (0-255)
		Although these representations all result in an integer, they have different valid ranges.
		Further they also allow special escape characters

			2. rune - UTF32
				- rn := '<single rune char>'
				- type alias for int32
				- special methods required to process
			
			Valid literals:
				'a'
				'ä'
				'本'
				'\t'
				'\000'
				'\007'
				'\377'
				'\x07'
				'\xff'
				'\u12e4'
				'\U00101234'
				'\''         // rune literal containing single quote character
			
			Invalid literals:
				'aa'         // illegal: too many characters
				'\xa'        // illegal: too few hexadecimal digits
				'\0'         // illegal: too few octal digits
				'\uDFFF'     // illegal: surrogate half
				'\U00110000' // illegal: invalid Unicode code point

		Refer Go language specification for more info: https://golang.org/ref/spec
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
