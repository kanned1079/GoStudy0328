package main

import (
	"fmt"
)

// channle本质就是数据是先进先出
// 1)Channel的本质上就是一个数据结构 -> 队列
// 2)数据是先进先出 [FIFO]
// 3)线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
// 4)channel是有类型的，一个string的channel只能存放string类型数据

// Channel的声明/定义
// channel是引用类型
// channel必须初始化后才能使用
// channel是有类型的

func main() {
	// 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan = %v, &intChan = %v\n", intChan, &intChan)
	// intChan = 0x140000b8000, &intChan = 0x140000aa018

	// 向管道写入数据
	intChan <- 10
	numb := 20
	intChan <- numb

	intChan <- 30
	num2 := <-intChan
	fmt.Println("num2 = ", num2) // 先进先出 那么这里出来的就是10
	intChan <- 40
	// 这里会报错 fatal error: all goroutines are asleep - deadlock!
	// 但是只要取了一个就又可以装了

	// 从Channel中读取数据
	// 输出intChan的length和cap
	fmt.Println("len(intChan) = ", len(intChan))
	fmt.Println("cap(intChan) = ", cap(intChan))
	// 在没有使用协程的情况下 如果Channel的数据已经全部取出 再取就会报告deadlock

	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan) // 到这里都是可以取到的
	//fmt.Println(<-intChan) // 到这里会报错因为已经取完了数据

	//test63_1()
	test63_2()

}

func test63_1() {
	// 创建一个Channel 最多存放3个int 演示存入和读取
	var intChannel1 chan int
	intChannel1 = make(chan int, 3)
	for i := 1; i <= cap(intChannel1); i++ {
		intChannel1 <- i * 10
		fmt.Println("data = ", i*10)
	}
	fmt.Println("len(intChannel1) = ", len(intChannel1))
	fmt.Println("cap(intChannel1) = ", cap(intChannel1))
	for i := 1; i <= cap(intChannel1); i++ {
		fmt.Printf("intChannel1[%d] = %d\n", i, <-intChannel1)
	}
}

func test63_2() {
	fmt.Println("-----------------------------")
	// 创建一个mapChan 最多可以存放10个map[string]string的kyy-value
	var mapChan1 chan map[string]string
	mapChan1 = make(chan map[string]string, 10)
	m1 := make(map[string]string, 1)
	for i := 1; i <= cap(mapChan1); i++ {
		m1[fmt.Sprintf("%s", "city")] = "上海"
		mapChan1 <- m1
	}
	fmt.Println(len(m1))
	fmt.Println("len(mapChan1) = ", len(mapChan1))
	fmt.Println("cap(mapChan1) = ", cap(mapChan1))
	for i := 1; i <= cap(mapChan1); i++ {
		fmt.Printf("mapChan1[%02d] = %v\n", i, <-mapChan1)
	}

}
