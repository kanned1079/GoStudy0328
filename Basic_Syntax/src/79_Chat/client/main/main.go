package main

import (
	"GoStudy0328/src/79_Chat/client/process"
	"fmt"
	"os"
)

// 定义暂时的全局变量 一个id 一个密码
var userId int
var userPassword string
var userName string

func main() {
	// 接受用户输入
	var key int
	// 判断是否还一直显示菜单
	//var loop = true

	for true {
		fmt.Println("---------------欢迎登录聊天室---------------")
		fmt.Println("\t\t1.登录聊天室")
		fmt.Println("\t\t2.注册用户")
		fmt.Println("\t\t3.退出系统")
		fmt.Print("请选择(1-3) > ")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			{
				fmt.Println("登录聊天室")
				fmt.Print("请输入用户的ID > ")
				_, _ = fmt.Scanf("%d\n", &userId) // 记得这里的\n
				fmt.Print("输入密码 > ")
				_, _ = fmt.Scanf("%s\n", &userPassword)
				//loop = false
				// 完成登录
				up := &process.UserProcess{}
				up.Login(userId, userPassword)
			}
		case 2:
			{
				fmt.Println("注册用户")
				//loop = false
				fmt.Print("输入用户的ID > ")
				_, _ = fmt.Scanf("%d\n", &userId)
				fmt.Print("输入用户的密码 > ")
				_, _ = fmt.Scanf("%s\n", &userPassword)
				fmt.Print("输入用户昵称 > ")
				_, _ = fmt.Scanf("%s\n", &userName)
				// 创建up的实例 来完成注册的请求
				up := &process.UserProcess{}
				_ = up.Register(userId, userPassword, userName)

			}
		case 3:
			{
				fmt.Println("退出系统")
				//loop = false
				os.Exit(0)
				//fmt.Println(loop)
			}
		default:
			{
				fmt.Println("输入有误")
			}
		}

	}

	//// 根据用户输入 显示新的菜单
	//if key == 1 {
	//	// 说明用户要登录
	//	//fmt.Print("请输入用户的ID > ")
	//	//// 暂时使用上面的全局变量
	//	//_, _ = fmt.Scanf("%d\n", &userId) // 记得这里的\n
	//	//fmt.Print("输入密码 > ")
	//	//_, _ = fmt.Scanf("%s\n", &userPassword)
	//	// 先把登录的函数写到另一个文件中 需要进行用户检验 设计Message
	//
	//	//if err := login(userId, userPassword); err != nil {
	//	//	fmt.Println("登录失败", err)
	//	//} else {
	//	//	fmt.Println("登录成功")
	//	//}
	//	//login(userId, userPassword) // 这里要改
	//	// 在新结构中这里要使用MCN
	//
	//} else if key == 2 {
	//	// 这里进行用户注册
	//}

}
