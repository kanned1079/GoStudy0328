package main

import (
	"fmt"
	"time"
)

func main() {
	// 开启一个writeData协程 向管道中写入50个整数
	// 开启一个readData协程 读取数据
	// 主线程要等两个执行完后才能退出
	var intChan chan int = make(chan int, 50)
	var endChan chan bool = make(chan bool, 1)
	go writeData63(intChan)
	go readData63(intChan, endChan)
	//time.Sleep(time.Second)
	for v := range endChan {
		fmt.Println(v)
		fmt.Println("ended")
	}

}

func writeData63(intChan chan int) {
	for i := 0; i <= 50; i++ {
		intChan <- i
		fmt.Println("inserted =", i)
		time.Sleep(time.Millisecond)
	}
	close(intChan)
}

func readData63(intChan chan int, endChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("data =", v)
	}
	endChan <- true
	close(endChan)
}
