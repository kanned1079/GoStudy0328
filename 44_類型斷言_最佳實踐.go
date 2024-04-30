package main

import (
	"fmt"
)

type USB_44 interface {
	Start()
	Stop()
	//Call()
}

type Computer_44 struct {
}

func (pc *Computer_44) Working(usb USB_44) {
	usb.Start()
	usb.Stop()
	// 在這裏使用類型斷言 判斷是不是來自Phone的
	if phone1, ok := usb.(Phone_44); ok {
		phone1.Call()
	} else {
		fmt.Println("類型不對")
	}
}

type Phone_44 struct {
	Name string
}

func (p Phone_44) Start() {
	fmt.Println("手機開始工作")
}

func (p Phone_44) Stop() {
	fmt.Println("手機停止工作")
}

func (p Phone_44) Call() {
	fmt.Println("手機在calling")
}

type Camera_44 struct {
	Name string
}

func (c Camera_44) Start() {
	fmt.Println("相機開始工作")
}

func (c Camera_44) Stop() {
	fmt.Println("相機停止工作")
}

// 類型斷言的最佳實踐2
func TypeJudge_44(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Paramater[%v] is bool, val = %v\n", i, x)
		case int, int64:
			fmt.Printf("Paramater[%v] is int, val = %v\n", i, x)
		case float32, float64:
			fmt.Printf("Paramater[%v] is float, val = %v\n", i, x)
		case string:
			fmt.Printf("Paramater[%v] is string, val = %v\n", i, x)
		case byte:
			fmt.Printf("Paramater[%v] is byte, val = %v\n", i, x)
		case Student_44:
			fmt.Printf("Paramater[%v] is Student, val = %v\n", i, x)

		case *Student_44:
			fmt.Printf("Paramater[%v] is *Student, val = %v\n", i, x)

		}
	}
}

type Student_44 struct {
}

// 那么怎么判断Student和*Student类型

func main() {
	var usbArr [3]USB_44
	usbArr[0] = Phone_44{Name: "XIAOMI"}
	usbArr[1] = Phone_44{Name: "SONY"}
	usbArr[2] = Camera_44{Name: "Nikon"}

	// 遍歷usbArr數組
	// Phone還有一個特有的方法Call() 遍歷usbArr數組
	// 如果是Phone變量 除了調用接口聲明的方法外 還需要調用特有的Call()方法
	var pc1 Computer_44
	for index, val := range usbArr {
		fmt.Printf("index = %d\n", index)
		pc1.Working(val)
	}

	var n1 float64 = 1.1
	var n2 int = 35
	var str string = "dwef"
	st1 := Student_44{}
	st2 := &Student_44{}
	TypeJudge_44(n1, n2, str, st1, st2)

}
