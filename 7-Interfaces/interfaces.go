package main

import (
	"bytes"
	"fmt"
	"io"
)

// Writer is exported
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter is exported
type ConsoleWriter struct{}

// Write is exported
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Print(string(data))
	if err != nil {
		return 0, err
	}
	fmt.Println("\tBytes written to console: ", n)
	return n, err
}

type incrementer interface {
	increment() int
}

type intCounter int

func (ic *intCounter) increment() int {
	*ic++
	return int(*ic)
}

func main() {
	/*
		Interfaces:
			- structs are used as data containers, but interfaces define behavior
			- in go, we do not explicity tell go that we are implementing an interface,
			  we use interfaces implicity
			- interfaces in go are used to achieve polymorphism like behavior
			- you can compose interfaces together similar to structs
	*/
	var myConsoleWriter Writer = ConsoleWriter{}
	myConsoleWriter.Write([]byte("Hello World"))

	fmt.Println(*new(intCounter))
	var myIncrementer incrementer = new(intCounter)
	for i := 0; i < 5; i++ {
		fmt.Print(myIncrementer.increment(), " ")
	}
	fmt.Println()

	// this is a comprehensive example for composition of interfaces
	fmt.Println("\nCreating a byte buffer for some string:")
	var myWriterReader WriterReader = NewBufferedWriterReader()
	totalBytes, err := myWriterReader.Write(
		[]byte("This repository consists of code snippets for golang basics"),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Total bytes buffered -", totalBytes)

	fmt.Println("\nRecursively Reading exactly 16 bytes from the buffer:")
	err = myWriterReader.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("\nFlushing out all remaining bytes from the buffer:")
	err = myWriterReader.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//type conversion using interfaces
	fmt.Println("\nType conversion:")
	bwr, ok := myWriterReader.(*BufferedWriterReader)
	if ok {
		fmt.Println(bwr)
	} else {
		fmt.Println("Conversion to BufferedWriterReader Failed")
	}
	r, ok := myWriterReader.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion to io.Reader Failed", r)
	}

	//empty interface
	fmt.Println("\nEmpty Interfaces for type casting:")
	var myEmptyInterface interface{} = NewBufferedWriterReader()
	if wr, ok := myEmptyInterface.(WriterReader); ok {
		wr.Write(
			[]byte("This repository consists of code snippets for golang basics"),
		)
		wr.Close()
	}
	fmt.Println()
	r, ok = myEmptyInterface.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Coversion to io.Reader Failed", r)
	}
	// empty interfaces are most commonly used with type switches to provide different
	// implementations based on the type

	/*
		Best Practices:
			- prefer many small interfaces over large monolithic ones, if you need
			  large monolithic ones, go ahead and compose small interfaces
			- io.Writer, io.Reader, interface{} are some of the most powerful interfaces
			  and they have 1 or 0 methods
			- if you're creating a library for others to consume, publish the concrete type,
			  don't create interface assuming you know how people are going to use it, allow
			  them to create interfaces that your type will implement, that way they don't have
			  to implement a whole bunch of methods that they'll never even use
			- go has totally inverted approach towards interfaces comapred to other languages
			- do export interfaces for the types that you'll be consuming, so when you're defining
			  a type that you'll be consuming export interfaces for others to implement then you
			  only have to worry about the behavior they are exposing to you
			- if possible define your methods and functions to receive interfaces
	*/
}

// NewBufferedWriterReader is exported
func NewBufferedWriterReader() *BufferedWriterReader {
	return &BufferedWriterReader{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// Reader is exported
type Reader interface {
	Read() error
}

// Closer is exported
type Closer interface {
	Close() error
}

// WriterReader is exported
type WriterReader interface {
	Writer
	Reader
	Closer
}

// BufferedWriterReader is exported
type BufferedWriterReader struct {
	buffer *bytes.Buffer
}

// Write is exported
func (bwc *BufferedWriterReader) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// Read is exported
func (bwc *BufferedWriterReader) Read() error {
	byteData := make([]byte, 16)
	for bwc.buffer.Len() > 16 {
		_, err := bwc.buffer.Read(byteData)
		if err != nil {
			return err
		}

		var myWriter Writer = new(ConsoleWriter)
		_, err = myWriter.Write(byteData)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close is exported
func (bwc *BufferedWriterReader) Close() error {
	for bwc.buffer.Len() > 0 {
		byteData := bwc.buffer.Next(16)

		var myWriter Writer = new(ConsoleWriter)
		_, err := myWriter.Write(byteData)
		if err != nil {
			return err
		}
	}
	return nil
}
