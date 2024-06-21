package main

import (
	"GoStudy0328/SimpleJSONTest/dao"
	"GoStudy0328/SimpleJSONTest/encrypt"
	"GoStudy0328/SimpleJSONTest/users"
	_ "crypto/rand"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func randPhoneNumber(n int) (res string) {
	for _ = range n {
		res += strconv.Itoa(rand.Intn(8))
	}
	return
}

func randPassword(length int) (res string) {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	password := make([]byte, length)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(password)
}

func TestAddMultiUser(t *testing.T) {
	var ency encrypt.Encryptor
	for i := 1; i <= 100; i++ {
		var user = users.WxUsers{
			Email:    fmt.Sprintf("test0mail%d@ikanned.com", i),
			Password: randPassword(20),
			Name:     "test0user" + strconv.Itoa(i),
			Age:      1,
			Phone:    randPhoneNumber(8),
			Level:    0,
		}
		//dao.Db.Model(&users.WxUsers{}).Where("email = ?", user.Email).Delete(&user)
		user.Password = ency.Base64Encrypt(user.Password)
		dao.Db.Create(&user)

	}

}
