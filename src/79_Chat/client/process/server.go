package process

import (
	"GoStudy0328/src/79_Chat/server/utils"
	"fmt"
	"net"
	"os"
)

// 保持和服务器连接

// ShowMenu 显示登录成功的菜单
func ShowMenu() {
	fmt.Println("---------------xxx登录成功---------------")
	fmt.Println("\t\t1.显示在线用户列表")
	fmt.Println("\t\t2.发送消息")
	fmt.Println("\t\t3.信息列表")
	fmt.Println("\t\t4.退出")
	fmt.Print("请选择(1-4) >")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("查看信息列表")
	case 4:
		{
			fmt.Println("退出了系统")
			// 最好是要通知服务器要关闭连接了
			os.Exit(0)
		}
	default:
		fmt.Println("非法输入")

	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	} // 创建一个Transfer实例 不停向服务器发送消息
	for { // 保持通讯
		fmt.Println("客户端正在等服务器发送的消息...")
		mes, err := tf.ReadPkg() // 这里的mes是已经是序列化后的
		if err != nil {
			fmt.Println("服务器端 tf.ReadPkg() err:", err)
			return
		}
		// 如果读取到消息 又是下一步处理逻辑
		fmt.Printf("mes = %v\n", mes)

	}
}
