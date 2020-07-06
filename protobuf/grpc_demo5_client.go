package main

import (
	"fmt"
	pb "github.com/polar9527/learning-go/protobuf/example"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// 创建一个客户端连接
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("connect failed : ", err.Error())
	}
	// 关闭客户端连接
	defer conn.Close()
	// 客户端服务api
	client := pb.NewDemo5Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	// 客户端请求数据
	req := pb.Demo5Request{
		Name: "Tom",
		Age:  99,
		Info: &pb.Demo5Request_Other{
			Addr:  "Beijing",
			Hobby: "climbing",
			G:     0,
		},
	}
	// 调用服务接口
	r, err := client.GetInfo(ctx, &req)
	if err != nil {
		log.Fatal("called failed : ", err.Error())
	}
	// 打印结果
	fmt.Println(r.Info)
}
