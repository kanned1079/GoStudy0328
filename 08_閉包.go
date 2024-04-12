package main

import (
	"fmt"
	"strings"
)

// AddUpper 累加器
// AddUpper 是一個函數 它的返回值是 func(int)int
func AddUpper01() func(int) int {
	var n int = 10
	return func(i int) int {
		n += i
		return n
	}
}

/*
	這一段就是一個閉包
	返回的是一个匿名函数，--但是这个匿名函数引用到函数外的n--
	因此这个匿名函数就和n形成个整体，构成闭包。
		閉包是一個類 函數就是操作 n就是字段 函數和它使用到的n構成了一個閉包

	var n int = 10
	return func(i int) int {
		n += i
		return n
	}

	當反覆調用函數f時候 因為n是累加的 （n只會初始化一次） 每次就會進行累加
	搞清楚閉包的關鍵是 分析清除返回的函數使用到了哪些變量 因為函數和它引用到的構成閉包
*/

func AddUpper02() func(int) int {
	var n int = 10
	var str string = "hello "
	return func(i int) int {
		n += i
		str += "A"
		fmt.Println("str = ", str)
		return n
	}
}

// strings.HasSuffix()函數可以用來判斷一個文件又沒有指定的後綴
// 傳入的suffix是後綴
// 雖然使用傳統的方法也可以實現 但是使用閉包後可以使用上次傳傳入的值 做到傳入一次就可以反覆使用
func makeSuffix(suffix string) func(string) string {
	// 在這裏的suffix字段和下邊的匿名函數 構成了一個閉包
	return func(name string) string {
		// 如果name沒有指定後綴就加上 否則就返回原來的name
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 不使用閉包 這裏就需要傳入兩個變量
func makeSuffix2(name, suffix string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main() {
	f := AddUpper01()
	fmt.Println(f(1))
	fmt.Println(f(2)) // 上邊運行結束後返回的已經是11了 所以 11+2 = 13

	f2 := AddUpper02()
	fmt.Println(f2(1))
	fmt.Println(f2(2)) // 上邊運行結束後返回的已經是11了 所以 11+2 = 13
	fmt.Println(f2(3))

	// 使用閉包
	f3 := makeSuffix(".jpg")
	fmt.Println("文件處理後的文件名是：", f3("winter"))   // 這裏給加上了後綴
	fmt.Println("文件處理後的文件名是：", f3("bird.jpg")) // 這裏已經有了後綴

	// 不使用閉包
	fmt.Println("文件處理後的文件名是：", makeSuffix2("dog", ".jpg"))
	fmt.Println("文件處理後的文件名是：", makeSuffix2("cat.jpg", ".jpg"))

}
