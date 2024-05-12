package process2

import (
	"fmt"
)

// 因为UserMgr实例在服务器端只会有一个
// 因为在很多地方都会用到 因此定义为全局变量

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对UserMger的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUser的增删改查
// AddOnlineUser 增加 如果修改也使用这个 因为map中的key是唯一的
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// DelOnlineUser 删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// GetOnlineUser 返回在线的用户
func (this *UserMgr) GetOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// GetOnlineUserById 根据Id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//return this.onlineUsers[userId]
	// 如何从map中取出一个值 带检测的方式
	up, ok := this.onlineUsers[userId]
	if !ok { // 这里说明要查找的这个用户当前不在线
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}
