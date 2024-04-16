package main

import (
	"encoding/json"
	"fmt"
)

type monster57 struct {
	Name     string  `json:"name"` // 反射机制
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

/*
序列化后 =  {"name":"牛魔王","age":500,"birthday":"2011-09-23","sal":5600.6,"skill":"不知道"}
序列化后 =  {"address":"china","age":18,"name":"aaaa"}
序列化后 =  [{"address":"北京","age":7,"name":"jack"},{"address":["墨西哥","夏威夷"],"age":7,"name":"jack"}]
序列化后 =  4543.352
*/

func testUnmarshalStruct() {
	str := "{\"name\":\"牛魔王\",\"age\":500,\"birthday\":\"2011-09-23\",\"sal\":5600.6,\"Skill\":\"不知道\"}"
	// 定义一个monster实例
	var monst1 monster57
	if err := json.Unmarshal([]byte(str), &monst1); err != nil {
		fmt.Println("FAILURE")
	} else {
		fmt.Println("SUCCESS")
	}
	fmt.Println("monst1.Name = ", monst1.Name)
	fmt.Println("monst1.Age = ", monst1.Age)
	fmt.Println("monst1.Birthday = ", monst1.Birthday)
	fmt.Println("monst1.Sal = ", monst1.Sal)
	fmt.Println("monst1.Skill = ", monst1.Skill)
}

func testUnmarshalMap() {
	var a map[string]interface{}
	str := "{\"address\":\"china\",\"age\":18,\"name\":\"aaaa\"}"
	//a = make(map[string]interface{})
	// 注意
	// map的反序列化不需要make 因为make的操作已经被封装到Unmarshal()里了
	if err := json.Unmarshal([]byte(str), &a); err != nil {
		fmt.Println("FAILURE")
	} else {
		fmt.Println("SUCCESS")
	}
	fmt.Println("反序列化map = ", a)
}

func testUnmarshalSlice() {
	str := `[{"address":"北京","age":7,"name":"jack"},{"address":["墨西哥","夏威夷"],"age":7,"name":"jack"}]`
	var slice []map[string]interface{} // slice
	if err := json.Unmarshal([]byte(str), &slice); err != nil {
		fmt.Println("FAILURE")
	} else {
		fmt.Println("SUCCESS")
	}
	fmt.Println("反序列化slice = ", slice)
}
func main() {
	// 反序列化一个json字符串时 要保证数据类型一致
	// 字段顺序无所谓
	testUnmarshalStruct()
	testUnmarshalMap()
	testUnmarshalSlice()
}
