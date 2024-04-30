package main

import "fmt"

// map是引用类型，遵守引用类型传递的机制，在一个函数接收map
func modify(map1 map[int]int) {
	map1[1] = 1000
}

// 定义一个学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// 修改后，会直接修改原来的map
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 34
	map1[3] = 87
	map1[4] = 367
	map1[5] = 17
	fmt.Println(map1)
	modify(map1)
	fmt.Println(map1)

	// map的value也经常使用struct结构体更方便
	var students map[string]Stu
	students = make(map[string]Stu, 1)
	stu1 := Stu{"Tom", 18, "Beijnng"}
	stu2 := Stu{"Mary", 20, "Shanghai"}
	students["NO1"] = stu1
	students["NO2"] = stu2
	fmt.Println(students)
	// 進行遍歷
	for key, v := range students {
		fmt.Printf("學生的學生編號是%v, 名字是%v，年齡是%v，住址是%v\n", key, v.Name, v.Age, v.Address)

	}
}
