package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// Redis链接池
// 通过Golang 对Redis操作 还可以通过Redis链接池
// 流程如下
// 1)事先初始化一定数量的链接 放入到链接池公
// 2)当Go需要操作Redis时 直接从Redis链接池取出链接即可
// 3)这样可以节省临时获取Redis链接的时间 从而提高效率

// 使用set
// 存放商品信息
// 商品名
// 价格
// 生产日期

// 定义一个全局的Pool
var pool *redis.Pool

// 这个init会在main执行前被执行
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "api.ikanned.com:16379")
		},
	}
}

func main() {
	//var err error
	// 连接池核心代码
	/*
		var pool *redis.Pool
		pool = &redis.Pool{ // 这个Pool是一个结构体
			MaxIdle:     8, // 最大空闲连接数
			MaxActive:   0, // 和数据库的最大连接数 0表示不限制
			IdleTimeout: 300, // 最大空闲时间 有一个连接在100s中没有用过就会释放连接
			Dial: func() (redis.Conn, error) { // 初始化连接池
				return redis.Dial("tcp", "api.ikanned.com:16379") // 指定连接哪个服务器
			},
		}
		//c := pool.Get() // 从连接池中获取一个连接
		pool.Close() // 一旦关闭连接池后就不能再从连接池中取连接了
	*/

	// 先从Pool中取出一个连接 即使到达8个后会自动增长
	// 如果要从Pool中取出连接 一定要保证这个连接池是没有被关闭的
	conn1 := pool.Get()
	defer conn1.Close()
	// 存入
	_, err := conn1.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println(err)
	}
	// 取出
	r, err := redis.String(conn1.Do("Get", "name"))
	fmt.Println(r)

	pool.Close() // 如果在这里关闭连接池 那么下面就会报错 redigo: get on closed pool

	conn2 := pool.Get()
	defer conn2.Close()
	// 存入
	_, err = conn2.Do("Set", "name2", "tom2")
	if err != nil {
		fmt.Println(err)
	}
	// 取出
	r, err = redis.String(conn2.Do("Get", "name2"))
	fmt.Println(r)
}
