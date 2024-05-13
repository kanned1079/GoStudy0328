package process

import (
	"GoStudy0328/Basic_Syntax/src/79_Chat/common/message"
	"GoStudy0328/Basic_Syntax/src/79_Chat/server/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

// 保持和服务器连接

// ShowMenu 显示登录成功的菜单
func ShowMenu() {
	fmt.Println("---------------xxx登录成功---------------")
	fmt.Println("\t\t1.显示在线用户列表")
	fmt.Println("\t\t2.发送消息")
	fmt.Println("\t\t3.信息列表")
	fmt.Println("\t\t4.退出")
	fmt.Print("请选择(1-4) >")
	var key int
	var content string
	//up := &UserProcess{}
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		// 前面完成了到这里修改 添加一个方法
		outputOnlineUser()
	case 2:
		{
			//fmt.Println("发送消息")
			fmt.Print("输入消息(群聊) > ")
			_, _ = fmt.Scanf("%s\n", &content)
			// 创建到外面
			//up := &UserProcess{}
			smsProcess.SendGroupMes(content)

		}

	case 3:
		fmt.Println("查看信息列表")
	case 4:
		{
			fmt.Println("退出了系统")
			// 最好是要通知服务器要关闭连接了
			os.Exit(0)
		}
	default:
		fmt.Println("非法输入")

	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	} // 创建一个Transfer实例 不停向服务器发送消息
	for { // 保持通讯
		fmt.Println("客户端正在等服务器发送的消息...")
		mes, err := tf.ReadPkg() // 这里的mes是已经是序列化后的
		if err != nil {
			fmt.Println("服务器端 tf.ReadPkg() err:", err)
			return
		}
		// 如果读取到消息 又是下一步处理逻辑
		log.Println("读取到服务器发送的消息 =", mes)
		switch mes.Type {
		case message.NotifyUserStatusType: // 通知有人上线了
			{
				// 那么就要在这里做处理
				// 1)取出NotifyUSerStatusMes 使用反序列化
				var notifyUSerStatusMes message.NotifyUserStatusMes
				err = json.Unmarshal([]byte(mes.Data), &notifyUSerStatusMes)
				if err != nil {
					fmt.Println("json.Unmarshal err:", err)
					return
				}
				// 2)把这个用户的状态加入到客户端维护的map中去 现在要去写map的
				updateUserStatus(&notifyUSerStatusMes)
				// 编写一个方法处理返回的信息
			}
		case message.SmsMesType:
			{
				log.Println("服务器发送的消息类型是 = ", message.SmsMesType)
				outputGroupMes(&mes)
			}

		default:
			fmt.Println("服务器端返回了一个未知消息 Type =", mes.Type)
		}
		//fmt.Printf("mes = %v\n", mes)

	}
}
