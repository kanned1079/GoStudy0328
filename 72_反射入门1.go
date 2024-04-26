package main

import (
	"fmt"
	"reflect"
)

// 基本介绍
// 1)反射可以在运行时动态获取变量的各种信息，比如变量的类型，类别
// 2)如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段、方法)
// 3)通过反射，可以修改变量的值，可以调用关联的方法。
// 4)使用反射，需要 import (“reflect”)

// reflect包实现了运行时反射，允许程序操作任意类型的对象。典型用法是用静态类型interface{}保存一个值
// 通过调用TypeOf获取其动态类型信息，该函数返回一个Type类型值。调用ValueOf函数返回一个Value类型值
// 该值代表运行时的数据。Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。

// 1)reflect.TypeOf(变量名)，获取变量的类型，返回reflect.Type类型
// 2)reflect.ValueOf(变量名)获取变量的值，返回reflect.Value类型 reflect.Value是一个，结构体类型。【看文档】
//   通过reflect.Value，可以获取到关于该变量的很多信息。
// 3)变量、interface{} 和 reflect.Value是可以相互转换的，这点在实际开发中，会经常使用到。画出示意图

// 在实际使用reflect的过程中

func main() {
	//num := 100 // 目标是通过reflect 拿到 value, type, kind
	//test72_1(num)

	var stu Student72 = Student72{
		Name: "Tom",
		Age:  24,
	}

	test72_2(stu)
}

// 基本流程
// test72_1 对基本数据类型进行反射 专门演示反射
func test72_1(b interface{}) { // 因为这里可以对任何数据类型进行反射
	// 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b) // 前面的r是代表反射到的类型
	fmt.Println("rTyp =", rTyp)
	fmt.Println("rTyp.Name =", rTyp.Name())

	// 获取到reflect.value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal =", rVal) // 到这里可以进行任何操作
	fmt.Printf("rval.type = %T\n", rVal)

	// 取到值
	n2 := 2 + rVal.Int()
	fmt.Printf("n2 = %d, type = %T\n", n2, n2) // n2 = 102, type = int64

	// 转换回去 将rVal转换为interface{}
	iV := rVal.Interface()
	// 将Interface通过断言转换为需要的类型
	num2 := iV.(int)
	fmt.Printf("num2 = %d, type = %T\n", num2, num2) // n2 = 102, type = int64
}

type Student72 struct {
	Name string
	Age  int
}

type Monster72 struct {
	Name string
	Age  int
}

func test72_2(b interface{}) {
	// 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b) // 前面的r是代表反射到的类型
	fmt.Println("rTyp =", rTyp)

	// 获取到reflect.value
	rVal := reflect.ValueOf(b)

	// 获取到对应的Kind
	// 1> rVal.Kind()
	// 2> rTyp.Kind() 这两个获取到的是一样的
	fmt.Printf("kind = %v, kind = %v\n", rVal.Kind(), rTyp.Kind())

	// 转为接口
	iV := rVal.Interface()
	fmt.Printf("iv = %v, iv = %T\n", iV, iV) // iv = {Tom 24}, iv = main.Student72

	// 取到字段信息 需要进行断言
	stu, ok := iV.(Student72)
	if ok {
		fmt.Printf("Stu.Name = %v, Stu.Age = %v\n", stu.Name, stu.Age)
	}

}
