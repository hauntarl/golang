# Introduction to gRPC

In **[gRPC](https://www.grpc.io/)**, a client application can directly call a method on a server application on a different machine as if it were a local object, making it easier for you to create distributed applications and services. As in many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. On the server side, the server implements this interface and runs a gRPC server to handle client calls. On the client side, the client has a stub (referred to as just a client in some languages) that provides the same methods as the server.

![High Level Architecture](https://www.grpc.io/img/landing-2.svg)

gRPC clients and servers can run and talk to each other in a variety of environments - from servers inside Google to your own desktop - and can be written in any of gRPC’s supported languages. So, for example, you can easily create a gRPC server in Java with clients in Go, Python, or Ruby. In addition, the latest Google APIs will have gRPC versions of their interfaces, letting you easily build Google functionality into your applications.

## Working with Protocol Buffers

By default, gRPC uses **[Protocol Buffers](https://developers.google.com/protocol-buffers/docs/overview)**, Google’s mature open source mechanism for serializing structured data (although it can be used with other data formats such as JSON). Not familiar with protocol buffers? Refer *[protobuf](https://github.com/hauntarl/golang/tree/master/protobuf)* to get started.

gRPC uses `protoc` with a special gRPC plugin to generate code from your proto file: you get generated gRPC client and server code, as well as the regular protocol buffer code for populating, serializing, and retrieving your message types.

## Description

>Custom implementation of **[Quick start](https://www.grpc.io/docs/languages/go/quickstart/)** guide for **gRPC in Go**.

A simple client-server application with gRPC, which takes name of client as request and returns appropriate greetings as response.

Define gRPC services in ordinary proto files, with RPC method parameters and return types specified as protocol buffer messages:

``` proto
message GreetRequest {
    // The request message containing the user's name.
    string name = 1;
}

message GreetResponse {
    // The response message containing the greetings
    string msg = 1;
}

// The greeting service definition.
service Greeter {
    // Sends a greeting
    rpc Greet (GreetRequest) returns (GreetResponse);
}
```

### Project Structure

- *[greeter](https://github.com/hauntarl/golang/tree/master/gRPC/greeter)* - contains `.proto` and generated files
- *[client](https://github.com/hauntarl/golang/tree/master/gRPC/client)* - contains implementation of client
- *[server](https://github.com/hauntarl/golang/tree/master/gRPC/server)* - contains implementation of server

### Run Commands

- go run server/main.go
- go build -o greet.exe client/main.go
- greet `<name>[optional]...`

### Output

``` terminal
D:\golang\gRPC>go run server/main.go

D:\golang\gRPC>go build -o greet.exe client/main.go

D:\golang\gRPC>greet
Hello World

D:\golang\gRPC>greet hauntarl
Hello hauntarl

D:\golang\gRPC>greet Sameer Mungole
Hello Sameer Mungole
```

## What's Next?

- [Basics tutorial](https://grpc.io/docs/languages/go/basics/) in Go
- Go [Generated Code](https://www.grpc.io/docs/languages/go/generated-code/)
- Go api doc [google.golang.org/grpc](https://pkg.go.dev/google.golang.org/grpc)
- gRPC [Documentation](https://www.grpc.io/docs/)
- [Core concepts, architecture and lifecycle](https://www.grpc.io/docs/what-is-grpc/core-concepts/)
- gRPC [Guides](https://grpc.io/docs/guides/)
- gRPC [FAQs](https://www.grpc.io/docs/what-is-grpc/faq/)
