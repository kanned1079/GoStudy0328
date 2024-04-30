package main

import (
	"encoding/json"
	"fmt"
)

// 任何数据类型都可以转换成JSON格式

type monster struct {
	Name     string  `json:"name"` // 反射机制
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	var monster1 monster = monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-09-23",
		Sal:      5600.6,
		Skill:    "不知道",
	}
	toJson(monster1)
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{}, 1)
	a["name"] = "aaaa"
	a["age"] = 18
	a["address"] = "china"

	toJson(a)
}

func testSlice() {
	var slice []map[string]interface{} // slice
	var m1 map[string]interface{}      // map
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{} // map
	m2 = make(map[string]interface{})
	m2["name"] = "jack"
	m2["age"] = 7
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	toJson(slice)

}

func testVar() {
	// 对基本数据类型序列化意义不大
	var num1 float64 = 4543.352
	toJson(num1)
}

func toJson(a interface{}) {
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败 err = ", err)
	}
	fmt.Println("序列化后 = ", string(data))
}

func main() {
	// 这里是反序列化
	testStruct()
	testMap()
	testSlice()
	testVar()
}
