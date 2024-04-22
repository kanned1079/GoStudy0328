package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	var intChan = make(chan int, 1000) // 这里只给1000
	var primeChan = make(chan int, 10000)
	var endChan = make(chan bool, 4) // 协程个数
	goroutineCount := 1

	go storeNums(intChan)

	start := time.Now()
	for i := 1; i <= goroutineCount; i++ {
		go judgePrime(i, intChan, primeChan, endChan)
	}

	//for i := 1; i <= goroutineCount; i++ {
	//	<-endChan
	//}
	// 那么 何时关闭管道
	// 只要能从endChan中取出4个true 那才算处理结束

	//for i := 1; i <= goroutineCount; i++ {
	//	<-endChan
	//}
	// 如果上面有四个true 那么就可以关闭endChan管道
	// 这里启动协程来做
	go func() {
		for i := 1; i <= goroutineCount; i++ {
			<-endChan
		}
		close(endChan)
		close(primeChan)
		end := time.Now()

		fmt.Println("花费时间 = ", end.Sub(start))
	}()
	//<-endChan
	//<-primeChan

	var primeSlice = make([]int, 1)
	for i := range primeChan {
		primeSlice = append(primeSlice, i)
	}
	sort.Ints(primeSlice)
	fmt.Println(primeSlice)

}

// <<<<<<< HEAD
func main1111() {
	traditionMethod()
}

// =======
// >>>>>>> origin/main
// 启动一个协程来存入数据1-8000
func storeNums(intChan chan int) {
	for i := 1; i <= 20000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 启动四个协程来判断素数
// 如果是素数 那么就放置到另一个管道中
func judgePrime(n int, intChan chan int, primeChan chan int, endChan chan bool) {
	for i := range intChan {
		if isPrime(i) {
			primeChan <- i
			//fmt.Println(i, "是素数 (协程", n, ")")
		}
	}
	endChan <- true
	fmt.Println("协程", n, "退出")
	//close(primeChan) // 这里还不能关闭管道
	//close(endChan)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//<<<<<<< HEAD

func traditionMethod() {
	start := time.Now()
	for i := 1; i <= 20000; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
	end := time.Now()
	fmt.Println("spend time =", end.Sub(start))

}

//=======
//>>>>>>> origin/main
