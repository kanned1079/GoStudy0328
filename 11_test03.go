package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now().Unix()
	test03()
	endTime := time.Now().Unix()
	fmt.Printf("執行test03()耗時為 = %vs\n", endTime-startTime)
}

func test03() {
	const s = ""
	str := s
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
}
