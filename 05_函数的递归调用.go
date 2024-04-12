package main

import "fmt"

// 函數在內部調用了自身

func main() {
	//test01(4)
	test02(4)
}

// 函數的遞歸從內向外執行
func test01(n int) {
	if n > 2 {
		n--
		test01(n)
	}
	fmt.Println("n = ", n)
}

func test02(n int) {
	if n > 2 {
		n--
		test01(n)
	} else {
		fmt.Println("n = ", n)

	}
}
