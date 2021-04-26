# Introduction to Protocol Buffers

This tutorial provides a basic Go programmer's introduction to working with **[protocol buffers](https://developers.google.com/protocol-buffers)**, using the **[proto3](https://developers.google.com/protocol-buffers/docs/proto3)** version of the protocol buffers language. By walking through creating a simple example application, it shows you how to:

- Define message formats in a .proto file.
- Use the protocol buffer compiler.
- Use the Go protocol buffer API to write and read messages.

Custom implementation of **[Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)** tutorial

## More

- [Language Guide](https://developers.google.com/protocol-buffers/docs/proto3) (proto3)
- Go [Generated Code](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
- Go api doc [google.golang.org/protobuf/proto](https://pkg.go.dev/google.golang.org/protobuf/proto)
- Go [FAQ](https://developers.google.com/protocol-buffers/docs/reference/go/faq)

## Description

The example we're going to use is a very simple "address book" application that can read and write people's contact details to and from a file. Each person in the address book has a name, an ID, an email address, and a contact phone number.

How do you serialize and retrieve structured data like this? There are a few ways to solve this problem:

- Use **[gobs](https://golang.org/pkg/encoding/gob/)** to serialize Go data structures. This is a good solution in a Go-specific environment, but it doesn't work well if you need to share data with applications written for other platforms.
- You can invent an ad-hoc way to encode the data items into a single string – such as encoding 4 ints as "12:3:-23:67". This is a simple and flexible approach, although it does require writing one-off encoding and parsing code, and the parsing imposes a small run-time cost. This works best for encoding very simple data.
- Serialize the data to XML. This approach can be very attractive since XML is (sort of) human readable and there are binding libraries for lots of languages. This can be a good choice if you want to share data with other applications/projects. However, XML is notoriously space intensive, and encoding/decoding it can impose a huge performance penalty on applications. Also, navigating an XML DOM tree is considerably more complicated than navigating simple fields in a class normally would be.
Protocol buffers are the flexible, efficient, automated solution to solve exactly this problem. With protocol buffers, you write a .proto description of the data structure you wish to store. From that, the protocol buffer compiler creates a class that implements automatic encoding and parsing of the protocol buffer data with an efficient binary format. The generated class provides getters and setters for the fields that make up a protocol buffer and takes care of the details of reading and writing the protocol buffer as a unit. Importantly, the protocol buffer format supports the idea of extending the format over time in such a way that the code can still read data encoded with the old format.

Protocol buffers are the flexible, efficient, automated solution to solve exactly this problem. With protocol buffers, you write a `.proto` description of the data structure you wish to store. From that, the protocol buffer compiler creates a class that implements automatic encoding and parsing of the protocol buffer data with an efficient binary format. The generated class provides getters and setters for the fields that make up a protocol buffer and takes care of the details of reading and writing the protocol buffer as a unit. Importantly, the protocol buffer format supports the idea of extending the format over time in such a way that the code can still read data encoded with the old format.

### Project Structure

- *schema* contains .proto file which defines our messages
- *adb* contains the generated .go file using defined schema
- *cmd* contains main.go file which utilizes file from *adb*

### Run Commands

- [Compiling your protocol buffers](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)
- go build -o addressbook.exe cmd/main.go
- addressbook list (to display all entries made into address book)
- addressbook add `[name of person]` (to create a new person entry in address book)

### Output

``` terminal
D:\godemo\protobuf>protoc -I=schema --go_out=. schema/addressbook.proto

D:\godemo\protobuf>go build -o addressbook.exe cmd/main.go

D:\godemo\protobuf>addressbook
missing subcommand: list or add

D:\godemo\protobuf>addressbook unknown
unregistered subcommand: unknown

D:\godemo\protobuf>addressbook list
open address.pb: The system cannot find the file specified.

D:\godemo\protobuf>addressbook add Sameer Mungole
Enter email address (blank for none): sameer.mungole@gmail.com
Enter a phone number (or leave blank to finish): 555-666-7777
Is this a mobile, home, or work phone? home
Enter a phone number (or leave blank to finish): 111-222-3333
Is this a mobile, home, or work phone? work
Enter a phone number (or leave blank to finish):

D:\godemo\protobuf>addressbook list
Person ID: 0
  Name: Sameer Mungole
  E-mail address: sameer.mungole@gmail.com
  Home phone #: 555-666-7777
  Work phone #: 111-222-3333

D:\godemo\protobuf>addressbook add hauntarl
Enter email address (blank for none):
Enter a phone number (or leave blank to finish): 9876543210
Is this a mobile, home, or work phone? default
Enter a phone number (or leave blank to finish):

D:\godemo\protobuf>addressbook list
Person ID: 0
  Name: Sameer Mungole
  E-mail address: sameer.mungole@gmail.com
  Home phone #: 555-666-7777
  Work phone #: 111-222-3333
Person ID: 1
  Name: hauntarl
  Mobile phone #: 9876543210
```
