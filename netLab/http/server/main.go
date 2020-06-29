package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// func MyFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.Method, r.RemoteAddr)
// 	fmt.Fprintf(w, "%s", "https request")
// }
// func main() {
// 	http.HandleFunc("/", MyFunc)
// 	// $ openssl genrsa -out server.key 2048
// 	// $ openssl req -new -x509 -key server.key -out server.crt -days 365
// 	cf := "E:/Go/src/GoNote/chapter9/demo12/main/server.crt"
// 	ck := "E:/Go/src/GoNote/chapter9/demo12/main/server.key"
// 	// 监听8081端口
// 	http.ListenAndServeTLS(":8081", cf, ck, nil)
// }

type responseToClient struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	var dataForm map[string][]string
	// 标识一个客户端的连接
	fmt.Println("client connect success ", r.RemoteAddr)
	// 获取地址栏内容
	fmt.Println(r.Method, r.RequestURI)
	// 获取请求头内容
	for k, v := range r.Header {
		fmt.Println(k, v[0])
	}
	data := make(map[string]string)
	if err := r.ParseForm(); err == nil {
		if r.Form != nil {
			dataForm = r.Form
		}
	}
	// 读取客户端的内容
	buf := make([]byte, 2048)
	n, _ := r.Body.Read(buf)
	// 获取请求体中的内容
	fmt.Println("receive data from body", string(buf[:n]))

	if r.Method == "GET" {
		// 解析参数
		r.ParseForm()
		for k, v := range r.Form {
			data[k] = v[0]
		}
	}
	// 处理客户端发送的POST请求和PUT请求
	if r.Method == "POST" || r.Method == "PUT" {
		ct, ok := r.Header["Content-Type"]
		if ok {
			// 如果是json数据根据请求头判断
			if ct[0] == "application/json" {
				json.Unmarshal(buf[:n], &data)
			}
			// 如果是POST表单数据
			if ct[0] == "application/x-www-form-urlencoded" {
				if dataForm != nil {
					for k, v := range dataForm {
						data[k] = v[0]
					}
				}
			}
		}
	}
	// 处理客户端的DELETE请求
	// 记录当前时间 `2006-01-02 15:04:05` 是指的格式格式
	data["time"] = time.Now().Format("2006-01-02 15:04:05")
	m := responseToClient{200, "success", data}
	mjson, e := json.Marshal(m)
	if e != nil {
		fmt.Println(e)
	}
	// 以json格式响应给客户端
	fmt.Fprintf(w, "%v\n", string(mjson))

}
func main() {
	http.HandleFunc("/", defaultFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
