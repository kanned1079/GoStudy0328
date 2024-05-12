package main

import (
	"testing"
)

// testing 提供对 Go 包的自动化测试的支持。通过 `go test` 命令
// 能够自动执行如下形式的任何函数：
// func TestXxx(*testing.T)

//func main() {
//
//}

// 测试用例
// 命名需要Test开头加上函数名称
// 其中 Xxx 可以是任何字母数字字符串（但第一个字母不能是 [a-z]），用于识别测试例程。
func TestCalUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误， 期望值= %v，实际值= %v\n", 55, res)
		t.Fatalf("AddUpper(10) 执行错误， 期望值= %v，实际值= %v\n", 55, res)
	}
	// 如果正确 输出日志
	t.Logf("AddUpper(10) 执行正确， 期望值= %v，实际值= %v\n", 55, res)
}

func TestHello(t *testing.T) {
	//fmt.Println("TestHello()被调用")
	t.Logf("\"TestHello()被调用\"")
}

func TestGetSub(t *testing.T) {
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("getSub(10, 3) 执行错误， 期望值= %v，实际值= %v\n", 7, res)
	}
	t.Logf("getSub(10, 3) 执行正确， 期望值= %v，实际值= %v\n", 7, res)
}
