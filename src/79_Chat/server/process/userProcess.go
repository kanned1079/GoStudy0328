package process2

import (
	"GoStudy0328/src/79_Chat/common/message"
	"GoStudy0328/src/79_Chat/server/model"
	"GoStudy0328/src/79_Chat/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn // 连接
}

// ServerProcessLogin 专门处理登录的请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 1.先从mes中取出data 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes) // 将string强制转换为[]byte给反序列化
	if err != nil {
		fmt.Println("mes.Data反序列化失败 err = ", err)
		return
	}
	// 先声明一个返回的ResMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 再声明一个LoginResMes 因为需要封装
	var loginResMes message.LoginRespMes // 要把这个填到上面去才能返回

	// 先不到数据库了去 如果用户的ID为100 密码为123456就是合法的 反之
	// 到数据库里去验证
	// 1)使用model.UserDao去数据库验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPassword)
	if err != nil {
		// 这里需要封装
		if err == model.ERROR_USER_NOT_EXIST { // 	用户不存在
			loginResMes.Code = 500
			//loginResMes.Error = "该用户不存在 请注册再使用"
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403 // 密码不正确
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505 // 未知错误
			loginResMes.Error = "未知服务器内部错误..."
		}
		fmt.Println(user, "登录失败 by serverProcessLogin")
		// 先测试成功 后面再根据返回具体错误信息
	} else {
		// 没有错误
		loginResMes.Code = 200 // 合法表示200 登录成功
		fmt.Println(user, "登录成功 by serverProcessLogin")
	}
	//if loginMes.UserId == 100 && loginMes.UserPassword == "123456" { // 需要改这里
	//	// 这样就是合法的 先不到数据库里去认证
	//	loginResMes.Code = 200 // 合法表示200 登录成功
	//	loginResMes.Error = "用户登录成功"
	//} else {
	//	// 不合法的话就构建一个LoginResMes就返回给客户端
	//	loginResMes.Code = 500              // 不合法表示 该用户不存在
	//	loginResMes.Error = "该用户不存在 请注册再使用" // 注释
	//}
	// 3.将loginresMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
		return
	}
	// 4. 将data赋值给resMes
	resMes.Data = string(data)
	// 5. 再对resMes进行序列化 发送至客户端需要
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
		return
	}
	// 6. 发送data 不能够直接发 也要防止丢包 还有很多步骤 因此这里也用封装成一个函数
	// 将其封装到一个writePkg中
	tf := &utils.Transfer{ // 新建实例
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	//err = writePkg(conn, data) // 现在这个已经没有了 需要调用Transfer对象的方法
	// 下面要去客户端写
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 	前面的userDao入库操作写完后来这里
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes) // 将string强制转换为[]byte给反序列化
	if err != nil {
		fmt.Println("mes.Data反序列化失败 err = ", err)
		return
	}
	// 先声明一个返回的ResMes
	var resMes message.Message
	resMes.Type = message.RegisterRespMesType
	var registerRespMes message.RegisterRespMes // 要把这个填到上面去才能返回
	// 下面要去Redis中完成注册
	// 这里的user因为不是一个包来得所以还是会报错
	//var mesUser = message.User{registerMes.User.UserId, registerMes.User.UserPassword, registerMes.User.UserName}
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_NOT_EXIST {
			registerRespMes.Code = 505
			registerRespMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerRespMes.Code = 506
			registerRespMes.Error = "ServerProcessRegister 注册时候发生未知错误"
		}
	} else {
		registerRespMes.Code = 200
		fmt.Println("ServerProcessRegister 成功")
	}
	// 下面开始组合数据
	data, err := json.Marshal(registerRespMes)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
		return
	}
	// 6. 发送data 不能够直接发 也要防止丢包 还有很多步骤 因此这里也用封装成一个函数
	// 将其封装到一个writePkg中
	tf := &utils.Transfer{ // 新建实例
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return err
}
