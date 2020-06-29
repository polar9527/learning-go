package main

import (
	"fmt"
	"math/rand"
	"strconv"

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
func randString(n int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
	key := "zdemo"
	key1 := "zdemo1"
	client.Del(key)
	client.Del(key1)
	for i := 0; i < 10; i++ {
		score := float64(rand.Intn(100) - 50)
		member := "golang-" + strconv.Itoa(i)
		data := &redis.Z{
			score,
			member,
		}
		// 向有序集合中添加成员
		client.ZAdd(key, data).Val()
	}
	for i := 0; i < 5; i++ {
		score := float64(rand.Intn(100) - 50)
		member := "golang-" + strconv.Itoa(i)
		data := &redis.Z{
			score,
			member,
		}
		// 向有序集合中添加成员
		client.ZAdd(key1, data).Val()
	}
	// 计算成员个数
	n1 := client.ZCard(key).Val()
	fmt.Println(n1) //10
	// 获取成员分数
	s1 := client.ZScore(key, "golang-6").Val()
	fmt.Println(s1) // -25
	// 修改成员分数
	v1 := client.ZIncrBy(key, 60.00, "golang-6").Val()
	fmt.Println(v1)
	// 从低到高返回排名
	s2 := client.ZRank(key, "golang-6").Val()
	fmt.Println(s2) //8
	// 从高到低返回排名
	s3 := client.ZRevRank(key, "golang-6").Val()
	fmt.Println(s3) //1
	// 获取指定范围的成员排名,从低到高排名
	s4 := client.ZRange(key, 0, n1-5).Val()
	fmt.Println(s4) //[golang-9 golang-5 golang-7 golang-2 golang-8 golang-3]
	// 获取指定范围的成员排名,从高到低排名
	s5 := client.ZRevRange(key, 0, n1-5).Val()
	fmt.Println(s5) //[golang-1 golang-6 golang-4 golang-0 golang-3 golang-8]
	// 删除成员
	v2 := client.ZRem(key, "golang-6").Val()
	fmt.Println(v2) //1
	// 计算两个有序集合的交集
	key2 := "zdemo2"
	// kslice := []string{key, key1}
	wslice := []float64{1.00, 1.00}
	z := redis.ZStore{
		wslice,
		"SUM",
	}
	r1 := client.ZInterStore(key2, z, key, key1).Val()
	fmt.Println(r1) //5
	// 计算两个有序集合的并集
	key3 := "zdemo3"
	r2 := client.ZUnionStore(key3, z).Val()
	fmt.Println(r2) // 9

}
