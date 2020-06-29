package main

import (
	"fmt"
	"log"
	"net/rpc"

	example "github.com/polar9527/learning-go/protobuf/proto"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	request := example.Demo3Request{Id: 1}
	var response example.Demo3Response
	err = client.Call("Demo3Service.GetUser", request, &response)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(response.GetName())
	fmt.Println(response.GetInfo().GetAddr())
}
