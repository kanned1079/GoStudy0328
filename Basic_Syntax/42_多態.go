package main

import "fmt"

// 变量(实例)具有多种形态。面向对象的第三大特征
// 在Go语言，多态特征是通过接口实现的。
// 可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。

// 接口體現多態的兩種形式 (多態Poly)
// 1. 多態參數
// 2. 多態數組

type USB_42 interface {
	Start()
	Stop()
}

type Computer_42 struct {
}

func (pc *Computer_42) Working(usb USB_42) {
	usb.Start()
	usb.Stop()
	//usb.Call() // 這樣不行 因為Camera沒有Call()方法
	// 所以這裏要進行判斷這個usb是Phone的還是Camera的 使用到這個需要用到類型斷言

}

type Phone_42 struct {
	Name string
}

type Camera_42 struct {
	Name string
}

func (p Phone_42) Start() {
	fmt.Println("手機開始工作...")
}

func (p Phone_42) Stop() {
	fmt.Println("手機停止工作...")
}

func (c Camera_42) Start() {
	fmt.Println("相機開始工作...")
}

func (c Camera_42) Stop() {
	fmt.Println("相機停止工作...")
}

func main() {
	// 這裏是第一種方法
	//var phone1 Phone_42
	//var camera1 Camera_42
	//var pc1 Computer_42
	//pc1.Working(phone1)
	//pc1.Working(camera1)

	// 這裏是第二種方法
	// 多态数组
	// 给Usb数组中，存放 Phone 结构体 和 Camera结构体变量
	// Phone还有一个特有的方法请遍历Usb数组
	// 如果是Phone变量，除了调用Usb 接口声明的方法外，还需要调用Phone特有方法 call.

	// 定義一個USB接口數組 可以存放Phone和Camera結構體變量
	var usbArr [3]USB_42
	// 按理來說 數組裡面只能存放一種數據類型 但是這裏放了Phone又放了Camera
	// 這就是基於接口的多態來實現的
	usbArr[0] = Phone_42{Name: "SONY"}
	usbArr[1] = Phone_42{Name: "Xiaomi"}
	usbArr[2] = Camera_42{Name: "Nikon"}
	fmt.Println(usbArr)

}
