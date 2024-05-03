package main

import "github.com/gomodule/redigo/redis"

// 使用set
// 存放商品信息
// 商品名
// 价格
// 生产日期

func main() {
	//var err error
	// 连接池
	var pool *redis.Pool
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "api.ikanned.com:16379")
		},
	}
	//c := pool.Get()
	pool.Close()

}
