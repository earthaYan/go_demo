package main

import (
	"context"
	"log"
	"net"

	pb "github.com/earthaYan/go_demo/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// 定义了一个 Go 结构体 server
type server struct {
	pb.UnimplementedGreeterServer
}

// 为 server 结构体添加SayHello(context.Context, pb.HelloRequest) (pb.HelloReply, error)方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "hello" + in.GetName()}, nil
}
func main() {
	// 指定监听客户端请求的端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 通过 grpc.NewServer() 创建一个 gRPC Server 实例
	s := grpc.NewServer()
	//通过pb.RegisterGreeterServer(s, &server{}) 将该服务注册到 gRPC 框架中
	pb.RegisterGreeterServer(s, &server{})
	// 通过 s.Serve(lis) 启动 gRPC 服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
