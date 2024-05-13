package process

import (
	"GoStudy0328/Basic_Syntax/src/79_Chat/common/message"
	"GoStudy0328/Basic_Syntax/src/79_Chat/server/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

// SendGroupMes 群发消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	// 创建一个Message
	var mes message.Message
	mes.Type = message.SmsMesType
	// 创建一个smsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content        // 说得内容
	smsMes.UserId = CurrUser.UserId //
	smsMes.UserStatus = CurrUser.UserStatus
	// 序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes.smsMes 序列化失败, err =", err)
		return
	}
	mes.Data = string(data)
	// 再次序列化sms
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes.mes 序列化失败, err =", err)
		return
	}
	// 将mes发送到服务器
	tf := &utils.Transfer{
		Conn: CurrUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes.tf.Write() 错误, err =", err)
		return
	}
	return
	// 写完客户端发送消息后去写服务端
}
