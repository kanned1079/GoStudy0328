package message

type User struct {
	UserId       int    `json:"user_id"`
	UserPassword string `json:"user_password"`
	UserName     string `json:"user_name"` // 后面序列化 这里肯定要用json
} // 先声明一个用户的结构体
