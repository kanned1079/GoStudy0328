package utils

import (
	"database/sql"
	"log"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	// 格式为 数据库用户名:数据库密码@[tcp(数据库地址:端口)]/数据库名
	// Open函数只是验证其参数 而不创建与数据库的连接 如果要检查数据源的名称是否合法 应该调用返回值Ping方法
	Db, err = sql.Open("mysql", "go:PassCode987!@[tcp(api.ikanned.com:4407)]/db0")
	if err != nil {
		log.Println("open db error:", err)
		return
	}
	log.Println("open db success")
}

//func (*sql.DB) Ping() error {
//
//}
