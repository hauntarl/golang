package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	Defer:
		- used to delay execution of a statement until function exits
		- deferred functions are executed in LIFO
		- used to close out resources, so it makes sense for LIFO execution
		  because resources may be dependent on each other

	Pattern:
		- open the resource
		- check for error
		- close the resource using defer
		- perform operations on resource

	NOTE:
		- if you are using loop for fetching resources, you should not use defer
		  as defer statements runs at the end of the function, so we'll be opening
		  all the resources without closing any, which may cause some memory issue
		- or you can delegate the processing of resources to another function and
		  make that function close the resource
*/
func main() {
	funcState := "\n\tEnd of function"  // defer saves the argument values
	defer fmt.Println(funcState)        // from where it was triggered,
	funcState = "\n\tStart of function" // changes made to variables after defer statement
	fmt.Println(funcState)              // won't affect the value passed to defer function

	slice := []string{"Data old"}    // unless arguments are reference to a memory location
	defer fmt.Println("\n\t", slice) // changes made to variables after defer statement
	slice[0] = "Data updated"        // will affect the end output
	fmt.Println("\n\t", slice)

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fmt.Println("\n\tres.Body closed!")
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s", robots)
}
