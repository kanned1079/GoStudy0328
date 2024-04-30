package main

import (
	"fmt"
	"io"
	"net"
)

// type Listener interface {
//    // Addr返回该接口的网络地址
//    Addr() Addr
//    // Accept等待并返回下一个连接到该接口的连接
//    Accept() (c Conn, err error)
//    // Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
//    Close() error
//}

const (
	ipaddr string = "localhost"
	port   string = "8888"
)

// 先新建Listen 结束前这个要Close
// 再新建conn  这个也要

func main() {
	fmt.Printf("服务器正在监听 %s:%s......\n", ipaddr, port)
	//fmt.Printf("服务器正在监听%s的%s端口.......\n", ipaddr, port)
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ipaddr, port))
	if err == io.EOF {
		fmt.Println("Listen ERROR =", err)
		return
	}

	defer func() { // 只要main函数不退出就不关闭
		if err = listen.Close(); err != nil {
			fmt.Println("关闭连接失败 ERROR =", err)
		}
	}()

	// 循环等待客户端来连接
	for {
		fmt.Println("等待客户端连接......")
		// 创建conn
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Listen Accept 失败")
			return
		} else {
			fmt.Println("Accept成功 远程主机地址 = ", conn.RemoteAddr().String())
		}
		// 这里准备启动协程为客户端服务
		// 这里不要写在main中
		//go process(conn)
		go myProcess(conn)
		// 每当有一个客户端连接 这里就开启一个协程
	}

	//fmt.Println("listen = ", listen)
}

//func process(conn net.Conn) {
//	// 必须要拿到一个Conn
//	// 这里进行循环接收
//	defer conn.Close() // 这里很重要 要记得关闭 不然别的客户端不能连接
//	// 循环接收
//	for {
//		// 创建一个新的切片
//		buf := make([]byte, 1024)
//		// 1)等待客户端通过conn发消息
//		// 2)如果客户端没有write 那么这个协程会阻塞在这里
//		fmt.Printf("服务器在等待%v发送的数据......\n", conn.RemoteAddr())
//		n, err := conn.Read(buf)
//		if err != nil {
//			fmt.Println("服务器端Read出错 ERR = ", err)
//			return
//		}
//		// 3. 在服务器终端显示客户端内容
//		fmt.Print(string(buf[:n])) // 这里记得不要用ln
//		fmt.Println()
//	}
//}

func myProcess(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("连接关闭失败 ERROR =", err)
		}
	}()
	for {
		buffer := make([]byte, 512)
		fmt.Printf("服务器正在等待%v发送的数据......\n", conn.RemoteAddr().String())
		length, err := conn.Read(buffer)
		//if err != nil {
		//	fmt.Println("服务端Read出错 ERROR =", err)
		//	return
		//}
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}
		fmt.Println("消息长度 = ", length)
		fmt.Println("Message = ", string(buffer[:length])) // 这里的length是因为要取切片的第0到第length-1个元素 否则会将切片全部转换
		fmt.Println()
	}
}
