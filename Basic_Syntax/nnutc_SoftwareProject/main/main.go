package main

import (
	"fmt"
)

func main() {
	fmt.Println("阶乘(10) = ", factorial(10))
	fmt.Println("fibracci(10) = ", fibonacci(10))
	fmt.Println("十进制26转换十六进制 = ", decimalToBaseN(26, 16))

	originalString := "abcdefg"
	reversedString := reverseString(originalString)

	fmt.Printf("原始字符串：%s\n", originalString)
	fmt.Printf("逆序字符串：%s\n", reversedString)

	var num1 int = 123
	fmt.Println("123倒序数 = ", reverse_int(num1))
	//reverseNum(123)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func decimalToBaseN(decimal, base int) string {
	if decimal == 0 {
		return "0"
	}
	digits := "0123456789ABCDEF"
	result := ""
	for decimal > 0 {
		remainder := decimal % base
		result = string(digits[remainder]) + result
		decimal /= base
	}
	return result
}

func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return string(s[len(s)-1]) + reverseString(s[0:len(s)-1])
}

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}
