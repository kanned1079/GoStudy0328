package main

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"GoStudy0328/SimpleTest/1_ExampleApi/user"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {

	dao.InitDB()

	for i := 1; i < 1000; i++ {
		var testUser1 user.MyUser = user.MyUser{
			Name:        "TestUser" + strconv.Itoa(i),
			Gender:      "男",
			Age:         16,
			PhoneNumber: strconv.Itoa(15344839 + i),
			Email:       fmt.Sprintf("TestUser%d@ikanned.com", i),
			Password:    generateRandomString(26),
			Premium:     "lv-0",
			UpdatedAt:   time.Now(),
		}
		testUser1.Insert()
	}

	//testUser1.Insert()

}

func generateRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
