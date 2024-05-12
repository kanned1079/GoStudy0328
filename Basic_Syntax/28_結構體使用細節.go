package main

import (
	"encoding/json"
	"fmt"
)

// 結構體在內存中分配的地址是連續的
type Point28 struct {
	x int
	y int
}

type Rect28 struct {
	leftUp, rightDown Point28
}

type Rect28_2 struct {
	leftUp, rightDown *Point28
}

type A struct {
	num int
}

type B struct {
	num int
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	r1 := Rect28{Point28{1, 2}, Point28{2, 8}}
	fmt.Printf("r1.leftUp.x = %p, r1.leftUp.y = %p\nr1.rightDown.x = %p, r1.rightDown.y = %p\n", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	fmt.Println("ri = ", r1)

	// r2有两个 *Point类型，这个两个*Point类型的本身地址也是连续的
	// 但是他们指向的地址不一定是连续
	r2 := Rect28_2{&Point28{10, 20}, &Point28{20, 80}}
	fmt.Printf("r2.leftUp = %p, r2.leftUp.y = %p\n", &r2.leftUp, &r2.leftUp)

	// 兩個結構體如果要相互轉換需要保持字段完全相同
	var a A
	var b B
	//a = b
	// 這樣不行
	a = A(b)
	fmt.Println("a = ", a)

	// struct的每个字段上，可以写上一个tag
	// 该tag可以通过反射机制获取，常见的使用场景就是--序列化--和反序列化。

	// json处理后的字段名也是首字母大写，这样如果我们是将json后的字符串返回给其它程序使用
	// 比如jqueryphp等，那么可能他们的不习惯这个命名方式，怎么办?
	// 将Monster的字段首字母小写，这个行不通，你会发现处理后，返回的是空字符串
	// 因为ison.Marshal相当于在其它包访问monster结构体，你首字母小写就不能在其它包访问。

	monster1 := Monster{"牛魔王", 500, "芭蕉扇"}
	// 需要進行序列化轉為字符串 -> JSON
	// json.Marshal()中使用到了反射
	jsonMonster1, err := json.Marshal(monster1)
	if err != nil {
		fmt.Println("error = ", err)
	} else {
		fmt.Println("jsonMonster1 = ", string(jsonMonster1)) // 那麼這裏是大寫的怎麼辦
	}

	//	[	]
	//	[	]
	//	[	]

}
