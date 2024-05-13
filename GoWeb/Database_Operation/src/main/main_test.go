package main

import (
	"GoStudy0328/GoWeb/Database_Operation/src/model"
	"log"
	"testing"
)

// TestMain 会在整个测试之前运行
func TestMain(m *testing.M) {
	log.Println("Starting TestMain")
	m.Run() //还需要运行m.Run()
}

func TestAddUser(t *testing.T) {
	log.Println("添加测试用户1")
	user := &model.User{
		UserName: "kingyyo",
		Password: "123456",
		Email:    "kingyyo@gmail.com",
	}
	user.AddUser()
	log.Println("调用子测试函数执行添加第二个用户")
	t.Run("第二个测试函数", testAddUser2)
}

func testAddUser2(t *testing.T) {
	log.Println("添加测试用户2")
	user := &model.User{
		UserName: "izumi kanna",
		Password: "Passcode1!",
		Email:    "kanna@kamiizumi.onmicrosoft.com",
	}
	user.AddUser2()
}

func testGetUserById(t *testing.T) {
	log.Println("查询单行记录")
	user := &model.User{}
	user, err := model.GetUserById(1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("查询到: ", user)
}
