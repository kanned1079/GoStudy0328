package main

import (
	"GoStudy0328/src/79_Chat/common/message"
	"fmt"
	"net"
)

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
