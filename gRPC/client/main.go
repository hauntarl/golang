package main

import (
	"context"
	"fmt"
	"grpcdemo/greeter"
	"io"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
)

const address = ":8888"

func main() {
	name := "World"
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		exit(os.Stderr, err.Error())
	}
	defer conn.Close()

	var (
		client      = greeter.NewGreeterClient(conn) // establish rpc connection
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	)
	defer cancel()

	res, err := client.Greet(ctx, &greeter.GreetRequest{Name: name})
	if err != nil {
		exit(os.Stderr, err.Error())
	}
	fmt.Println(res.GetMsg())
}

// terminate the program and log the custom message to given io.Writer
func exit(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
	os.Exit(1)
}
