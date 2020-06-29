package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"

	example "github.com/polar9527/learning-go/protobuf/proto"
)

func main() {
	// 相当于示例化一个Person的对象
	// 这个Person的结构体就是我们通过代码自动生成的
	d1 := example.Person{
		Name: "tom",
		Age:  99,
		Info: &example.Person_Other{
			Addr:  "beijing",
			Hobby: "code",
			G:     example.Person_MALE,
		},
	}
	// 进行序列化操作
	d1Encode, err := proto.Marshal(&d1)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(d1Encode)
	// 进行反序列化操作
	d1Decode := example.Person{}
	err = proto.Unmarshal(d1Encode, &d1Decode)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(d1Decode.GetName())
	fmt.Println(d1Decode.GetAge())
	fmt.Println(d1Decode.GetInfo().GetG())
}
