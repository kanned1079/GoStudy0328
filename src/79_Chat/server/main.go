package main

import (
	"GoStudy0328/src/79_Chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// readPkg 用于读取数据包
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

// writePkg 用于向客户端再发送数据 包含与客户端发送相似逻辑
// 过来的信息已经是字节了就不需要再转换成字符串了
func writePkg(conn net.Conn, data []byte) (err error) {
	// 1.先发送长度给客户端
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buffer [4]byte
	binary.BigEndian.PutUint32(buffer[0:4], pkgLen)
	n, err := conn.Write(buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("server write err:", err)
		return
	}
	// 2. 发送data消息本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("server write err:", err)
		return
	}
	return
}

// ServerProcessLogin 专门处理登录的请求
func ServerProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 1.先从mes中取出data 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), loginMes) // 将string强制转换为[]byte给反序列化
	if err != nil {
		fmt.Println("mes.Data反序列化失败 err = ", err)
		return
	}
	// 先声明一个返回的ResMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 再声明一个LoginResMes 因为需要封装
	var loginResMes message.LoginRespMes // 要把这个填到上面去才能返回

	// 先不到数据库了去 如果用户的ID为100 密码为123456就是合法的 反之
	if loginMes.UserId == 100 && loginMes.UserPassword == "123456" {
		// 这样就是合法的 先不到数据库里去认证
		loginResMes.Code = 200 // 合法表示200 登录成功
		loginResMes.Error = "用户登录成功"
	} else {
		// 不合法的话就构建一个LoginResMes就返回给客户端
		loginResMes.Code = 500              // 不合法表示 该用户不存在
		loginResMes.Error = "该用户不存在 请注册再使用" // 注释
	}
	// 3.将loginresMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
		return
	}
	// 4. 将data赋值给resMes
	resMes.Data = string(data)
	// 5. 再对resMes进行序列化 发送至客户端需要
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
		return
	}
	// 6. 发送data 不能够直接发 也要防止丢包 还有很多步骤 因此这里也用封装成一个函数
	// 将其封装到一个writePkg中
	err = writePkg(conn, data)
	// 下面要去客户端写
	return
}

// ServerProcessMes 根据客户端发送消息种类不同 决定调用哪个函数来处理
func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		{
			// 处理登录的逻辑
			err = ServerProcessLogin(conn, mes)
		}
	case message.LoginResMesType:
		{
			// 处理注册
		}
	default:
		fmt.Println("消息类型不存在 无法处理")

	}
	return
}

// process 处理和客户端的通讯
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
		// 3)根据反序列化后对应的消息 判读是否是合法的用户 返回LoginResMes
		// 但是要根据不同消息来 让这个协程调用不同函数
		// ServerProcessMes() 处理消息
		// ServerProcessLoginMess() 处理登录请求

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
