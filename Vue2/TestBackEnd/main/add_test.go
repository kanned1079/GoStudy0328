package main

import (
	"GoStudy0328/Vue2/TestBackEnd/dao"
	"GoStudy0328/Vue2/TestBackEnd/users"
	"math/rand"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {
	dao.Db.Model(&users.UniUsers{}).Create(&users.UniUsers{
		Name:  "张三",
		Age:   20,
		Sex:   1,
		Birth: time.Date(1997, time.July, getRand(30), getRand(24), 0, 0, 0, time.UTC),
		Addr:  "Jiangsu_NanjingCity",
	})
}

func getRand(num int) (age int) {
	return rand.Intn(num)
}
