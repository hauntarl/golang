package main

import (
	"fmt"
	"net/http"
)

/*
	Panic:
		- when something unplanned happens in our go application
		- their are two ways to go:
			1. return an error value (normal case)
			2. panic
		- panics are used when your application gets into a state that it cannot recover from
		  eg. divide by zero, cannot obtain TCP port for web server
		- funtion will stop executing, but deferred function will still fire
		- if nothing handles that panic then program will exit
*/
func main() {
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("Start")
	// defer fmt.Println("I am always executed, before panic stops the function execution")
	// panic("Something bad happened")

	// small note about panic. Go doesn't panic, we do!
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})

	err := http.ListenAndServe(":8080", nil) // run this application twice on same port to get error
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error()) // then you panic
	}
}
