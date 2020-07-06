package main

import (
	"fmt"
	pb "github.com/polar9527/learning-go/grpc/example"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 实现服务
type Demo5Server struct {
	pb.UnimplementedDemo5Server
}

// 实现服务定义的接口
func (d *Demo5Server) GetInfo(ctx context.Context, in *pb.Demo5Request) (*pb.Demo5Response, error) {
	InfoS := fmt.Sprintf("my name is %s, i am %d, i live in %s", in.Name, in.Age, in.Info.Addr)
	return &pb.Demo5Response{Info: InfoS}, nil
}
func (d *Demo5Server) SetName(ctx context.Context, in *pb.Demo5Request) (*pb.Demo5Response, error) {
	return &pb.Demo5Response{Info: "getName"}, nil
}
func main() {
	// 监听8080端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen : ", err.Error())
	}
	// 创建一个没有注册服务的新的gRpc服务器
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterDemo5Server(s, &Demo5Server{})
	fmt.Println("gRpc 服务端开启")
	// 接收gRpc请求
	if err := s.Serve(listen); err != nil {
		log.Fatal("this is error : ", err.Error())
	}
}
