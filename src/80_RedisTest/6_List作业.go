package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 将中国的自大名著 以List的形式存放到Redis
// 并完成crud操作

func main() {
	var err error
	conn, err := redis.Dial("tcp", "api.ikanned.com:16379")
	if err != nil {
		fmt.Println("redis dial error", err)
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("conn close error", err)
			panic(err)
		} else {
			fmt.Println("conn closed")
		}
	}()

	conn.Do("FLUSHALL")

	// 添加到List中
	var classics = [4]string{"西游记", "红楼梦", "水浒传", "三国演艺"}
	for i := 0; i < len(classics); i++ {
		content := classics[i]
		_, err = conn.Do("rPush", "classics", content)
		if err != nil {
			fmt.Println("rPush error", err)
			panic(err)
		}
	}

}
