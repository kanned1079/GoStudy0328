package main

import "fmt"

var myMap1 = make(map[int]int, 10)

func jiecheng(n int) {
	var sum int = 1
	for i := 0; i <= n; i++ {
		sum *= i
	}
	myMap1[n] = sum
}

func main() {
	for i := 1; i <= 200; i++ {
		go jiecheng(i)
	}
	for index, value := range myMap1 {
		fmt.Println(index, value)
	}
}
