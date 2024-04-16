package main

import (
	"fmt"
	"testing"
)

// Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试
// testing框架和其他语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例。
// 通过单元测试可以解决如下问题:
//		1)确保每个函数是可运行，并且运行结果是正确的
//		2)确保写出来的代码性能是好的
//		3)单元测试能及时的发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决
//	   	  而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定

func main() {
	var res int
	// 以下是传统的测试方法 不方便 不利于管理
	if res = addUpper(100); res != 5050 {
		fmt.Printf("返回错误，返回值 = %v， 期望值 = %v\n", res, 5050)
	} else {
		fmt.Printf("返回正确，返回值 = %v， 期望值 = %v\n", res, 5050)

	}

}

// 这里有一个函数 怎样确定它运行的结果是否是正确的
func addUpper(a int) int {
	res := 0
	for i := 1; i <= a; i++ {
		res += i
	}
	return res
}

func addUpper2(a int) int {
	res := 0
	for i := 1; i <= a; i++ {
		res += i
	}
	return res
}

func TestAddUpper(t *testing.T) {
	// 定义测试用例
	testCases := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 3},
		{input: 3, expected: 6},
		{input: 4, expected: 10},
		{input: 5, expected: 15},
		{input: 10, expected: 55},
	}

	// 遍历测试用例
	for _, tc := range testCases {
		t.Run("addUpper test", func(t *testing.T) {
			// 调用 addUpper 函数
			result := addUpper(tc.input)
			// 检查结果是否与预期值相符
			if result != tc.expected {
				t.Errorf("addUpper(%d) = %d; want %d", tc.input, result, tc.expected)
			}
		})
	}
}
