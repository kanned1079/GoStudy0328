package main

import "fmt"

func main() {
	// Channel可以声明为只读或只写
	// 默认情况下Channel是可读可写的

	//var chan1 chan int // 这是可读可写
	var chan2 chan<- int //这是只写
	chan2 = make(chan<- int, 3)
	chan2 <- 10
	//fmt.Printf("chan2 = ", <-chan2) // 这是不能读取的

	var chan3 <-chan int
	chan3 = make(<-chan int, 3)
	//chan3 <- 30 // 这是不能写入的
	fmt.Println("chan3 = ", chan3)
}

func send69(chan<- chan int) {

}

func receive69(<-chan chan int) {

}
