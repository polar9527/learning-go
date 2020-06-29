package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Arguments struct {
	A int
	B int
}

func main() {
	// DialHTTP连接到指定网络地址的HTTP RPC服务器
	// 返回一个rpc客户端
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	arg := Arguments{99, 1}
	var resp int
	// 调用指定的函数并等待其完成
	err = client.Call("DemoRpc.Add", arg, &resp)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("rpc DemoRpc Add %v\n", resp)
	err = client.Call("DemoRpc.Minus", arg, &resp)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("rpc DemoRpc Minus %v\n", resp)
	// 模拟一个错误的rpc调用
	err = client.Call("DemoRpc.Nothing", arg, &resp)
	if err != nil {
		log.Fatal(" call err:", err.Error())
	}
	fmt.Printf("rpc DemoRpc Nothing %v\n", resp)

}
