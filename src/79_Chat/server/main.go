package main

import (
	"GoStudy0328/src/79_Chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buffer := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	// 这里的conn只有在没有被关闭的情况下 才会阻塞
	// 如果客户端关闭了conn就不会阻塞了
	_, err = conn.Read(buffer[:4]) // 先读取前面四个字节
	if err != nil {
		fmt.Println("read err:", err) // 这里是读取头部长度出错
		//err = errors.New("read pkg header err")
		// 进行错误处理
		return
	}
	fmt.Println("读取到buf长度 = ", buffer[:4])
	// 将这个长度转换为uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buffer[0:4]) // 这里是Uint32转回来 前面是Putuint32
	// 根据pkgLen来读取内通
	n, err := conn.Read(buffer[:pkgLen]) // 从conn套接字中读取[:pkgLen]长度的内容到buffer里
	if n != int(pkgLen) || err != nil {  // 第一种可能读到了但是没有读全 还有一种就是对方关闭了管道
		fmt.Println("read err:", err) // 这里是读取内容出错
		//err = errors.New("read pkg body err")
		return
	} // 到这里buffer里面就有收到的消息了 但是还没完 这需要反序列化
	err = json.Unmarshal(buffer[:pkgLen], &mes) // 注意这里要加上& 很重要 很重要 很重要
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	return
}

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 也需要延迟关闭
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("close conn err:", err)
		}
	}()
	//buffer := make([]byte, 8096)
	// 读取客户端发送的信息 这里该去写客户端的代码了
	for {
		// 不建议在这里乱写 建议封装为函数 readPkg
		// 这里将读取消息的 直接封装为一个函数 readPkg(conn net.Conn) 返回message和error
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了连接，服务器端也正常退出")
				return
			} else {
				fmt.Println("read err:", err)
				return
			}
			//return // 因为这里的return就不会一直读取
		}
		fmt.Println(mes)
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
