package main

import "fmt"

/*
	Interfaces:
		- structs are used as data containers, but interfaces define behavior
		- in go, we do not explicity tell go that we are implementing an
		  interface, we use interfaces implicity
		- interfaces in go are used to achieve polymorphism like behavior
		- you can compose interfaces together similar to structs
		eg. interface {
				Read([]byte) (int, error)
				Write([]byte) (int, error)
				Close() error
			}

		Invalid:
		interface {
			String() string
			String() string  // illegal: String not unique
			_(x int)         // illegal: method must have non-blank name
		}

		Illegal: Bad cannot embed itself
		type Bad interface {
			Bad
		}

		Illegal: Bad1 cannot embed itself using Bad2
		type Bad1 interface {
			Bad2
		}
		type Bad2 interface {
			Bad1
		}
*/

// Writer defines the structure of methods it exports
type Writer interface{ Write([]byte) (int, error) }

// ConsoleWriter is exported
type ConsoleWriter struct{}

// Write is associated with the context of ConsoleWriter which implicity
// implements the Writer interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Print(string(data))
	if err != nil {
		return 0, err
	}
	fmt.Println("\tBytes written to console: ", n)
	return n, err
}

func main() {
	var cw Writer = ConsoleWriter{}
	cw.Write([]byte("Hello World"))
}
