package main

import (
	"GoStudy0328/src/79_Chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 写一个函数 完成登录

func login(userId int, userPassword string) (err error) {
	// 注意给函数或参数做好命名
	// 最好是使用error来反回错误 而不是使用bool
	// 下一步是定义协议 怎么组织发送的数据
	//fmt.Println("你输入的用户名和密码 = ", userId, userPassword)
	// 开始完成用户登录的第一个功能 获取消息的长度 并且服务端可以正常收到该值
	// 1.连接到服务器端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	// 2. 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//mes.Data // 这个要序列化后放进去
	// 3.创建一个LoginMes结构体
	var loginMes message.LoginMes // 要序列化这个给上面的Mes
	loginMes.UserId = userId
	loginMes.UserPassword = userPassword
	fmt.Println(loginMes)
	//mes.Data = loginMes // 这样不行
	// 4.将LoginMes序列化
	data, err := json.Marshal(loginMes) // 这里的data是[]byte切片
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	// 5.将data赋给Message的Data字段
	mes.Data = string(data) // 将[]byte转换为string
	// 6.将Mes进行序列化
	data, err = json.Marshal(mes) // 这里的data就是需要发送的数据了 这是个byte切片
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	// 7.到这里data就是要发送的消息 根据规则要先发送消息长度
	// 7.1.发送消息长度 先获取到data的长度->转换成一个表示长度的[]byte切片
	// binary包实现了简单的数字与字节序列的转换以及变长值的编解码。
	// ByteOrder规定了如何将字节序列和 16、32或64比特的无符号整数互相转化。
	var pkgLen uint32                               // 这个大小是 (2^32)-1字节 大约有4G
	pkgLen = uint32(len(data))                      // 把int转换为uint32
	var buffer [4]byte                              // 存放上main的len 发送需要的是[]byte
	binary.BigEndian.PutUint32(buffer[0:4], pkgLen) // 相当于把长度转换为了bytes
	// 发送长度
	n, err := conn.Write(buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	fmt.Println("客户端发送消息成功 长度 = ", pkgLen)
	fmt.Println("发送的data = ", string(data))
	// Class3 update20240506
	// 完成客户端可以发送消息本身 服务端能正常收到消息 并根据客户端发送的的消息判断是不是合法用户
	// 并返回相应的LoginResMes
	// 1)让客户端发送消息本身
	// 2)服务器收到收将其反序列化成对应的消息结构体
	// 3)根据反序列化后对应的消息 判读是否是合法的用户 返回LoginResMes
	// 4)客户端根据接收到的消息显示对应的页面
	// 5)这里需要做一些函数的封装
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	// 这里还需要进行处理服务器返回的消息
	return
}

// 发送数据的流程
//	1)创建一个Message结构体
//	2)设置消息类型 mes.Type = 登录消息类型
//	3)mes.Data = 登录消息的序列化后的数据
//	4)对mes进行序列化
//	5)在网络传输中需要解决丢包 那么怎么防止丢包
//		1)先不发送mes本身 先发送mes长度【有多少字节】
//		2)再发送消息本身

// 接受数据的流程
// 1)先接收到客户端发送的消息长度 []byte -> int
// 2)根据接收到的长度len来再次接收消息本身
// 3)接收时要判断实际接收到的消息内容是否等于len
// 4)如果不想等 就有纠错协议
// 5)取到消息后 反序列化->Message
// 6)取出message.Data(string) 反序列化
// 7)取出LoginMes.userId和LoginMes.userPasswd
// 8)进行比较
// 9)根据比较结果 返回Mess 里面比较简单 比如1成功 100错误等等
// 10)发送给客户端

// Message 定义消息的格式和它的结构
