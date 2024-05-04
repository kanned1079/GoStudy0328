package main

import "fmt"

// 写一个函数 完成登录

func login(userId int, userPassword string) (err error) {
	// 注意给函数或参数做好命名
	// 最好是使用error来反回错误 而不是使用bool
	// 下一步是定义协议 怎么组织发送的数据
	fmt.Println("你输入的用户名和密码 = ", userId, userPassword)

	err = nil
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

type Message struct {
	Type string `json:"type"` // 这里以后会有很多种不同的类型
	Data string `json:"data"` // 这里是序列化后放入的

}
