package main

import (
	"fmt"
	"reflect"
)

// 使用反射的方式来获取结构体tag标签
// 遍历字段的值 修改字段值
// 调用结构体方法
// 要求使用传地址的方式

type A76 struct {
	Str string `json:"str"`
	Num int    `json:"num"`
}

func (this A76) GetSum(a, b int) int {
	return a + b
}

func main() {
	examp := A76{
		Str: "swegtvedrh",
		Num: 89,
	}
	reflectTest76(&examp)
	fmt.Println("examp(main) =", examp)

}

func reflectTest76(in interface{}) {
	rTpy := reflect.TypeOf(in)
	rVal := reflect.ValueOf(in)
	rKind := rVal.Kind()
	if rKind != reflect.Ptr && rVal.Elem().Kind() == reflect.Struct {
		fmt.Println("非法输入")
		return
	}

	fmt.Println("rTpy = ", rTpy)
	fmt.Println("rVal = ", rVal)

	fieldNums := rTpy.Elem().NumField()
	funcNums := rVal.Elem().NumMethod()
	fmt.Println("字段数 = ", fieldNums, ", 方法数 = ", funcNums)

	for i := 0; i < fieldNums; i++ { // 遍历字段名称和别称
		// 获取字段的值
		fmt.Printf("字段%d的值 = %v\n", i, rVal.Elem().Field(i))
		tagVal := rTpy.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("字段%d的标签 = %v\n", i, tagVal)
		}
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf(20))
	params = append(params, reflect.ValueOf(79))
	res := rVal.Elem().Method(0).Call(params) // 调用方法
	fmt.Println("调用了方法 res = ", res[0].Int())

	var fieldName1 string = "Str" // 要修改的字段的名称
	if rVal.Elem().FieldByName(fieldName1).IsValid() {
		fmt.Println("字段合法")
		if rVal.Elem().FieldByName(fieldName1).CanSet() {
			fmt.Println("字段可以被设置")
			rVal.Elem().FieldByName(fieldName1).SetString("kinggyo")
		} else {
			fmt.Println("字段不可以被设置")
		}
	} else {
		fmt.Println("字段不合法")
	}

}
