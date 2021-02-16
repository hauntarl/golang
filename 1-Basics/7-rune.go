package main

import "fmt"

/*
	Runes
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

	For more information, refer the following articles listed in readme:
		Runes and character encoding
		UTF-8: Bits, Bytes, and Benefits - Russ Coxx

	Refer Go language specification for more info: https://golang.org/ref/spec
*/
func main() {
	j := 'a' // rune datatype
	fmt.Printf("Rune '%v' - %v, %T\n", string(j), j, j)
}
