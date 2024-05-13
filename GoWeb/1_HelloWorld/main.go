package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World", r.URL.Path)
	log.Println("请求路径: ", r.URL.Path)
}

// type HandlerFunc func(ResponseWriter, *Request)
// HandlerFunc type是一个适配器，通过类型转换让我们可以将普通的函数作为HTTP处理器使用。
// 如果f是一个具有适当签名的函数，HandlerFunc(f)通过调用f实现了Handler接口。

func main() {
	http.HandleFunc("/", handler)     // 请"/"时候 调用handler处理器
	http.ListenAndServe(":8080", nil) // 这里写nil就使用默认的多路复用器DefaultServeMux
	// 上面端口如果不写那么就默认使用80
	// ListenAndServe监听srv.Addr指定的TCP地址，并且会调用Serve方法接收到的连接。
	// 如果srv.Addr为空字符串，会使用":http"。
}
