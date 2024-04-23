package main

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

}
