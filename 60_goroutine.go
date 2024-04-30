package main

import (
	"fmt"
	"strconv"
	"time"
)

// 需求:要求统0000 的数字中，哪些是素数?
// 分析思路:
// 传统的方法，就是使用一个循环，循环的判断各个数是不是素数。
// 使用并发或者并行的方式，将统计素数的任务分配给多个goroutine去完成，这时就会使用到goroutine.

// Go协程的特点
//		有独立的棧空间
//		共享程序堆空间
//		调度由用户控制
//		协程是轻量级的线程

// Goroutine的调度模型
// MPG模式
// M: 操作系统的主线程 (是物理线程)
// P: 协程执行需要的上下文
// G: 协程

func test601() {
	for i := 0; i < 10; i++ {
		fmt.Println("test601() hello world " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 100)
	}
}

func test602() {
	for i := 0; i < 10; i++ {
		fmt.Println("test602() hello world " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go test601() // 开启了一个协程
	go test602()

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello golang " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 100)

	}
}
