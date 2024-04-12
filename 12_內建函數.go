package main

import "fmt"

func main() {
	// new 用來分配內存 主要用來分配值類型

	num1 := 10
	fmt.Printf("num1's type=%T, num1=%v, num1's addr=%p\n", num1, num1, &num1)

	var num2 = new(int)
	fmt.Printf("num2's type=%T, num2=%v, num2's addr=%p, 指向的值=%v\n", num2, num2, &num2, *num2)

}
