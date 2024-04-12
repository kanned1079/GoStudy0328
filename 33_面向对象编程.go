package main

import (
	"fmt"
)

// 一、声明(定义)结构体，确定结构体名
// 二、编写结构体的字段
// 三、编写结构体的方法

// 学生案例:
// 字段：分别为string、编写一个Student结构体，包含name、gender、age、id、score类型。
// 结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
// 在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。
type Student_33 struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu Student_33) sayHello() (str string) {
	str = fmt.Sprintf("name=%s, gender=%s, age=%d, ID=%d, score=%v\n", stu.name, stu.gender, stu.age, stu.id, stu.score)
	return
}

// 编程创建一个 Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
// 声明一个方法获取立方体的体积。
// 建一个Box结构体变量，打印给定尺寸的立方体的体积
type Box_33 struct {
	length float64
	width  float64
	height float64
}

func (box Box_33) getVolum() (vol float64) {
	vol = box.length * box.width * box.height
	return
}

// 一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元 其它情况门票免费
// 请编写visitor结构体，根据年龄段决定能够购买的门票价格并输出
type Visitors_33 struct {
	Name string
	Age  int
}

func (v Visitors_33) showPrice() {
	if v.Age >= 90 || v.Age <= 8 {
		fmt.Printf("不给玩")
	}
	if v.Age > 18 {
		fmt.Printf("游客名字是%s， 收费为%d\n", v.Name, 20)
	} else {
		fmt.Printf("免费")
	}

}

func main() {
	stu1 := Student_33{
		name:   "kanna",
		gender: "male",
		age:    19,
		id:     34534235,
		score:  98.32,
	}
	fmt.Println(stu1.sayHello())

	var box1 Box_33
	fmt.Print("length: ")
	fmt.Scan(&box1.length)
	fmt.Print("width: ")
	fmt.Scan(&box1.width)
	fmt.Print("height: ")
	fmt.Scan(&box1.height)
	fmt.Println("box1's volum = ", box1.getVolum())

}
