package main

import (
	"log"
	"net"
	"net/rpc"
)

// 1.定义一个远程方法
type Hello struct {
}

// SayHello
// 基本结构 方法中必须有一个指针类型的
// request 获取服务器发送的消息
// reply 返回消息需要使用指针类型
// req和rep不能是chan func 这两个都不能被序列化
func (h *Hello) SayHello(request string, reply *string) (err error) {
	log.Println("客户端传回: ", request)
	*reply = "Hello" + request
	return nil
}
func main() {
	log.Println("Hello Server")

	// 2.实现微服务
	err := rpc.RegisterName("hello", new(Hello)) // 注册微服务
	if err != nil {
		log.Println(err)
	}

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("开始建立连接")
		defer listener.Close() // 应用退出的时候关闭监听
		// 建立连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// 绑定服务
		rpc.ServeConn(conn)
	}

}
