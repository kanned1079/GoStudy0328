package main

import (
	"GoStudy0328/SimpleJSONTest/dao"
	"GoStudy0328/SimpleJSONTest/users"
	"testing"
)

func TestAddAUser(t *testing.T) {
	var u1 = users.WxUsers{
		Email:    "test1@gmail.com",
		Password: "123456",
		Name:     "test1",
		Age:      19,
		Phone:    "23455675",
		Level:    5,
	}
	//u1.CreateNew()
	dao.Db.Create(&u1)
}
