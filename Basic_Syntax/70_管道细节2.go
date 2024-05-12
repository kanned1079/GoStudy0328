package main

import (
	"fmt"
	"time"
)

// 使用select可以解决从管道取数据的阻塞问题
// goroutine中使用recover，解决协程中出现panic，导致程序崩渍问题.【案例演示】
// 说明:如果我们起了一个协程，但是这个协程出现了panic，如果我们没有捕获这个panic,就会造成整个程序崩溃
// 这时我们可以在goroutine中使用recover来捕获panic,进行处理，这样即使这个协程发生的问题，但是主线程仍然不受影响，可以继续执行。

func main() {
	var numChan = make(chan int, 10)
	for i := 0; i < 10; i++ {
		numChan <- i
	}
	var strChan = make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- fmt.Sprintf("hello%d", i)
	}

	// 传统方法在遍历管道时 如果不关闭会导致阻塞而deadlock
	// 实际开发中 不知道什么时候关闭管道

	func() {
		for {
			select {
			case v := <-numChan:
				fmt.Println("读取到数据 = ", v)
				time.Sleep(time.Second)
			case v := <-strChan:
				{
					fmt.Println("取到数据 = ", v)
					time.Sleep(time.Second)

				}
			default:
				fmt.Println("err")
				time.Sleep(time.Second)
				return // 也可以使用 break lable1
			}
		}
	}()

}
