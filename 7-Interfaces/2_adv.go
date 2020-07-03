package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// this is a comprehensive example for composition of interfaces
	fmt.Println("\nCreating a byte buffer for some string:")
	var wr WriteReadCloser = NewBufferedWriterReader()
	totalBytes, err := wr.Write(
		[]byte("This repository consists of code snippets for golang basics"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Total bytes buffered -", totalBytes)

	fmt.Println("\nRecursively Reading exactly 16 bytes from the buffer:")
	err = wr.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("\nFlushing out all remaining bytes from the buffer:")
	err = wr.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//type conversion using interfaces
	fmt.Println("\nType conversion:")
	bwr, ok := wr.(*BufferedWriterReader)
	if ok {
		fmt.Println(bwr)
	} else {
		fmt.Println("Conversion to BufferedWriterReader Failed")
	}

	r, ok := wr.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion to io.Reader Failed", r)
	}

	//empty interface
	fmt.Println("\nEmpty Interfaces for type casting:")
	var myEmptyInterface interface{} = NewBufferedWriterReader()
	if wr, ok := myEmptyInterface.(WriteReadCloser); ok {
		wr.Write(
			[]byte("This repository consists of code snippets for golang basics"),
		)
		wr.Close()
	}
	fmt.Println()
	if r, ok = myEmptyInterface.(io.Reader); ok {
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
			- go has totally inverted approach towards interfaces compared to other languages
			- do export interfaces for the types that you'll be consuming, so when you're defining
			  a type that you'll be consuming, export interfaces for others to implement then you
			  only have to worry about the behavior they are exposing to you
			- if possible define your methods and functions to receive interfaces
	*/
}

// Writer defines the structure of methods it exports
type Writer interface {
	Write([]byte) (int, error)
}

// Reader provides the definition of Read method
type Reader interface {
	Read() error
}

// Closer provides the definiton of Close method
type Closer interface {
	Close() error
}

// WriteReadCloser is a composite interface comprised of Writer, Reader and Closer
type WriteReadCloser interface {
	Writer
	Reader
	Closer
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

// BufferedWriterReader is a user defined struct which holds
// pointer to bytes.Buffer object
type BufferedWriterReader struct {
	buffer *bytes.Buffer
}

// NewBufferedWriterReader is like a constructor of BufferedWriterReader
// to explicitly initialize the BufferedWriterReader object
func NewBufferedWriterReader() *BufferedWriterReader {
	return &BufferedWriterReader{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// Write method of BufferedWriterReader stores data into the buffer
func (bwc *BufferedWriterReader) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// Read method reads data 16 bytes at a time, if there are <16 bytes
// present in the buffer or it reaches the EOF, it stops reading
func (bwc *BufferedWriterReader) Read() error {
	byteData := make([]byte, 16)
	for bwc.buffer.Len() > 16 {
		_, err := bwc.buffer.Read(byteData)
		if err != nil {
			return err
		}

		var cw Writer = new(ConsoleWriter)
		_, err = cw.Write(byteData)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close method flushes the remaining bytes from the buffer, applicable
// if the reamining bytes are less than 16 in this case
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
