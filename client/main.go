package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/earthaYan/go_demo/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "amy"
)

func main() {
	// 设置到server的连接
	// 指定不同的选项，用来控制创建连接的方式
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 创建一个客户端 stub，用来执行 RPC 请求
	c := pb.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 像调用本地函数一样，调用远程的方法
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
