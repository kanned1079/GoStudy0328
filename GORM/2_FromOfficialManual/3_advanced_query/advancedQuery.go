package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var Db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "go:PassCode987!@tcp(api.ikanned.com:4407)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// parseTime=True&loc=Local 用于处理时间
	//Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 第一种写法
	// MySQL 驱动程序提供了 一些高级配置 可以在初始化过程中使用，例如：
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{})
	if err != nil {
		log.Println("Database connection failed. err: ", err)
	} else {
		log.Println("Database connection success.")
	}
}

type User struct {
	Id       int
	Name     string
	Age      int
	Birthday time.Time
}

// 假如这是需要的数据
type APIUser struct {
	ID   uint
	Name string
}

func specField() {
	// 在查询时，GORM 会自动选择 `id `, `name` 字段
	list1 := []APIUser{}
	Db.Order("id desc").Model(&User{}).Limit(100).Find(&list1)

	rangeList(list1)
	fmt.Println("-----------------------------")

}

func rangeList(users []APIUser) {
	for _, user := range users {
		fmt.Println(user)
	}
}

func main() {
	specField()
}
