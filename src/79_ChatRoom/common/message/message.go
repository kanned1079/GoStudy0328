package message

// 确定一些消息类型
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string // 消息类型 这个类型最好是使用一些常量
	Data string // 消息内容
}

// 先定义两个具体的消息 后面需要的话再增加 一个是登录消息 一个是登录结果的消息
type LoginMes struct {
	UserId       int    // 用户Id
	UserPassword string // 用户密码
	UserName     string // 用户名
}

type LoginRespMes struct {
	Code  int    // 状态码 500表示用户未注册 200表示登录成功
	Error string // 返回的错误信息
}
