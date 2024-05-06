package main

import (
	"fmt"
	"os"
)

// 定义暂时的全局变量 一个id 一个密码
var userId int
var userPassword string

func main() {
	// 接受用户输入
	var key int
	// 判断是否还一直显示菜单
	var loop = true

	for loop {
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
				loop = false
			}
		case 2:
			{
				fmt.Println("注册用户")
				loop = false
			}
		case 3:
			{
				fmt.Println("退出系统")
				//loop = false
				os.Exit(0)
				fmt.Println(loop)
			}
		default:
			{
				fmt.Println("输入有误")
			}
		}

	}

	// 根据用户输入 显示新的菜单
	if key == 1 {
		// 说明用户要登录
		fmt.Print("请输入用户的ID > ")
		// 暂时使用上面的全局变量
		_, _ = fmt.Scanf("%d\n", &userId) // 记得这里的\n
		fmt.Print("输入密码 > ")
		_, _ = fmt.Scanf("%s\n", &userPassword)
		// 先把登录的函数写到另一个文件中 需要进行用户检验 设计Message

		if err := login(userId, userPassword); err != nil {
			fmt.Println("登录失败", err)
		} else {
			fmt.Println("登录成功")
		}

	} else if key == 2 {
		// 这里进行用户注册
	}

}
