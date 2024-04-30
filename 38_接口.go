package main

import "fmt"

// 接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。
// 接口体现了程序设计的 -多态- 和 -高内聚- -低偶合- 的思想。
// Golang中的接口，不需要显式的实现。
// 只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口。
// 因此，Golang中没有implement这样的关键字

type USB interface {
	// interface类型可以定义一组方法
	// 某个自定义类型(比如结构体Phone)要使用的时候,在根据具体情况把这些方法写出来
	// 但是这些不需要实现。并且interface不能包含任何变量
	Start()
	Stop()
}

type Phone struct {
}

// 让Phone实现USB接口的方法
func (p Phone) Start() {
	fmt.Println("手机開始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

// 让Camera实现USB接口的方法
func (c Camera) Start() {
	fmt.Println("相机開始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

// 编写一个方法 接收一個USB接口類型的變量
// 這個方法實現一個USB方法
// 説顯現了 Usb接口 就是指的是Usb接口中聲明的所有方法
func (pc Computer) Working(usb USB) {
	// 通過usb接口變量來實現Start()和Stop()方法
	usb.Start()
	usb.Stop()
}

func main() {
	pc1 := Computer{}
	phone1 := Phone{}
	camera1 := Camera{}

	pc1.Working(phone1) // 因為這裏傳入的是phone 所以實現的方法也就是phone的start和stop
	pc1.Working(camera1)

}
