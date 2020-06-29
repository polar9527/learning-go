package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
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

// 生成指定长度的随机字符串
func randString(n int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = s[rand.Intn(len(s))]
	}
	return string(b)
}
func main() {
	defer client.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var data string
	var channels []string = []string{"news", "it", "sports", "shopping"}
	for {
		fmt.Printf("please input some data:")
		fmt.Scanln(&data)
		// 退出
		if data == "quit" {
			break
		}
		channel := channels[rand.Intn(4)]
		// 向一个频道发布消息
		result := client.Publish(channel, data+randString(10)).Val()
		if result == 1 {
			fmt.Println("send info to", channel, "success")
		}
	}
}
