package main

import (
	"context"
	"log"
	"net"

	pb "github.com/earthaYan/go_demo/print"
	"google.golang.org/grpc"
)

const (
	port = ":5001"
)

type print struct {
	pb.UnimplementedPrintServer
}

func (s *print) PrintHello(ctx context.Context, r *pb.PrintRequest) (*pb.PrintResponse, error) {
	log.Printf("Received: %v", r.GetName())
	return &pb.PrintResponse{Message: "Hello " + r.GetName()}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPrintServer(s, &print{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
