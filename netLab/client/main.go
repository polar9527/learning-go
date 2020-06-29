package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:9000"
	// 连接服务器
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer conn.Close()
	// 设置缓冲
	buf := make([]byte, 2048)
	for {
		// 接收输入内容
		fmt.Println("please input :")
		fmt.Scan(&buf)
		// 写入内容
		conn.Write(buf)
		// 获取响应
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := buf[:n]
		// 输入响应
		fmt.Println(string(result))
	}

}
