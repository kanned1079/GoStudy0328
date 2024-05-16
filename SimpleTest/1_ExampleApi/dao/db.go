package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	dbUser  = "go"
	dbPass  = "PassCode987!"
	dbProto = "tcp"
	dbHost  = "api.ikanned.com"
	dbPort  = "4407"
	dbName  = "db0"
)

var (
	Db  *gorm.DB
	err error
)

func InitDB() {
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbProto, dbHost, dbPort, dbName),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		log.Println("MySQL数据库连接失败 err: ", err)
	} else {
		log.Println("MySQL数据库连接成功")
	}
}
