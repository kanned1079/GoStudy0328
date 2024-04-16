package main

import "testing"

//func main() {
//	var monster12 Monster
//	monster12.Name = "aaaa"
//	monster12.Age = 600
//	monster12.Skill = "unkonw"
//	monster12.Store()
//
//}

func TestThisMain(t *testing.T) {
	var monster1 = Monster{
		Name:  "aaaa",
		Age:   600,
		Skill: "unkonw",
	}
	if res := monster1.Store(); res == false {
		t.Fatalf("运行错误")
	}
	t.Logf("ok")
}
