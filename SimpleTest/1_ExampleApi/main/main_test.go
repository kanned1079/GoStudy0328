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
		Name:        "崔俊2",
		Gender:      "女",
		Age:         16,
		PhoneNumber: "153453460839",
		Email:       "cj2@ikanned.com",
		Password:    "cjpasfergfesrs",
		Premium:     "lv-7",
		UpdatedAt:   time.Now(),
	}
	testUser1.Insert()

	//testUser1.Insert()

}
