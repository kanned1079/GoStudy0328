package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

func main() {
	conn, err := redis.Dial("tcp", "api.ikanned.com:16379")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	fmt.Println("SUCC")
	defer conn.Close()

	_, _ = conn.Do("FLUSHALL")

	_, _ = conn.Do("SELECT", "2")

	for i := 0; i < 2000; i++ {
		go testWrite()
	}
	time.Sleep(time.Second * 10)

}

func testWrite() {
	conn, err := redis.Dial("tcp", "api.ikanned.com:16379")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	fmt.Println("SUCC")
	defer conn.Close()
	for i := 0; i < 900000; i++ {
		_, err := conn.Do("SET", "cc"+strconv.Itoa(i), strconv.Itoa(i))
		if err != nil {
			fmt.Println("set failed:", err)
		}
		fmt.Println("Seted", i)
	}
}
