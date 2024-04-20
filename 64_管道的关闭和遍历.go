package main

import (
	"fmt"
	"time"
)

// 请完成goroutine和chanl同工作的案例，具体要求:
// 1)开启一个writeData协程向管道intchan中写入50个整数
// 2开启一个readData协程从管道intChan中读取writeData写入的数据
// 3)注意: writeData和readDate操作的是同一个管道
// 4)主线程需要等待writeData和readDate协程都完成工作才能退出

func testCloseChannel() {
	// Channel的关闭
	// 当Channel关闭后 里面的数据只能读取不能写入
	var intChan1 chan int
	intChan1 = make(chan int, 100)
	intChan1 <- 10
	intChan1 <- 20
	close(intChan1) // 这之后 panic: send on closed channel
	//for i := 0; i < 100-2; i++ {
	//	intChan1 <- i
	//}
	fmt.Println(<-intChan1)
	fmt.Println("len(intChan1) =", len(intChan1))
	fmt.Println("cap(intChan1) =", cap(intChan1))
}

func traverseChan() {
	fmt.Println("----------------------------")
	var intChan2 chan int
	intChan2 = make(chan int, 100)
	//close(intChan1)
	for i := 0; i < 100; i++ {
		intChan2 <- i
	}
	//fmt.Println(<-intChan1)
	fmt.Println("len(intChan1) =", len(intChan2))
	fmt.Println("cap(intChan1) =", cap(intChan2))

	// * 在遍历的时候如果Channel没有关闭 就会出现deadlock错误
	// * 如果Channel已经关闭 则会正常遍历数据 遍历完后 就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v = ", v)
	}
}

// 1)开启一个writeData协程向管道intchan中写入50个整数
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("写入的值 =", i)
		time.Sleep(time.Millisecond)
	}
	close(intChan)
}

// 2开启一个readData协程从管道intChan中读取writeData写入的数据
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan // v是值，ok是是否读取成功
		if !ok {
			break
		}
		fmt.Println("data = ", v)
		time.Sleep(time.Millisecond)
	}
	// 读取完数据后 即任务完成
	exitChan <- true
	close(exitChan)
}

// 3)注意: writeData和readDate操作的是同一个管道
// 4)主线程需要等待writeData和readDate协程都完成工作才能退出

func maintThread() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	// 用这个不行
	//time.Sleep(time.Second * 5)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}

// 启动一个协程，将 1-2000的数放入到一个channel中 比如 numChan
// 启动8个协程，从numChan取出数(比如n)
// 并计算 1+..*n的值，并存放到resChan最后8个协程协同完成工作后
// 再遍历reschan,显示结果 如 res[1]=1, res[10]=55..
// 注意:考虑 resChan chan int是否合适

func fib64(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func calcu(i int, numChan, resChan chan int) {
	for num := range numChan {
		fact := fib64(num)
		resChan <- fact
		fmt.Printf("calcu[%d], num = %v, fact = %v\n", i, num, fact)
		time.Sleep(time.Millisecond)
	}
	//close(resChan)
}

func ptrMain() {

	var numChan chan int = make(chan int, 200)
	var resChan chan int = make(chan int, 200)
	go func() {
		for i := 1; i <= 200; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for i := 1; i <= 8; i++ {
		go calcu(i, numChan, resChan)
	}

	go func() {
		for i := 1; i <= 8; i++ {
			<-resChan
		}
		close(resChan)
	}()

	results := make(map[int]int)
	for num := range resChan {
		results[num] = num
	}
	for i := 0; i < len(results); i++ {
		fmt.Println("results =", results)
	}

}

func ptrMain2() {
	numChan := make(chan int, 200)
	resChan := make(chan int, 200)

	go func() {
		for i := 1; i <= 200; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for i := 1; i <= 8; i++ {
		go calcu(i, numChan, resChan)
	}

	// Wait for all calcu goroutines to finish
	go func() {
		for i := 1; i <= 8; i++ {
			<-resChan
		}
		close(resChan)
	}()

	results := make(map[int]int)
	for num := range resChan {
		results[num] = num
	}

	for num := range results {
		fmt.Println("result =", num)
	}
}

func main() {
	//testCloseChannel()
	//traverseChan()
	//defer maintThread()
	//maintThread()
	ptrMain2()
}
