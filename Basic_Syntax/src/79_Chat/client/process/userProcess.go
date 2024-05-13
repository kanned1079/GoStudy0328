package process

import (
	"GoStudy0328/Basic_Syntax/src/79_Chat/client/utils"
	"GoStudy0328/Basic_Syntax/src/79_Chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
	// 暂时不用任何字段
}

// Login 只处理登录功能
func (this *UserProcess) Login(userId int, userPassword string) (err error) {
	// 注意给函数或参数做好命名
	// 最好是使用error来反回错误 而不是使用bool
	// 下一步是定义协议 怎么组织发送的数据
	//fmt.Println("你输入的用户名和密码 = ", userId, userPassword)
	// 开始完成用户登录的第一个功能 获取消息的长度 并且服务端可以正常收到该值
	// 1.连接到服务器端
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	// 2. 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//mes.Data // 这个要序列化后放进去
	// 3.创建一个LoginMes结构体
	var loginMes message.LoginMes // 要序列化这个给上面的Mes
	loginMes.UserId = userId
	loginMes.UserPassword = userPassword
	fmt.Println("func(Login) >", loginMes)
	//mes.Data = loginMes // 这样不行
	// 4.将LoginMes序列化
	data, err := json.Marshal(loginMes) // 这里的data是[]byte切片
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	// 5.将data赋给Message的Data字段
	mes.Data = string(data) // 将[]byte转换为string
	// 6.将Mes进行序列化
	data, err = json.Marshal(mes) // 这里的data就是需要发送的数据了 这是个byte切片
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	// 7.到这里data就是要发送的消息 根据规则要先发送消息长度
	// 7.1.发送消息长度 先获取到data的长度->转换成一个表示长度的[]byte切片
	// binary包实现了简单的数字与字节序列的转换以及变长值的编解码。
	// ByteOrder规定了如何将字节序列和 16、32或64比特的无符号整数互相转化。
	var pkgLen uint32                               // 这个大小是 (2^32)-1字节 大约有4G
	pkgLen = uint32(len(data))                      // 把int转换为uint32
	var buffer [4]byte                              // 存放上main的len 发送需要的是[]byte
	binary.BigEndian.PutUint32(buffer[0:4], pkgLen) // 相当于把长度转换为了bytes
	// 发送长度
	n, err := conn.Write(buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	fmt.Println("客户端发送消息成功 长度 = ", pkgLen)
	fmt.Println("发送的data = ", string(data))
	// 这是一串信息
	// "{\"user_id\":100,\"user_password\":\"123456\",\"user_name\":\"kanna\"}"
	// Class3 update20240506
	// 完成客户端可以发送消息本身 服务端能正常收到消息 并根据客户端发送的的消息判断是不是合法用户
	// 并返回相应的LoginResMes
	// 1)让客户端发送消息本身
	// 2)服务器收到收将其反序列化成对应的消息结构体
	// 3)根据反序列化后对应的消息 判读是否是合法的用户 返回LoginResMes
	// 4)客户端根据接收到的消息显示对应的页面
	// 5)这里需要做一些函数的封装
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	// 这里还需要进行处理服务器返回的消息
	//time.Sleep(10 * time.Second)
	//fmt.Println("休眠了10s")
	tf := utils.Transfer{ // 创建Transfer实例
		Conn: conn,
	}
	mes, err = tf.ReadPkg() // mes还需要反序列化
	if err != nil {
		fmt.Println("readPkg err:", err)
		return
	}
	// 将mes的Data部分发序列化成LoginRespMes
	var loginResMes message.LoginRespMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// 初始化CurrUser
		CurrUser.Conn = conn
		CurrUser.UserId = userId
		CurrUser.UserStatus = message.UserOnline
		// 显示在线用户的列表 遍历loginResMes.UserIds
		fmt.Println("当前在线用户列表如下: ")
		for _, v := range loginResMes.UserIds {
			// 如果要求不显示自己在线 下面需要增加的
			if v == loginMes.UserId {
				continue
			}
			fmt.Println("用户Id： ", v)
			// 在这里把Id放到onlineUser中
			user := &message.User{ // 这里是用的指针
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user // 这里的value也是要的指针

		}
		fmt.Println("\n\n")
		go serverProcessMes(conn) // 在这里启动协程
		go fmt.Println("登录成功")    // 要修改这里
		for {                     // 显示登录成功的菜单 且要做循环显示
			ShowMenu()
			// 下面要做保持服务器沟通 因此需要在客户端启动一个协程
			// 如果服务器有数据推送过来给客户端就接受并显示在客户端的终
		}

	} else { // 修改了这里 根据状态码直接返返回
		fmt.Println("登录失败 可能的原因 err = ", loginResMes.Error)
	}
	return
}

func (this *UserProcess) Register(userId int, userPassword string, userName string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	// 复制的Login的代码
	var mes message.Message
	mes.Type = message.RegisterMesType // 修改消息类型
	//mes.Data // 这个要序列化后放进去
	// 3.创建一个LoginMes结构体
	var registerMes message.RegisterMes // 要序列化这个给上面的Mes
	registerMes.User.UserId = userId
	registerMes.User.UserPassword = userPassword
	registerMes.User.UserName = userName
	// 下面要完成序列化 将regMes序列化
	data, err := json.Marshal(registerMes) // 这里的data是[]byte切片
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	mes.Data = string(data) // 将[]byte转换为string
	// 将mes序列化
	data, err = json.Marshal(mes) // 这里的data就是需要发送的数据了 这是个byte切片
	fmt.Println("发送前 ", mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	// 需要封装上面的
	// 创建一个Transfer
	tf := utils.Transfer{
		Conn: conn,
	}
	// 发送data给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("请求注册信息 writePkg err:", err)
	}
	mes, err = tf.ReadPkg() // 这个mes就是RegisterRespMes
	if err != nil {
		fmt.Println("请求注册信息 readPkg err:", err)
	}
	// 将mes的Data部分发序列化成RegisterRespMes
	var registerRespMes message.RegisterRespMes
	err = json.Unmarshal([]byte(mes.Data), &registerRespMes)
	if registerRespMes.Code == 200 { // 正确
		fmt.Println("注册成功 需要重新登录")
		os.Exit(0) // pass
	} else { // 修改了这里 根据状态码直接返返回
		fmt.Println("注册失败 可能的原因 err = ", registerRespMes.Error)
		os.Exit(0) // pass
	}

	return
}
