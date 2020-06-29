package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/polar9527/learning-go/golangWebsocketVue/common"
)

var modeDescribe = `webSocket response model :
--echo default value
--robot The Turing robot responded
`
var addr = flag.String("addr", "", "http service addr")
var model = flag.String("model", "", modeDescribe)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echoFunc(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrader :", err)
		return
	}
	defer func() {
		log.Println(c.RemoteAddr().String(), "connect close")
	}()
	defer c.Close()
	log.Println(c.RemoteAddr().String(), "connect success")
	var robot = common.NewRobot()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		//t := time.Now().Format("2006-01-02 15:04:05")
		m := string(message)
		log.Println("server receive message :", m)
		// robot 机器人应答
		if *model == "robot" {
			err, result := robot.Chat(m)
			if err == nil {
				for _, k := range result {
					if s, ok := k.(string); ok {
						message = []byte(s)
						err = c.WriteMessage(mt, message)
					}
				}
			}
			log.Println(err)
		} else if *model == "echo" {
			// echo 回音服务
			log.Println("echo 回音服务")
			message = []byte(m)
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("write :", err)
			}
		}
	}
}
func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("addr : ", *addr)
	log.Println("model :", *model)
	http.HandleFunc("/echo", echoFunc)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
