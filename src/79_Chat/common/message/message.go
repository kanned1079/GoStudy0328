package message

// 确定一些消息类型
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` // 消息类型 这个类型最好是使用一些常量
	Data string `json:"data"` // 消息内容
}

// 先定义两个具体的消息 后面需要的话再增加 一个是登录消息 一个是登录结果的消息
type LoginMes struct {
	UserId       int    `json:"user_id"`       // 用户Id
	UserPassword string `json:"user_password"` // 用户密码
	UserName     string `json:"user_name"`     // 用户名
}

type LoginRespMes struct {
	Code  int    `json:"code"`  // 状态码 500表示用户未注册 200表示登录成功
	Error string `json:"error"` // 返回的错误信息
}

type RegisterMes struct {
	// pass
}