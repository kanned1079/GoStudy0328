package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Db  *gorm.DB
	err error
)

type Users struct { // 这里的结构体名对应的表 也就是 users
	Id      int
	Name    string
	Age     int
	Email   string
	AddTime int
}

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "go:PassCode987!@tcp(api.ikanned.com:4407)/db0?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Database connection failed. err: ", err)
	} else {
		log.Println("Database connection success.")
	}
}

func (u *Users) TableName() string {
	return "users"
}
