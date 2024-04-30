package main

import (
	"fmt"
	"time"
)

func main() {
	var intChan = make(chan int, 10)
	var exitChan = make(chan bool, 1)

	go write66(intChan)
	go read66(intChan, exitChan)

	<-exitChan

}

func write66(intChan chan int) {
	for i := 0; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入数据 =", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(intChan)
}

func read66(intChan chan int, endChan chan bool) {
	for {
		i, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("读到数据 = %v\n", i)
		time.Sleep(time.Second)
	}
	endChan <- true
	close(endChan)
}
