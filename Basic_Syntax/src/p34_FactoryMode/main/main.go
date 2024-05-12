package main

import (
	"GoStudy0328/src/p34_FactoryMode/model"
	"fmt"
)

// Golang的结构体没有构造函数，通常可以使用式来解决这个问题
//看一个需求
//一个结构体的声明是这样的:
// package model
// type Student struct { Name string.... }
// 因为这里的student 的首字母S是大写的，如果我们想在其它包创建Student的实例(比如main包)
// 引入model包后，就可以直接创建Student结构体的变量(实例)。
// 但是问题来了，如果首字母是小写的，比如 是 type student struct{...}就不不行了
// 怎么办---> 工厂模式来解决.

func main() {
	// 這裏大寫是ok的
	//var stu1 = model.Student{
	//	Name:  "Tom",
	//	Score: 78.9,
	//}

	// 當student小寫時候 可以通過工廠模式解決
	var stu1 = model.NewStudent("Tom", 88.9)
	//fmt.Println(*stu1)
	//fmt.Println(stu1.Name, stu1.Score) // 是指針 所以前面有一個地址符&

	stu1.SetName("Kanna")
	fmt.Println("name = ", stu1.GetName())

}
