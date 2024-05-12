package main

import (
	"fmt"
	"reflect"
)

// Tpye是类型
// Kind是类别
// Type和Kind可能是相同 也可能是不同的

// 变量 <----> interface{} <-----> reflect.Value

// 使用反射获取 类型的值 要保证类型匹配 不然会导致panic 需要用断言
// iVal.Int()  iVal.Float()

// 通过反射来 修改num int和stu
func main() {
	var num int = 120
	reflectTest73(&num) // 这里要取地址
	fmt.Println("main中num =", num)

	// Elem()的作用可以理解为
	//num := 9
	//ptr *int = &num
	//num2 := *ptr
}

func reflectTest73(b interface{}) {
	//rType := reflect.TypeOf(b)
	//fmt.Println("rType =", rType)

	rVal := reflect.ValueOf(b)
	// 这个rVal有提供一个方法
	//rVal.SetInt(160)
	fmt.Println("rVal =", rVal)
	fmt.Printf("rVal.Kind() = %v\n", rVal.Kind()) // 这是个指针
	// 那这里传指针才能改变值 那么得有方法获取到值
	//rVal.SetInt(130)
	// 报错 panic: reflect: reflect.Value.SetInt using unaddressable value
	// 因为上面传递的是地址
	fmt.Println(rVal)

	rVal.Elem().SetInt(240) // 这样会修改num的值
	// func (v Value) Elem() Value
	// Elem返回v持有的接口保管的值的Value封装，或者*v持有的指针指向的值的*Value封装。
	// 如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。

	//iV := rVal.Interface()
	//fmt.Println("num2 =", iV.(int))
}
