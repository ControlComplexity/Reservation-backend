package message

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// WebSocket部署服务器外网无法连接解决方案: 1. 服务器的监听地址localhost改为0.0.0.0  2. 腾讯云的安全组里面放行 3.linux开通防火墙放行
var addr = flag.String("addr", "0.0.0.0:10020", "http service address")
var ch1 chan string
var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("echo1")

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	fmt.Println("echo ch1: ", ch1)
	defer c.Close()
	err = c.WriteMessage(1, []byte("kkkkkk"))
	go func() {
		for {
			err = c.WriteMessage(1, []byte("WebSocket测试"))
			time.Sleep(5 * time.Second)
		}
	}()
	for {
		select {
		case msg := <-ch1:
			fmt.Println("msg content: ", msg)
			err = c.WriteMessage(1, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}

}

func Send(ch chan string) {
	fmt.Println("ch11: ", ch1)
	ch1 = ch
	fmt.Println("ch12: ", ch1)
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/message", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
