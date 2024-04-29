package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 1.新建连接
	conn, err := redis.Dial("tcp", "some-redis.orb.local:6379")
	if err != nil {
		fmt.Println("连接到Redis失败", err)
		return
	} else {
		fmt.Println("连接到Redis成功")
	}
	defer conn.Close()

	// 2.向Redis写入key-value数据
	_, err = conn.Do("SET", "key1", "value1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("写入数据成功")
	}

	// 2.向Redis写入key-value数据
	//r, err := conn.Do("GET", "key1")
	r, err := redis.String(conn.Do("GET", "key1"))
	// 这里返回的r是空接口类型
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("读取数据成功")
	}
	// 因此需要进行转换
	//r1 := reflect.ValueOf(r).
	// 不需要自己进行转换 Redis提供了很多方法
	fmt.Println(r)
}
