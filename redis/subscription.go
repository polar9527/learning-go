package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 定义一组常量
const (
	REDIS_IP   = "127.0.0.1"
	REDIS_PORT = "6379"
	REDIS_PWD  = ""
	REDIS_DB   = 0
)

// 定义一个redis.client类型的变量
var client *redis.Client

// 初始化函数
func init() {
	// 实例化一个redis客户端
	client = redis.NewClient(&redis.Options{
		Addr:     REDIS_IP + ":" + REDIS_PORT, // ip_port
		Password: REDIS_PWD,                   // redis连接密码
		DB:       REDIS_DB,                    // 选择的redis库
		PoolSize: 20,                          // 设置连接数,默认是10个连接
	})
}
func main() {
	defer client.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// 订阅多个频道
	sub := client.PSubscribe("news", "it", "sports", "shopping")
	_, err := sub.Receive()
	if err != nil {
		fmt.Println(err)
	}
	// 消息通道
	ch := sub.Channel()
	//  从通道中读取消息
	for message := range ch {
		fmt.Println(message.Channel, message.Payload)
	}
}
