package main

import (
	"fmt"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 也需要延迟关闭
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("close conn err:", err)
		}
	}()
	buffer := make([]byte, 8096)
	// 读取客户端发送的信息 这里该去写客户端的代码了
	for {
		//buffer := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据...")
		_, err := conn.Read(buffer[:4]) // 先读取前面四个字节
		if err != nil {
			fmt.Println("read err:", err)
			// 进行错误处理
			return
		}
		fmt.Println("读取到buf长度 = ", buffer[:4])
		// 不建议在这里乱写 建议封装为函数 readPkg
	}
}

func main() {
	// 提示信息
	fmt.Println("服务器正在监听8889端口...")
	listener, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 延时关闭
	defer func() {
		if err := listener.Close(); err != nil {
			fmt.Println("listener close err:", err)
		}
	}()
	// 一旦监听成功 就等待客户端来连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			// 先不退出
		}
		// 一旦连接成功就启动一个协程 和客户端保持通讯
		go process(conn) // 将套接字socket传递给协程
	}
}
