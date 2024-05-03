package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 集合Set 就是一个基于HashTable数据结构的 里面存放*无序的*字符串
// 并且元素的值是*不能重复的*

// 比如 存放多个电子邮件
// key -> email
// value aaa@bbb.com bbb@ccc.com ccc@ddd.com ...
// Redis指令 sadd emal aaa@bbb.com bbb@ccc.com ccc@ddd.com

func main() {
	var err error
	conn, err := redis.Dial("tcp", "api.ikanned.com:16379")
	if err != nil {
		panic(err)
	}
	fmt.Println("redis connect success")
	defer conn.Close()

	// 先清空
	_, err = conn.Do("Select", 1)
	_, err = conn.Do("FlushAll")

	// 添加数据
	_, err = conn.Do("sAdd", "emails", "aaa@bbb.com", "bbb@ccc.com", "ccc@ddd.com")

	// 取出所有元素
	allElems := make([]string, 0)
	allElems, err = redis.Strings(conn.Do("sMembers", "emails"))
	if err != nil {
		panic(err)
	}
	for _, elem := range allElems {
		fmt.Println(elem)
	}

	// 判断是否是成员
	isMem, err := redis.Bool(conn.Do("sIsMember", "emails", "aaa@bbb.com"))
	fmt.Println(isMem) // 如果是成员返回true 否则返回false

	// 删除指定值
	isRem, err := redis.Bool(conn.Do("sRem", "emails", "bbb@ccc.com"))
	fmt.Println(isRem) // 如果删除成功了返回true

}
