package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"time"
)

type ArgsDemo struct {
	A int
	B int
}
type DemoRpc3 struct{}

func (d *DemoRpc3) Add(req ArgsDemo, resp *int) error {
	for i := 0; i < 5; i++ {
		fmt.Println("sleep...", i)
		time.Sleep(1 * time.Second)
	}
	*resp = req.A + req.B
	fmt.Println("Add Do")
	return nil
}
func (d *DemoRpc3) Minus(req ArgsDemo, resp *int) error {
	*resp = req.A - req.B
	return nil
}

func main() {
	// 注册rpc服务
	rpc.Register(new(DemoRpc3))
	// 将用于RPC消息的HTTP处理程序注册到DefaultServer
	rpc.HandleHTTP()
	// 监听8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
