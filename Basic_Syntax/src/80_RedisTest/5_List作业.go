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
		_, err = conn.Do("rPush", "classics", classics[i])
		if err != nil {
			fmt.Println("rPush error", err)
			panic(err)
		}
	}

	read6(conn)

	popL6(conn) // 第一次pop出 西游记
	popL6(conn) // 第二次pop出 红楼梦
	popR6(conn) // 这是从右侧pop出 三国演绎
	//popR6(conn) // 再pop一次这个List就消失了

}

func read6(conn redis.Conn) {
	var err error
	fmt.Println("这是range")
	data, err := redis.Strings(conn.Do("lRange", "classics", 0, -1))
	if err != nil {
		fmt.Println("lRange error", err)
		panic(err)
	}
	for i := range data {
		fmt.Println(i, data[i])
	}
}

func popL6(conn redis.Conn) {
	var err error
	data, err := redis.String(conn.Do("lPop", "classics"))
	// 注意这里不要用Strings 因为pop是一个个推出 得到的是一个而不是一组
	if err != nil {
		fmt.Println("lPop error", err)
		panic(err)
	}
	fmt.Println("左侧pop = ", data)
}

func popR6(conn redis.Conn) {
	var err error
	data, err := redis.String(conn.Do("rPop", "classics"))
	// 如上 为什么要用string
	if err != nil {
		fmt.Println("lPop error", err)
		panic(err)
	}
	fmt.Println("右侧pop = ", data)
}
