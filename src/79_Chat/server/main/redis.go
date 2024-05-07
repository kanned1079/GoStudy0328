package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeOut time.Duration) { // 这个是不会自动调用的
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大空闲时间
		MaxActive:   maxActive,   //最大连接数
		IdleTimeout: idleTimeOut, // 等待超时时间
		Dial: func() (redis.Conn, error) { // 初始化链接Redis
			return redis.Dial("tcp", address)
		},
	}
}
