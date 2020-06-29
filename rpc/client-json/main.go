package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type JsonParams struct {
	X int
	Y int
}

func main() {
	// 连接到指定的json rpc服务器
	client, err := jsonrpc.Dial("tcp", ":8081")
	if err != nil {
		log.Fatal(err.Error())
	}
	var result int
	p := JsonParams{60, 40}
	err = client.Call("JsonDemo.Add", p, &result)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%d + %d = %d\n", p.X, p.Y, result)
	err = client.Call("JsonDemo.Minus", p, &result)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%d - %d = %d\n", p.X, p.Y, result)
	p.Y = 0
	err = client.Call("JsonDemo.Div", p, &result)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%d / %d = %d\n", p.X, p.Y, result)
}
