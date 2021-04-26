package main

import (
	"context"
	"fmt"
	"grpcdemo/greeter"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
)

const (
	network = "tcp"
	address = ":8888"
)

func main() {
	var (
		handler greeterServer // implements GreeterServer
		server  = grpc.NewServer()
	)
	// register a new server along with object to handler incoming requests
	greeter.RegisterGreeterServer(server, handler)
	// create a listener which accepts incoming requests on given port
	listener, err := net.Listen(network, address)
	if err != nil {
		exit(os.Stderr, err.Error())
	}
	defer listener.Close()
	// start the server at given listener details
	exit(os.Stderr, server.Serve(listener).Error())
}

// terminate the program and log the custom message to given io.Writer
func exit(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
	os.Exit(1)
}

type greeterServer struct {
	greeter.UnimplementedGreeterServer
}

func (s greeterServer) Greet(ctx context.Context,
	req *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	return &greeter.GreetResponse{Msg: fmt.Sprintf("Hello %s", req.Name)}, nil
}
