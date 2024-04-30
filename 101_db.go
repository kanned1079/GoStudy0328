package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

type Tag struct {
	uuid      string
	keySha256 string
}

const (
	dbDriver   = "mysql"
	dbUser     = "test"
	dbPassword = "password"
	dbName     = "gotestdb"
	dbHost     = "192.168.12.243"
	dbPort     = "3306"
)

func main() {
	// 打开数据库连接
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	//db, err := sql.Open("mysql", "test:password@tcp(192.168.12.243:3306)/gotestdb")
	if err != nil {
		fmt.Println("初始化数据库失败 err = ", err)
	} else {
		fmt.Println("数据库初始化成功")
	}

	// 程序结束时关闭数据库
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println("数据库关闭失败 err = ", err)
		}
	}()

	// 插入数据

}
