package main

import "fmt"

/*
	Strings: UTF8
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
		Refer: https://unicodebook.readthedocs.io/unicode_encodings.html

	There are four ways to represent the integer value as a numeric constant:
		- \x followed by exactly two hexadecimal digits
		- \u followed by exactly four hexadecimal digits
		- \U followed by exactly eight hexadecimal digits (limit < \U0010FFFF)
		- \ followed by exactly three octal digits. (0-255)
	Although these representations all result in an integer, they have different
	valid ranges. Further they also allow special escape characters
*/
func main() {
	fmt.Println("\nSTRINGS:")
	h := "hello world"
	cout(h)
	fmt.Printf("%v, %T\tbyte at index: 6\n", h[6], h[6]) // cannot modify the value
	fmt.Printf("%v, %T\tbyte converted to string\n", string(h[6]), string(h[6]))
	i := []byte(h) // byte slicing
	fmt.Printf("byte array: %v, %T\n", i, i)
}

func cout(val interface{}) { fmt.Printf("%v, %T\n", val, val) }
