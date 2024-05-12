package model

import (
	"GoStudy0328/src/79_Chat/common/message"
	"net"
)

// 因为正在客户端很多地方会用到 因此将此作为全局

type CurUser struct {
	Conn net.Conn
	message.User
}
