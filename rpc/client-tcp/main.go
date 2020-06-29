package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	X int
	Y int
}

func main() {
	// 连接到指定的rpc服务器
	client, err := rpc.Dial("tcp", ":8081")
	if err != nil {
		log.Fatal(err.Error())
	}
	var result int
	p := Params{99, 1}
	err = client.Call("DemoRpc.Add", p, &result)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%d + %d = %d\n", p.X, p.Y, result)
	err = client.Call("DemoRpc.Minus", p, &result)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%d - %d = %d\n", p.X, p.Y, result)
	p.Y = 0
	err = client.Call("DemoRpc.Div", p, &result)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%d / %d = %d\n", p.X, p.Y, result)
}
