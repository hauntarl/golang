package main

import (
	"fmt"
	"reflect"
)

/*
	Tags:
		- when defining a structure we can attach a tag or a simple message to each field
		- tag can be any arbitrary messsage and go does not use those tags in any way.
		- to get the tag information for any particular field we need to import "reflect"
		- it can look in the structure definition and fetch tags for respective fields
		- in order to make use of tags, you need to create your own validation library
		  which will have its own logic to look for keywords in tags
*/
func main() {
	fmt.Println("\nTAGS:")
	type simple struct {
		name    string `required:"true" max:"100"`
		address string // in the above tag, linter will give a warning if
		theDOB  string // space separated key:"value" is not present
	}
	t := reflect.TypeOf(simple{})
	field, ok := t.FieldByName("name")
	fmt.Println("Field Information:", field)
	fmt.Printf("tag: %v, isPresent: %v", field.Tag, ok)
}
