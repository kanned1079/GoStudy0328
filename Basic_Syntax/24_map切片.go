package main

import "fmt"

// map切片的使用
// 要求:使用一个map来记录monster的信息 name 和 age
// 也就是说一个monster对应一个map,并且妖怪的个数可以动态的增加=>map切片
func main() {
	// 生命一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	// map需要make 切片也需要make
	// 增加第一个
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}
	// 下面这个写法会越界
	if monsters[2] == nil {
		monsters[2] = make(map[string]string, 2)
		monsters[2]["name"] = "狐狸精"
		monsters[2]["age"] = "300"
	}
	// 需要使用到切片的append函数 可以动态增加
	// 先定一个一个newMonsters
	newMonster := map[string]string{
		"name": "新的妖怪",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
