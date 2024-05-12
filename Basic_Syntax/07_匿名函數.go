package main

import "fmt"

func main() {

	// 匿名函數
	resu := func(num1, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println("resu = ", resu)

	a := func(num1, num2 int) int { // 這裏的a就已經是一個函數變量了
		return num1 + num2
	}
	fmt.Println("sum = ", a(10, 30)) // 可以直接這樣調用

}
