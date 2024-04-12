package main

import "fmt"

// 继承基本介绍和示意图
// 继承可以解决代码复用,让我们的编程更加靠近人类思维，
// 当多个结构体存在相同的属性(字段)和方法时,可以从这些结构体中抽象出结构体(比如刚才的student)
// 在该结构体中定义这些相同的属性和方法。
// 其它的结构体不需要重新定义这些属性和方法，只需嵌套一个Student匿名结构体即可。
// 也就是说:在Golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性。
type Student_37 struct {
	Name  string
	Age   int
	Score int
}

type Pupil struct {
	Student_37
}

type Graduate struct {
	Student_37
}

// 給Student_37一個方法 兩個子類都可以調用這個方法
func (stu *Student_37) getSum(a, b int) (res int) {
	res = a + b
	return
}

func (stu *Student_37) ShowInfo() {
	fmt.Printf("姓名=%s, 年龄=%d, 成绩=%d\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student_37) SetScore(s int) {
	stu.Score = s
}

// 这是Pupil特有的方法 虽然方法名相同
func (pu *Pupil) Testing() {
	fmt.Printf("小学生在考试...\n")
}

func (gra *Graduate) Testing() {
	fmt.Printf("大学生在考试...\n")
}

// 一些細節
type A_37 struct {
	Name string
	Age  int
}

type B_37 struct {
	Name  string
	Score float64
}

type C_37 struct {
	A_37
	B_37
	// 上面的A和B都有相同的Name 但是C沒有
	Name string
	// 要麼字節寫一個 要麼指定是誰的
}

// 如果一个struct嵌套了一个有名结构体，这种模式就是组合
// 如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字
type AA_37 struct {
	Name string
}

// 有名結構體
type BB_37 struct {
	Name string
	aa   AA_37 // 也叫組合關係
}

// 嵌套匿名结构体后
// 也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值。
type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct { // 這裏是多重繼承
	Goods
	Brand // 嵌入多個繼承的結構體
}

type TV2 struct {
	*Goods
	*Brand
}

// 在Golang中也允許匿名字段
type Stu_37 struct {
	Name string
	Age  int
}

type Stu_37_A struct {
	Stu_37
	int // 這是匿名字段
	// 如果這裏有一個int匿名字段了 那麼就不能再有第二個 除非起名
}

func main() {
	pu1 := Pupil{}
	pu1.SetScore(86)
	pu1.Name = "a"
	pu1.Age = 11
	pu1.ShowInfo()
	pu1.Testing()
	fmt.Println("pu1.getSum(1, 1) = ", pu1.getSum(1, 1))

	gra1 := Graduate{}
	gra1.SetScore(61)
	gra1.Name = "b"
	gra1.Age = 20
	gra1.ShowInfo()
	gra1.Testing()
	fmt.Println("gra1.getSum(32, 64) = ", gra1.getSum(32, 64))

	var cc1 C_37
	//cc1.Name = "A" // 這樣會不明確 會報錯
	//cc1.A_37.Name = "A" // 得這樣寫
	cc1.Name = "ss"
	fmt.Println(cc1.Name)

	var bb1 BB_37
	//bb.Name = "kanna" //這樣不行
	bb1.aa.Name = "kanna" // 得這麼寫
	bb1.Name = "kinggyo"  // 如果裏邊有那就行
	fmt.Println("bb1.aa.Name = ", bb1.aa.Name, ", bb1.Name = ", bb1.Name)

	// //////

	tv1 := TV{
		Goods{"電視機", 5000},
		Brand{"海爾", "山東青島"},
	}
	tv2 := TV{
		Goods{
			Name:  "電視機",
			Price: 6000,
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}

	fmt.Println(tv1)
	fmt.Println(tv2)

	tv3 := TV2{
		&Goods{
			Name:  "華為",
			Price: 8600,
		},
		&Brand{
			Name:    "電視2",
			Address: "上海",
		},
	}
	fmt.Println(tv3.Goods.Name, tv3.Price, tv3.Brand.Name, tv3.Address)

	anoy1 := Stu_37_A{
		Stu_37{
			Name: "string",
			Age:  18,
		},
		19, // 這裏直接對匿名字段賦值
	}
	fmt.Println(anoy1)
	anoy1.int = 21 // 這樣也可以
	fmt.Println(anoy1)

}
