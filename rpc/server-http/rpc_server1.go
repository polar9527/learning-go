package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Arguments struct {
	A int
	B int
}
type DemoRpc struct{}

func (d *DemoRpc) Add(req Arguments, resp *int) error {
	*resp = req.A + req.B
	return nil
}
func (d *DemoRpc) Minus(req Arguments, resp *int) error {
	*resp = req.A - req.B
	return nil
}

func main() {
	// 注册rpc服务
	rpc.Register(new(DemoRpc))
	// 将用于RPC消息的HTTP处理程序注册到DefaultServer
	rpc.HandleHTTP()
	// 监听8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
