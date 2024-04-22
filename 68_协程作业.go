package main

import "fmt"

// 启动一个协程，将 1-2000的数放入到一个channel中 比如 numChan
// 启动8个协程，从numChan取出数(比如n)
// 并计算 1+..*n的值，并存放到resChan最后8个协程协同完成工作后
// 再遍历reschan,显示结果 如 res[1]=1, res[10]=55..
// 注意:考虑 resChan chan int是否合适

func storeNums68(numChan chan int) {
	for i := 1; i <= 200000; i++ {
		numChan <- i
	}
	close(numChan)
	fmt.Println("存储完成")
}

func calcu68(numChan chan int, resChan chan int, endChan chan bool) {
	for i := range numChan {
		resChan <- fib68(i)
	}
	endChan <- true
	// 这里还不能close
}

func fib68(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	var numChan = make(chan int, 200)
	var resChan = make(chan int, 2000)
	var endChan = make(chan bool, 8)
	workersNum := 256

	go storeNums68(numChan)

	for i := 1; i <= workersNum; i++ {
		go calcu68(numChan, resChan, endChan)
	}

	go func() {
		for i := 1; i <= workersNum; i++ {
			<-endChan
		}
		close(resChan)
		close(endChan)
	}()

	var mapFibNum = make(map[int]int, 1000)
	var num int = 1
	for i := range resChan {
		mapFibNum[num] = i
		num++
	}
	var lines int = 0
	for key, value := range mapFibNum {
		fmt.Printf("fib68(%d) = %d  ", key, value)
		lines++
		if lines%3 == 0 {
			fmt.Println()
		}
	}
}
