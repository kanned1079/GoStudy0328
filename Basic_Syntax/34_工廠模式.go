package main

// Golang的结构体没有构造函数，通常可以使用式来解决这个问题
//看一个需求
//一个结构体的声明是这样的:
// package 1_Create
// type Student struct { Name string.... }
// 因为这里的student 的首字母S是大写的，如果我们想在其它包创建Student的实例(比如main包)
// 引入model包后，就可以直接创建Student结构体的变量(实例)。
// 但是问题来了，如果首字母是小写的，比如 是 type student struct{...}就不不行了
// 怎么办---> 工厂模式来解决.

func main() {

}
