package main

import "fmt"

func main() {
	fmt.Println("function main...")
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
}

var name string
var age int

// 這個函數會會在定義的函數前被main函數調用
func init() {
	fmt.Println("function init...")

	// 可以用來做初始化
	age = 100
	name = "Tom"
}
