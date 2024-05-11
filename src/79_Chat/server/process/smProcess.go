package process2

import (
	"GoStudy0328/src/79_Chat/common/message"
	"GoStudy0328/src/79_Chat/server/utils"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// 服务器的群发消息
type SmsProcess struct {
}

// 转发消息
func (this *SmsProcess) SendGroupMes(msg *message.Message) {
	// 遍历服务器端的userMgr的map
	// 将消息转发出去
	// 先取出msg的内容 在Sms中
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(msg.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes.json 反序列化失败 err:", err)
		return
	}
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes.json 反序列化失败 err:", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		// 这里还需要过滤到自己
		if id == smsMes.UserId {
			continue
		}
		// 新建一个函数用来发送
		if err := this.SendMsgToEachOnlineUser(data, up.Conn); err != nil {
			fmt.Println("转发消息失败")
			log.Println("转发消息失败")
		}
		log.Println("成功转发了消息")

	}
}

// SendMsgToEachOnlineUser 单独封装成一个方法用来发送
func (this *SmsProcess) SendMsgToEachOnlineUser(data []byte, conn net.Conn) (err error) {
	// 创建Transfer
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	return
}
