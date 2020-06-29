package main

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/pkg/errors"

	example "github.com/polar9527/learning-go/protobuf/proto"
)

type Demo3Service struct {
}

func (d *Demo3Service) GetUser(request example.Demo3Request, response *example.Demo3Response) error {
	// 模拟数据
	// 数据获取逻辑自行设计
	datas := map[int64]example.Demo3Response{
		1: {Name: "AAA", Age: 999, Info: &example.Demo3Response_Other{Addr: "beijing", Hobby: "sport", G: example.Demo3Response_MALE}},
		2: {Name: "BBB", Age: 888, Info: &example.Demo3Response_Other{Addr: "上海", Hobby: "sport", G: example.Demo3Response_FEMALE}},
		3: {Name: "CCC", Age: 777, Info: &example.Demo3Response_Other{Addr: "wuhan", Hobby: "sport", G: example.Demo3Response_UNKNOWN}},
		4: {Name: "DDD", Age: 666, Info: &example.Demo3Response_Other{Addr: "重庆", Hobby: "sport", G: example.Demo3Response_MALE}},
		5: {Name: "EEE", Age: 555, Info: &example.Demo3Response_Other{Addr: "", Hobby: "sport", G: example.Demo3Response_FEMALE}},
	}
	// 模拟业务处理逻辑
	if value, ok := datas[request.GetId()]; ok {
		*response = value
	} else {
		return errors.New("not found")
	}
	return nil
}
func main() {
	rpc.Register(new(Demo3Service))
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
