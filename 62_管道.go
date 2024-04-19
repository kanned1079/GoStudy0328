package main

import (
	"fmt"
	"sync"
	"time"
)

// 现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
// 最后显示出来。要求使用goroutine完成

// 那么不同的goroutine怎么互相通信

// 编写一个函数 计算各个数的阶乘 并放入map中
// 启动多个协程 统计的将结果放入到map中
// map应该使用一个全局
var (
	myMap = make(map[int]int, 10)
	// 可以给全局变量加互斥锁
	// 声明一个全局的互斥锁
	lock sync.Mutex // 这个是全局的互斥锁
	// sync全称synchornized
	// sync包提供了基本的同步基元，如互斥锁。
	// 除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。

)

// 前面使用全局变量加锁同步来解决goroutine的通讯 但是不完美
// 1)主线程在等待所有goroutine全部完成的时间很难确定我们这里设置10秒，仅仅是估算。
// 2)如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有goroutine处于工作状态，这时也会随主线程的退出而销毁
// 3)通过全局变量加锁同步来实现通讯，也并不利用多个协程对全局变量的读写操作。
// 4)上面种种分析都在呼唤一个新的通讯机制-channel

// test62() 这个用来计算n的阶乘 然后放入map中
func test62(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 这里将res这个结果放入到map中
	lock.Lock() // 加锁
	myMap[n] = res
	lock.Unlock() // 解锁
}

func main() {
	// 在这里开启多个协程
	for i := 1; i <= 20; i++ {
		go test62(i)
		// 但是这里的200个协程在同时操作一个Map地址
	}

	// 但是这里等待的时间还是会不确定
	time.Sleep(time.Second * 10)
	// 仍然报错 fatal error: concurrent map writes

	// 遍历结果
	// 主线程完成后会关闭其他的协程 所以这里会看不到任何结果
	// go build -race 62_管道.go

	// 如果有在访问map 需要上锁 Lock
	// 其他的协程需要排队等待 直到map被Unlock
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}
