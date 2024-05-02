package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var ipaddr string = "api.ikanned.com"
var port string = "16379"

func main() {
	// 1.新建连接
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", ipaddr, port))
	if err != nil {
		fmt.Println("连接到Redis失败", err)
		return
	} else {
		fmt.Println("连接到Redis成功")
	}
	defer conn.Close()

	_, err = conn.Do("SELECT", "1")
	// 在这里选择了第二个数据库

	// 2.向Redis写入key-value数据
	_, err = conn.Do("HSET", "user1", "name", "kanna")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("写入数据成功")
	}
	_, err = conn.Do("HSET", "user1", "age", "18")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("写入数据成功")
	}

	// 2.向Redis写入key-value数据
	//r, err := conn.Do("GET", "key1")
	r, err := redis.String(conn.Do("HGET", "user1", "name"))
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

	e, err := redis.Int(conn.Do("HGET", "user1", "age"))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("读取数据成功")
	}

	fmt.Println(e)
	fmt.Printf("Type(e) = %T\n", e)

	// 上面都是一个个的放
	// 下面是一次放多个
	_, err = conn.Do("HMset", "name", "张三", "age", 40)
	// 同样的也可以多个读取
	rr, err := redis.Strings(conn.Do("Mget", "name", "age"))
	fmt.Printf("rr.Type() = %T\n", rr) // 这里的数据类型是 []string
	fmt.Println(rr)
	for i := range rr {
		fmt.Println(rr[i])
	}

	_, _ = conn.Do("SET", "UserName", "kinggyo")

	// 这里设置有效时间为60s
	_, _ = conn.Do("EXPIRE", "UserName", 60)
}
