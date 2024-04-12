package main

import "fmt"

// 注意事項1
// 接口本身不能创建实例,但是可以指向一个实现了该接口的 自定义类型的变量(实例)
type AInterface interface {
	Say()
}
type Stu_39 struct {
	Name string
}

func (stu Stu_39) Say() {
	fmt.Println("Stu Say()")
}

func main1() {
	var stu1 Stu_39
	var a AInterface = stu1
	a.Say() // 輸出Stu Say()
}

// 注意事項2
// 接口中所有的方法都没有方法体,即都是没有实现的方法。

// 注意事項3
// 在Golang中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口。

// 注意事項4
//一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量赋给接口类型。

// 注意事項5
// 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
type integer_39 int

func (i integer_39) Say() {
	fmt.Println("i = ", i)
}
func main2() {
	var i integer_39 = 10
	//i.Say()
	var inter AInterface = i
	inter.Say()
}

// 注意事項6
// 一個自定義類型可以實現多個接口
type BInterface interface {
	Hello()
}
type Monster_39 struct {
}

func (m Monster_39) Hello() {
	fmt.Println("Monster's Hello() ~~~")
}

func (m Monster_39) Say() {
	fmt.Println("Monster's Say() ~~~")
}

func main3() {
	// monster1實現了AInterface和BInterface接口
	var monster1 Monster_39
	var a1 AInterface = monster1
	var b1 BInterface = monster1
	a1.Say()
	b1.Hello()
}

// 注意事項7
// 接口中不能有任何變量

//type TestInterface interface {
//	Fun()
//	Say()
//	Name string // 接口中不能有任何變量
//}

// 注意事項8
// 比如(A接口)可以继承多个别的接口(比如B, C接口)
// 这时如果要实现A接口，也必须将B,C接口的方法也全部实现。

type B_Interface interface {
	test01()
}

type C_Interface interface {
	test02()
}

type A_Interface interface {
	B_Interface
	C_Interface
	test03()
}

func (stu Stu_39) test01() {

}

func (stu Stu_39) test02() {

}

// 如果這裏少任何一個實現的方法 那麼就會報錯
// 只有當多有方法都實現了才能稱為實現了接口 報錯：Stu_39 does not implement A_Interface (missing method test03)
func (stu Stu_39) test03() {

}

func main4() {
	var stu2 Stu_39
	var aa1 A_Interface = stu2
	aa1.test01()
}

// 注意事項9
// interface类型默认是一个指針(Interface是引用類型)
// 如果没有对interface初始化就使用，那么会输出nil

// 注意事項10
// 空接口interfacen 没有任何方法，所以所有类型都实现了空接口【案例演示】
// 也就是可以將任何數據類型傳遞給空接口
type T interface {
}

func main() {
	var stu1 Stu_39
	var t T = stu1
	fmt.Println(t)

	// 空接口也可以這樣寫
	var i1 interface{} = t
	fmt.Println(i1)
}
