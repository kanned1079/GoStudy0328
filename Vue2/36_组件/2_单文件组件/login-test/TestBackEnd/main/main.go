package main

import (
	"GoStudy0328/Vue2/36_组件/2_单文件组件/login-test/TestBackEnd/dao"
	"GoStudy0328/Vue2/36_组件/2_单文件组件/login-test/TestBackEnd/users"
)

func main() {
	testUser1 := users.UniUsers{}
	dao.Db.Model(&users.UniUsers{}).Create(&testUser1)
}
