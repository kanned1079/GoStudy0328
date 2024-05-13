package main

import (
	"GoStudy0328/GoWeb/Database_Operation/src/model"
	"fmt"
	"log"
	"testing"
)

//func TestMain(m *testing.M) {
//	log.Println("Starting TestMain")
//	m.Run()
//}

func TestGetUserById(t *testing.T) {
	log.Println("查询单行记录")
	user := &model.User{}
	user, err := model.GetUserById(1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("查询到: ", user)

	log.Println("测试查询全部记录")
	t.Run("GetUsersById", testGetUsersById)
}

func testGetUsersById(t *testing.T) {
	log.Println("查询全部记录")
	usersSlice, err := model.GetUsersById()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("查询到: ")
	for i := range usersSlice {
		fmt.Printf("%v %v %v %v\n", usersSlice[i].Id, usersSlice[i].UserName, usersSlice[i].Password, usersSlice[i].Email)
	}
}
