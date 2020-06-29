package main

import (
	"fmt"
	"log"
	"net"
)

func HandleCunc(c net.Conn) {
	defer c.Close()
	Addr := c.RemoteAddr().String()
	fmt.Println(Addr, "connect success")
	buf := make([]byte, 1024)
	for {
		// 读取内容
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := buf[:n]
		// 提提示信息
		fmt.Printf("来自<%s>的数据:%s\n", Addr, string(result))
		// 关闭连接
		if "exit" == string(result) {
			fmt.Println(Addr, "退出连接")
			return
		}
		// 发送给客户端数据
		c.Write([]byte(string(result)))
	}
}
func main() {
	// 监听本机的9000端口
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Println(err)
		return
	}
	// 关闭监听
	defer listener.Close()
	// 循环等待连接
	for {
		// 等待客户端连接,返回一个连接句柄,等待下一个连接
		connect, err := listener.Accept()
		// 连接错误处理
		if err != nil {
			log.Println(err)
			continue
		}
		// 开启一协程处理客户端的连接
		go HandleCunc(connect)
	}
}
