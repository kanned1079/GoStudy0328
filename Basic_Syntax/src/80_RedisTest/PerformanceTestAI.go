package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	pool := newPool("api.ikanned.com:16379")
	defer pool.Close()

	conn := pool.Get()
	defer conn.Close()

	// 设置要测试的键值对
	key := "test_key"
	value := "test_value"

	// 设置要进行的测试次数
	numTests := 10000

	// 开始测试
	start := time.Now()
	for i := 0; i < numTests; i++ {
		// 设置键值对
		_, err := conn.Do("SET", key, value)
		if err != nil {
			fmt.Println("Error setting:", err)
			return
		}

		// 获取键值对
		_, err = redis.String(conn.Do("GET", key))
		if err != nil {
			fmt.Println("Error getting:", err)
			return
		}
	}
	elapsed := time.Since(start)

	fmt.Printf("Elapsed time for %d tests: %s\n", numTests, elapsed)
	fmt.Printf("Average time per test: %s\n", elapsed/time.Duration(numTests))
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}
