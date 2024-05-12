package main

import (
	"GoStudy0328/src/79_Chat/common/message"
	process2 "GoStudy0328/src/79_Chat/server/process"
	"GoStudy0328/src/79_Chat/server/utils"
	"io"
	"log"
	"net"
)

type Processor struct {
	conn net.Conn // 连接
}

// ServerProcessMes 根据客户端发送消息种类不同 决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	// 是否能接收到客户端发送的消息
	log.Println("Client.mes =", mes)
	log.Println("mes.Type =", mes.Type)
	switch mes.Type {
	case message.LoginMesType:
		{
			// 处理登录的逻辑
			//err = ServerProcessLogin(conn, mes)
			up := &process2.UserProcess{ // 创建一个UserProcessor
				Conn: this.conn,
			}
			_ = up.ServerProcessLogin(mes)
		}
	case message.RegisterMesType:
		{
			// 处理注册
			up := &process2.UserProcess{ // 创建一个UserProcessor
				Conn: this.conn,
			}
			_ = up.ServerProcessRegister(mes) // 这个mes里面有 1)type 2)data
			// 然后先去写userDao
		}
	case message.SmsMesType:
		{
			// 在客户端每次要创建新的up
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes) // 是指针

		}

	default:
		log.Println("消息类型不存在 无法处理")

	}
	return
}

func (this *Processor) Process2() (err error) {
	for {
		// 不建议在这里乱写 建议封装为函数 readPkg
		// 这里将读取消息的 直接封装为一个函数 readPkg(conn net.Conn) 返回message和error
		tf := &utils.Transfer{
			Conn: this.conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				log.Println("客户端关闭了连接，服务器端也正常退出")
				return err
			} else {
				log.Println("read err:", err)
				return err
			}
			//return // 因为这里的return就不会一直读取
		}
		log.Println("读取到客户端的消息 =", mes)
		// 3)根据反序列化后对应的消息 判读是否是合法的用户 返回LoginResMes
		// 但是要根据不同消息来 让这个协程调用不同函数
		// ServerProcessMes() 处理消息
		// ServerProcessLoginMess() 处理登录请求
		err = this.serverProcessMes(&mes)
		if err != nil {
			log.Println("server process err:", err)
			return err
			// 以后要做错误处理
		}

	}

}
