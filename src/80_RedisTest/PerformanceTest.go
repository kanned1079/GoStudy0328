package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

func main() {
	conn, err := redis.Dial("tcp", "192.168.0.243:16379")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	fmt.Println("SUCC")
	defer conn.Close()

	//_, _ = conn.Do("FLUSHALL")
	_, _ = conn.Do("SELECT", "2")

	var endChan chan bool = make(chan bool, 1000)

	for i := 0; i < 1000; i++ {
		go testWrite(i, endChan)
	}

	for i := 0; i < 1000; i++ {
		<-endChan
	}
	close(endChan)

}

func testWrite(n int, endChan chan bool) {
	conn, err := redis.Dial("tcp", "api.ikanned.com:16379")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	fmt.Println("SUCC")
	defer conn.Close()
	for i := 0; i < 10000; i++ {
		fmt.Printf("goRoutine %d \t", n)
		_, err := conn.Do("SET", "cc"+strconv.Itoa(n)+strconv.Itoa(i), strconv.Itoa(i))
		if err != nil {
			fmt.Println("set failed:", err)
		}
		fmt.Println("Seted", i)

	}
	endChan <- true
}
