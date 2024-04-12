package main

import (
	"fmt"
)

// 使用 map[string]map[string]sting 的map类型key:表示用户名，是唯一的，不可以重复
// 如果某个用户名存在，就将其密码修改"888888"，
// 如果不存在就增加这个用户信息,(包括昵称nickname 和 密码pwd)
// 编写一个函数 modifyUser(users map[string]map[string]sting, name string)完成上述功能
func main() {
	var usersList map[string]map[string]string
	usersList = make(map[string]map[string]string, 1)

	usersList["smith"] = make(map[string]string, 1)
	usersList["smith"]["nickname"] = "小花貓"
	usersList["smith"]["password"] = "pass"

	modifyUser(usersList, "shy")
	modifyUser(usersList, "smith")

	fmt.Println(usersList)

}

func modifyUser(users map[string]map[string]string, name string) {
	//v, ok := users[name]
	if users[name] != nil { // 如果某个用户名存在 就将其密码修改"888888"
		users[name]["password"] = "88888888"
	} else { // 如果不存在就增加这个用户信息,(包括昵称nickname 和 密码pwd)
		users[name] = make(map[string]string, 2)
		var nickName, pwd string
		fmt.Printf("設置暱稱：")
		fmt.Scan(&nickName)
		users[name]["nickname"] = nickName
		fmt.Printf("設置密碼：")
		fmt.Scan(&pwd)
		users[name]["password"] = pwd
	}
}
