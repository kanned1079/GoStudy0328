package process

import (
	"GoStudy0328/src/79_Chat/client/model"
	"GoStudy0328/src/79_Chat/common/message"
	"fmt"
)

// 这个和服务器端不一样的是它管理的是客户端的
// 客户端要维护的map
// onlineUsers 在登录的时候会把所有Id放进去 那么在用户登录成功的时候才会把所有的值放进去
var onlineUsers map[int]*message.User = make(map[int]*message.User)
var CurrUser model.CurUser // 在用户登录成功后将会对这个初始化

// UpdateUserStatus 用于处理返回的NotifyUserOnlineMes信息
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	// 需要进行优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { // 如果没有的话 才会去创建
		user = &message.User{ // 移动到了这里
			UserId: notifyUserStatusMes.UserId,
			//UserStatus: notifyUserStatusMes.Status,	// 去除这行到下面
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	//user := &message.User{
	//	UserId:     notifyUserStatusMes.UserId,
	//	UserStatus: notifyUserStatusMes.Status,
	//} // 但是这样的效率不高 因为多创建了一次user 需要优化
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser() //更新完后顺便再调用显示
}

// 还需要在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线的用户：")
	// 对onlineUser进行遍历
	for id, user := range onlineUsers {
		fmt.Printf("用户Id：%d\t状态：%v\n", id, user.UserStatus)
	}

}
