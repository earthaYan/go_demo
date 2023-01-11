package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/earthaYan/go_demo/print"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:5001"
	defaultName = "world"
)

func main() {
	conn, error := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if error != nil {
		log.Fatalf("did not connect: %v", error)
	}
	defer conn.Close()
	c := pb.NewPrintClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PrintHello(ctx, &pb.PrintRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
