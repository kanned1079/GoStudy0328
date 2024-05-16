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
		Name:        "kinggyo inari",
		Gender:      "女",
		Age:         18,
		PhoneNumber: "34860839",
		Email:       "kinggyo@ikanned.com",
		Premium:     "lv3",
		UpdatedAt:   time.Now(),
	}

	//testUser1.Insert()

}
