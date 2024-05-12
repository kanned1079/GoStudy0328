package main

import (
	"fmt"
	"reflect"
)

// 给你一个变量 varvfloat64=1.2
// 请使用反射来得到它的reflect.Value，然后获取对应的TypeKind和值，
// 并将reflect.Value转换成interfacel{}
// 再将interfacel{} 转换成float64

func main() {
	var v float64 = 1.2
	reflectTest0425(&v)
	fmt.Println("main中的v =", v)
}

func reflectTest0425(b interface{}) {

	rVal := reflect.ValueOf(b)

	//reflect.ValueOf(b).Elem().SetFloat(2.76)

	num := rVal.IsValid()

	//rVal.SetFloat(5.6)
	fmt.Println("num =", num)
	fmt.Printf("num.Type = %T\n", num)
	rVal.Elem().SetFloat(67.87)

	iV := rVal.Interface()
	fmt.Println("iV =", iV)
	fmt.Printf("iV.Type =%T\n", iV)

	num3 := iV.(*float64)

	fmt.Println(*num3)

	//iV := rVal.Interface()
	//fmt.Println("iV =", iV)

	//iV := rVal.Interface()

	//if num, ok := iV.(float64); ok {
	//	fmt.Println("num =", num)
	//} else {
	//	fmt.Println(ok)
	//}

}
