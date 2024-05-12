package main

import (
	"fmt"
)

// 张老太养了两只猫猫:一只名字叫小白,今年3岁,白色。
// 还有一只叫小花,今年100岁,花色。
// 请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，年龄，颜色。
// 如果用户输入的小猫名错误，则显示 张老太没有这只猫猫。

type Cats struct {
	Name  string
	Age   int
	Color string
}

// Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其它00P语言不一样
// 比如继承Golang没有extends 关键子，继承是通过匿名字段来实现。
// Golang面向对象(00P)很优雅，00P本身就是语言类型系统(type system)的一部分通过接口(interface)关联，耦合性低
// 在Golang中--面向接口编程--是非常重要的特性。
type Persion1 struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type Person27 struct {
	Name string
	Age  int
}

type monster1 struct {
	Name string
	Age  int
}

type monster2 struct {
	Name string
	Age  int
}

func main() {
	// 使用現有技術解決
	// 1.使用定義變量
	// 2.使用數組
	// 3.使用map
	// 使用以上技術不利於數據的管理
	//var cat1, cat2 Cats
	//cat1.Name = "小白"
	//cat1.Age = 3
	//cat1.Color = "白色"
	//cat2.Name = "小花"
	//cat2.Age = 100
	//cat2.Color = "花色"
	//
	//fmt.Println("cat1 = ", cat1)
	//fmt.Println("cat2 = ", cat2)

	//如果结构体的字段类型是:指针，slice，和map的零值都是 nil
	// 如果需要使用这样的字段，需要先make，才能使用. 即还没有分配空间

	var p1 Persion1
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}
	// 结构体里的指针 切片 map都还是需要make后才能使用
	p1.slice = make([]int, 2)
	p1.slice[0] = 12
	p1.map1 = make(map[string]string)
	p1.map1["NO0"] = "aaaaaa"

	fmt.Println(p1)

	// 不同的結構體變量是相互獨立的 互不影響
	// 傳遞的是值類型
	var monst1 monster1
	var monst2 monster2
	monst1.Name = "aaaaa"
	monst1.Age = 200
	monst2.Name = "bbbbb"

	// 方式2
	p2 := Person27{"Tom", 18}
	fmt.Println("p2 = ", p2)

	// 方式3 返回的是结构体指针
	var p3 = new(Person27)
	// 因为p3是一个指针 因此标准的给字段赋值的方法是
	(*p3).Name = "smith"
	p3.Name = "Smith" // Go的設計者為了使用方便 在底層做了處理 會給p3做了取值運算
	(*p3).Age = 19
	fmt.Println("p3 = ", *p3)

	// 方式4 返回的是结构体指针
	var p4 *Person27 = &Person27{}
	//var p4 *Person27 = &Person27{"Mary", 20}
	// 也可以创建的时候赋值
	(*p4).Name = "kanna"
	p4.Age = 21
	fmt.Println("p4 = ", p4)

}
