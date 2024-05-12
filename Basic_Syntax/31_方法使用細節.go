package main

import (
	"fmt"
)

// 普通的float32也可以有自己的方法

type integer int

func (i integer) print() {
	fmt.Println("i = ", i)
}

func (i *integer) change() {
	(*i) = (*i) + 1
	fmt.Println("i = ", *i)
}

// 如果一个类型实现String()这个方法
// 那么fmt.Println默认会调用这个变量的String()進行輸出
type Student_31 struct {
	Name string
	Age  int
}

// 通常用指針 速度很快
func (s *Student_31) String() string {
	str := fmt.Sprintf("stu[name] = %s, stu[age] = %d\n", s.Name, s.Age)
	return str
}

func main() {
	var i integer = 10
	i.change()
	i.print()

	stu1 := Student_31{
		Name: "kinggyo",
		Age:  18,
	}
	//fmt.Println(stu1) 這樣不行
	fmt.Println(&stu1) // 如果實現了 *Student的String()方法 就調用那個方法的Sprintf輸出s

}
