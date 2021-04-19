package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"protobuf/adb"
	"strings"

	"google.golang.org/protobuf/proto"
)

const (
	get   = "list"
	put   = "add"
	fname = "address.pb"
)

func init() {
	flag.Parse()
	if flag.NArg() < 1 {
		exit(os.Stderr, fmt.Sprintf("missing subcommand: %s or %s", get, put))
	}
}

func main() {
	// get action associated with given flag
	do, ok := subcmd[flag.Arg(0)]
	if !ok {
		exit(os.Stderr, fmt.Sprintf("unregistered subcommand: %s", flag.Arg(0)))
	}
	// call the action associated with given flag
	if err := do(); err != nil {
		exit(os.Stderr, err.Error())
	}
}

// terminate the program and log the custom message to given io.Writer
func exit(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
	os.Exit(1)
}

type action func() error

// flag to action mapping
var subcmd = map[string]action{
	// reads the database and displays the content to provided io.Writer
	get: func() error {
		book, err := readDB()
		if err != nil {
			return err
		}

		for _, person := range book.People {
			display(os.Stdout, person)
		}
		return nil
	},
	// reads the database, creates a new person, writes it to database
	put: func() error {
		book, _ := readDB()
		if book == nil {
			book = &adb.AddressBook{}
		}

		person, err := getDetails(os.Stdin, book)
		if err != nil {
			return err
		}

		book.People = append(book.People, person)
		return writeDB(book)
	},
}

// reads content from fname, unmarshals the data into Go structure, returns it
func readDB() (*adb.AddressBook, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	book := &adb.AddressBook{}
	if err := proto.Unmarshal(data, book); err != nil {
		return nil, err
	}
	return book, nil
}

// writes the attributes of given person to given io.Writer
func display(w io.Writer, p *adb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "  Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "  E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case adb.Person_MOBILE:
			fmt.Fprint(w, "  Mobile phone #: ")
		case adb.Person_HOME:
			fmt.Fprint(w, "  Home phone #: ")
		case adb.Person_WORK:
			fmt.Fprint(w, "  Work phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}

// reads input from given io.Reader, creates a new Person based on given info
func getDetails(r io.Reader, book *adb.AddressBook) (*adb.Person, error) {
	p := &adb.Person{
		Id:   int32(len(book.People)),
		Name: strings.Join(flag.Args()[1:], " "), // fetch name from cmd flags
	}
	rd := bufio.NewReader(r) // creates a new buffered reader

	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return nil, err
	}
	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		pNum, err := rd.ReadString('\n')
		if err != nil {
			return nil, err
		}
		pNum = strings.TrimSpace(pNum)
		if pNum == "" {
			break
		}
		number := &adb.Person_PhoneNumber{Number: pNum}

		fmt.Print("Is this a mobile, home, or work phone? ")
		pType, err := rd.ReadString('\n')
		if err != nil {
			return nil, err
		}
		switch strings.TrimSpace(pType) {
		case "mobile":
			number.Type = adb.Person_MOBILE
		case "home":
			number.Type = adb.Person_HOME
		case "work":
			number.Type = adb.Person_WORK
		}

		p.Phones = append(p.Phones, number)
	}
	return p, nil
}

// marshals Go structure into bytes and writes them to database
func writeDB(book *adb.AddressBook) error {
	data, err := proto.Marshal(book)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fname, data, os.ModePerm)
}
