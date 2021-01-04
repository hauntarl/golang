package main

import "fmt"

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

func main() {
	fmt.Println("\nINTEGERS:")
	var signedInteger int8
	cout(signedInteger)
	var unsignedInteger uint16
	cout(unsignedInteger)
	var byteVar byte
	cout(byteVar)

	fmt.Println("\nBIWISE OPERATIONS:")
	a, b := 10, 3                                                        // defaults to int
	fmt.Printf("%v & %v = %v\t%b & %b = %b\n", a, b, a&b, a, b, a&b)     // AND
	fmt.Printf("%v | %v = %v\t%b | %b = %b\n", a, b, a|b, a, b, a|b)     // OR
	fmt.Printf("%v ^ %v = %v\t%b ^ %b = %b\n", a, b, a^b, a, b, a^b)     // XOR
	fmt.Printf("^%v = %v\t^%b = %b\n", b, ^b, b, ^b)                     // NOT
	fmt.Printf("%v &^ %v = %v\t%b &^ %b = %b\n", a, b, a&^b, a, b, a&^b) // AND-NOT (Bit-clear)
	fmt.Printf("%v << 3 = %v\t%b << 3 = %b\n", a, a<<3, a, a<<3)         // left-shift
	fmt.Printf("%v >> 3 = %v\t%b >> 3 = %b\n", a, a>>3, a, a>>3)         // right-shift
}

func cout(val interface{}) { fmt.Printf("%v, %T\n", val, val) }
