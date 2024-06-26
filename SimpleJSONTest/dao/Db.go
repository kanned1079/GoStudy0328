package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	dbUser  = "go"              // 使用哪个用户来连接数据库
	dbPass  = "PassCode987!"    // 用户的密码
	dbProto = "tcp"             // 连接的协议
	dbHost  = "api.ikanned.com" // 数据库主机的IP或域名
	dbPort  = "4407"            // 数据库连接端口
	dbName  = "db0"             // 操作哪一个数据库
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbProto, dbHost, dbPort, dbName),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Println("初始化数据库失败 err: ", err)
		panic(err)
	} else {
		log.Println("初始化数据库成功")
	}
}
