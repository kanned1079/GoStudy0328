package main

import (
	"fmt"
)

// 不管调用形式如何，真正决定是值拷贝还是地址拷贝 看这个方法是和哪个类型绑定
// 如果是和值类型，比如(p Person)则是值拷贝，如果和指针类型，比如是(p*Persom)则是地址拷贝。

// 编写结构体(MethodUtls)，编程一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形，在 main 方法中调用该方法。
// 编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
// 编写一个方法算该矩形的面积(可以接收长1en，和宽width)， 将其作为方法返回值。在 main方法中调用该方法，接收返回的面积值并打印。

type struct1 struct {
	width  float64
	height float64
}

func (s struct1) funcSt1() {
	fmt.Printf("width = %v, height = %v\n", s.width, s.height)
}

func (s struct1) funcSt2(m, n float64) {
	s.width = m
	s.height = n
	fmt.Printf("width = %v, height = %v\n", s.width, s.height)

}

func (s struct1) funcSt3(m, n float64) (area float64) {
	s.width = m
	s.height = n
	area = s.width * s.height
	return
}

// 判断一个数是奇数还是偶数
// 根据行、列、字符打印对应行数和列数的字符，比如:行:3，列:2，字符*,则打印相应的效果
// 定义小小计算器结构体(Calcuator)，实现加减乘除四个功能
// 实现形式1:分四个方法完成:
// 实现形式2:用一个方法搞定

type intger int

func (num intger) judgeNum() (res bool) {
	if num%2 == 0 {
		res = true
	} else {
		res = false
	}
	return
}

type draws struct {
	rows int
	cols int
	ch   byte
}

func (d draws) drawPic() {
	for i := 0; i < d.rows; i++ {
		for j := 0; j < d.cols; j++ {
			fmt.Printf("%c ", d.ch)
		}
		fmt.Println()
	}
}

type Calculator_32 struct {
	num1 float64
	num2 float64
	sign byte
}

// 方法1
func (c Calculator_32) cal_Add(n1, n2 float64) (res float64) {
	c.num1 = n1
	c.num2 = n2
	res = c.num1 + c.num2
	return
}

func (c Calculator_32) cal_Sub(n1, n2 float64) (res float64) {
	c.num1 = n1
	c.num2 = n2
	res = c.num1 - c.num2
	return
}

func (c Calculator_32) cal_Mul(n1, n2 float64) (res float64) {
	c.num1 = n1
	c.num2 = n2
	res = c.num1 * c.num2
	return
}

func (c Calculator_32) cal_Div(n1, n2 float64) (res float64) {
	c.num1 = n1
	c.num2 = n2
	res = c.num1 / c.num2
	return
}

// 方法2
func (c Calculator_32) cal_Auto(n1, n2 float64, sign byte) (res float64) {
	c.num1 = n1
	c.num2 = n2
	c.sign = sign
	switch {
	case sign == '+':
		{
			res = c.num1 + c.num2
		}
	case sign == '-':
		{
			res = c.num1 - c.num2
		}
	case sign == '*':
		{
			res = c.num1 * c.num2
		}
	case sign == '/':
		{
			res = c.num1 / c.num2
		}
	}
	return
}

func main() {
	rect1 := struct1{
		width:  10.0,
		height: 8.0,
	}
	fmt.Print("函数1 ：")
	rect1.funcSt1()

	var rect2 struct1
	fmt.Print("函数2 ：")
	rect2.funcSt2(4, 9)

	fmt.Print("函数3 ：")
	area := rect2.funcSt3(8, 9)
	fmt.Println("Area = ", area)

	var num1 intger = 12
	flag1 := false
	flag1 = num1.judgeNum()
	if flag1 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	var pic1 draws
	pic1.rows = 3
	pic1.cols = 5
	pic1.ch = '+'
	pic1.drawPic()

	var calculatorTest Calculator_32
	fmt.Println("方法1：10 + 43 = ", calculatorTest.cal_Add(10, 43))
	fmt.Println("方法2：128 / 8 = ", calculatorTest.cal_Auto(128, 8, '/'))

	// 方法和函数
	// 调用方式不一样
	// 函数的调用方式： 函数名(实参列表)
	// 方法的调用方式： 变量.方法名(实参列表)

	// 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
	// 对于方法(如struct的方法)接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以

}
