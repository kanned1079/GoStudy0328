package main

import (
	"GoStudy0328/src/p45_HomeAccount/s2_ObjectOriented/myUtils"
	"fmt"
)

// 將記帳軟件的所有功能 封裝到一個結構體中 然後調用該結構體的方法

// 定義一個結構體
//type FamilyAccount struct {
//	Balance float64
//	Money   float64
//}

func main() {
	// 在main方法中創建一個FamilyAccount實例
	var fam1 = myUtils.NewAccount(10000)
	// 创建用户
	var kanna = myUtils.NewAdmin("kanna", "123456")
	fmt.Print("输入用户名：")
	var usr, pwd string
	if fmt.Scanln(&usr); usr == kanna.GetName() {
		fmt.Print("输入密码：")
		if fmt.Scanln(&pwd); pwd == kanna.GetPwd() {
			fam1.MainMenu()
		} else {
			fmt.Println("密码错误")
		}
	} else {
		fmt.Println("用户名错误")
	}

}
