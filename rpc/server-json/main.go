package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type JsonDemo struct{}
type JsonParams struct {
	X int
	Y int
}

// 暴露对外的服务
func (d *JsonDemo) Add(p JsonParams, result *int) error {
	*result = p.X + p.Y
	return nil
}
func (d *JsonDemo) Minus(p JsonParams, result *int) error {
	*result = p.X - p.Y
	return nil
}
func (d *JsonDemo) Div(p JsonParams, result *int) error {
	if p.Y == 0 {
		return errors.New("dividend is zero")
	}
	*result = p.X / p.Y
	return nil
}
func main() {
	//注册一个自定义名称的rpc服务
	rpc.RegisterName("JsonDemo", new(JsonDemo))
	// 开启一个tcp服务,监听8081端口
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		// 等待连接
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println(conn.RemoteAddr().String())
		}
		//在单个连接上运行JSON-RPC服务器
		go jsonrpc.ServeConn(conn)
	}
}
