package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	fmt.Println("client")
	// 1.建立连接
	conn, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// 使用conn来调用远程处理函数
	var rep string
	err = conn.Call("hello.SayHello", "我是客户端", &rep)
	if err != nil {
		log.Println(err)
	}
	// 4.获取微服务返回的数据
	fmt.Printf(rep)
}
