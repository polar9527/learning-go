package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
	"net/rpc"
)

type Demo struct{}
type Params struct {
	X int
	Y int
}

// 暴露对外的服务
func (d *Demo) Add(p Params, result *int) error {
	*result = p.X + p.Y
	return nil
}
func (d *Demo) Minus(p Params, result *int) error {
	*result = p.X - p.Y
	return nil
}
func (d *Demo) Div(p Params, result *int) error {
	if p.Y == 0 {
		return errors.New("dividend is zero")
	}
	*result = p.X / p.Y
	return nil
}
func main() {
	//注册一个自定义名称的rpc服务
	//和rpc.Register作用是一样
	rpc.RegisterName("DemoRpc", new(Demo))
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
		}
		go rpc.ServeConn(conn)
	}

}
