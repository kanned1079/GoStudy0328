package process

import (
	"GoStudy0328/Basic_Syntax/src/79_Chat/common/message"
	"encoding/json"
	"fmt"
	"log"
)

func outputGroupMes(mes *message.Message) {
	// 上面来的消息肯定是SmsMestype来得
	// 显示消息
	var smsMes message.SmsMes
	if err := json.Unmarshal([]byte(mes.Data), &smsMes); err != nil {
		log.Println("smsMes反序列化失败 err: ", err)
		return
	}
	// 显示
	info := fmt.Sprintf("用户Id(%v)说: %v\n", smsMes.UserId, smsMes.Content)
	fmt.Println(info)

}
