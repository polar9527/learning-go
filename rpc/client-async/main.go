package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type ArgsDemo struct {
	A int
	B int
}

func main() {
	//DialHTTP连接到指定网络地址的HTTP RPC服务器
	//返回一个rpc客户端
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	arg := ArgsDemo{9999, 8888}
	var resp int
	//异步调用指定的函数并等待其完成
	call := client.Go("DemoRpc3.Add", arg, &resp, nil)
	// 正常的同步调用
	err = client.Call("DemoRpc3.Minus", arg, &resp)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("rpc DemoRpc Minus %v\n", resp)
	for {
		select {
		case <-call.Done:
			if call.Error != nil {
				log.Println(call.Error.Error())
				return
			}
			fmt.Printf("rpc DemoRpc Add %v\n", resp)
			return
		default:
			fmt.Println("wait...")
			time.Sleep(1 * time.Second)
		}
	}

}
