package main

import (
	"fmt"
	"net"
)

const (
	ipaddr string = "localhost"
	port   int    = 8080
)

func main() {
	//func Dial(network, address string) (Conn, error)
	//在网络network上连接地址address，并返回一个Conn接口。
	//对TCP和UDP网络，地址格式是host:port或[host]:port，参见函数JoinHostPort和SplitHostPort
	var err error
	// 创建net.Dial进行连接服务器 参数和net.Listen相同
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipaddr, port))
	if err != nil {
		fmt.Println("连接错误 ERR =", err)
	} else {
		fmt.Println("服务器连接成功 conn = ", conn)
	}

	//// 功能1 客户选可以发送单行数据 然后就退出
	////reader := bufio.NewReader(os.Stdin) // 这里的os.Stdin 是标准化输入的意思 standard input
	//
	//// 从终端读取一行用户输入 准备发送给服务器
	////line, err := reader.ReadString('\n') // 以换行符号为结束符
	//fmt.Print("输入消息 > ")
	//var str string
	//_, _ = fmt.Scan(&str)
	////if err != nil {
	////	fmt.Println("readString错误 ERR = ", err)
	////}
	//// 再将liene发送给服务器
	//// 客户端使用conn.Write来向服务器发送消息 注意这里需要的值是字符切片
	//n, err := conn.Write([]byte(str)) // 强制转换为[]byte
	//if err != nil {
	//	fmt.Println("conn.Write错误 ERROR = ", err)
	//} else {
	//	fmt.Printf("客户端发送了%v字节的数据 并退出", n)
	//}
	startWorker(conn)

	// 创建消息通道
	//msgChan := make(chan string)
	//go receiveAndDisplayMessages(conn, msgChan)

	// 客户端也要记得关闭连接
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("客户端关闭连接错误 ERROR = ", err)
		}
	}()

	//defer conn.Close()
	// 到这里结束 接下来是做服务器的接受
}

func startWorker(conn net.Conn) {
	for {
		fmt.Print("输入消息 > ")
		var str string
		_, _ = fmt.Scanln(&str)
		if str == "exit" {
			return
		}
		_, err := conn.Write([]byte(str)) // 强制转换为[]byte
		if err != nil {
			fmt.Println("conn.Write错误 ERROR = ", err)
		} else {
			//fmt.Printf("客户端发送了%v字节的数据\n\n", n)
		}
	}
}

func receiveAndDisplayMessages(conn net.Conn, msgChan chan string) {
	for {
		msg := <-msgChan
		fmt.Println("服务器发来消息:", msg)
	}
}
