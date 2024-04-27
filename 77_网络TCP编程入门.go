package main

// Golang的主要设计目标就是面向大规模后端服务程序，网络通信这块是服务端程序必不可少关重要的一部分。
// 网络编程有两种:
// 1)TCP socket编程，是网络编程的主流。之所以叫Tcp socket编程，是因为底层是基于Tcp/ip协议的.比如:QQ聊天[示意图]
// 2)B/S结构的http编程，我们使用浏览器去访问服务器时，使用的就是http协议，而http底层依旧是用tcp socket实现的。[示意图] 比如: 京东商城 【这属于go web开发范畴】

// 怎样去描述一个message
type Message77 struct {
	ID       int
	content  string
	SendMame string
	Getter   map[string]string
}

type Connection struct {
}

// 服务器)要尽可能的少开端口在计算机(尤其是做服务器尽量少开端口
// 一个程序监听一个端口只能被
// 如果使用 netstat -an 可以查看本机有哪些端口在监听
// 可以使用 netstat -anb 来查看监听端口的pid,在结合任务管理器关闭不安全的端口

// 服务端的处理流程
// 1)监听端口 8888
// 2)接收客户端的tcp链接，建立客户端和服务器端的链接
// 3)创建goroutine，处理该链接的请求(通常客户端会通过链接发送请求包)

// 客户端的处理流程
// 1)建立与服务端的链接
// 2)发送请求数据，接收服务器端返回的结果数据
// 3)关闭链接

func main() {

}
