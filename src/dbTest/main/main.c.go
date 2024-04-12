package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "os/user"
	"strconv"
	_ "strconv"
	"time"
)

const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "password"
	dbName     = "go"
	dbHost     = "some-mysql.orb.local"
	//dbHost = "192.168.12.244"
	dbPort = "3306"
)

var db *sql.DB

type User struct {
	UserName string
	Password string
	Age      int
	ID_Card  string
	Address  string
}

func initDBConnection() (err error) {
	db, err = sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		return err
	}
	return nil
}

func insertUser(user User) error {
	stmt, err := db.Prepare("INSERT INTO users(username, password, age, id_card, address) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.UserName, user.Password, user.Age, user.ID_Card, user.Address)
	if err != nil {
		return err
	}
	fmt.Println("数据插入成功 ")
	return nil
}

func readUserInfos() ([]User, error) {
	var users []User = make([]User, 1)
	rows, err := db.Query("SELECT username, password, age, id_card, address FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserName, &user.Password, &user.Age, &user.ID_Card, &user.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func main() {
	//operateMysql()
	deleteAllRedisKeys()
	operateRedis()
}

func operateMysql() {
	startTime := time.Now()
	if err := initDBConnection(); err != nil {
		fmt.Printf("数据库初始化失败")
	}
	defer db.Close()

	var kanna User
	kanna.UserName = "kanna"
	kanna.Password = "Password1"
	kanna.Age = 18
	kanna.ID_Card = "3258568896354"
	kanna.Address = "常州钟楼区永红街道"

	for i := 0; i < 10000; i++ {
		fmt.Printf("第%d次操作：", i+1)
		_ = insertUser(kanna)
	}
	endTime := time.Now()
	spendTime := endTime.Sub(startTime)
	fmt.Println("Time = ", spendTime)

	//users, err := readUserInfos()
	//if err != nil {
	//	log.Fatal("读取用户列表出错：", err)
	//}
	//
	//fmt.Println("Users:")
	//for _, user := range users {
	//	fmt.Printf("用户名=%s, 密码=%s, 年龄=%d, 身份证号=%s\n", user.UserName, user.Password, user.Age, user.ID_Card)
	//}

}

func operateRedis() {
	startTime := time.Now()

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		//Addr:     "192.168.12.244:6379",
		Addr:     "some-redis.orb.local:6379",
		Password: "",
		DB:       0,
	})

	for i := 1; i <= 1; i++ {
		// 设置键值对
		err := rdb.Set(ctx, "key1"+strconv.Itoa(i), "内容"+strconv.Itoa(i), 0).Err()
		if err != nil {
			panic(err)
		}
		fmt.Printf("设置成功第%d个\n", i)
	}

	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	endTime := time.Now()
	spendTime := endTime.Sub(startTime)
	fmt.Println("Time = ", spendTime)
}

func deleteAllRedisKeys() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "some-redis.orb.local:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.FlushAll(ctx).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
