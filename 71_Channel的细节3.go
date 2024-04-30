package main

import (
	"fmt"
	"time"
)

// goroutine中使用recover，解决协程中出现panic，导致程序崩渍问题.【案例演示】

func sayHe71() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}
}

func test71() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	var strMap map[int]string
	strMap[0] = "helloWorld" // 这里报错了会导致程序奔溃 panic: assignment to entry in nil map
	// 这里为了防止出现panic 可以用 defer和recover
}

func main() {
	go sayHe71()
	go test71()

	for i := 0; i < 10; i++ {
		fmt.Println("main")
		time.Sleep(time.Second)
	}
}
