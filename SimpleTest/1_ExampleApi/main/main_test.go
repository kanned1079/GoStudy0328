package main

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"GoStudy0328/SimpleTest/1_ExampleApi/user"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {

	dao.InitDB()
	var testUser1 user.MyUser = user.MyUser{
		Name:        "崔俊",
		Gender:      "男",
		Age:         17,
		PhoneNumber: "155460839",
		Email:       "cj@ikanned.com",
		Password:    "cjpass",
		Premium:     "lv-0",
		UpdatedAt:   time.Now(),
	}
	testUser1.Insert()

	//testUser1.Insert()

}
