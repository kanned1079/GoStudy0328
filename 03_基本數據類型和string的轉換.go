package main

import (
	"fmt"
	"strconv"
)

// 經常會有將基本數據類型轉換為string
// 或者是string轉換為基本數據類型

func main() {
	// 基本類型轉string類型
	var num1 int = 100
	var num2 float64 = 2.54
	var flag bool = false
	var num3 byte = 'A'

	// 方式1 這個方式用的是最多的
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Println("str = ", str)

	str = fmt.Sprintf("%f", num2)
	fmt.Println("str = ", str)

	str = fmt.Sprintf("%t", flag)
	fmt.Println("str = ", str)

	str = fmt.Sprintf("%d or %c", num3, num3)
	fmt.Println("str = ", str)

	// 方式2 使用strconv函數
	//str = strconv.FormatInt(num1, 10) // 這裏需要的是int64
	str = strconv.FormatInt(int64(num1), 10) // 這裏需要的是int64
	fmt.Println("int -> strconv str = ", str)

	// 'f'是代表轉成的一種格式 10表示小數部分保留10位 64表示這個數是float64
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Println("float -> strconv str = ", str)

	str = strconv.FormatBool(flag)
	fmt.Println("bool -> strconv str = ", str)

	// 有一個itoa()函數 比較好用
	num4 := 10
	str = strconv.Itoa(num4)
	fmt.Println("Itoa() -> strconv str = ", str)

	fmt.Println("-----------------------------")

	// string轉基本數據類型 使用的是strconv內部的函數
	//var str2 string = "true"
	//var num11  = 99999
	//var num12 float64 = 4456.435
	//var num13 byte = 4

	flag2, _ := strconv.ParseBool("true") // 這個函數有兩個返回值 這裏用 _ 忽略掉了一個
	// 如果沒轉換成功 會給賦一個默認值 false
	// 其他類型也一樣
	fmt.Println("bool to string flag2 = ", flag2)

	outNum1, _ := strconv.ParseInt("113", 10, 64)
	fmt.Println("string to int = ", outNum1)

	outNum2, _ := strconv.ParseFloat("123.234", 64)
	fmt.Println("string to float = ", outNum2)

}
