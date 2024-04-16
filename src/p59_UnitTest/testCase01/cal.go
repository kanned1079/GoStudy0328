package main

//  1)测试用例文件名必须以 test.go 结尾。比如 cal test.go,cal 不是固定的。
//  2)测试用例函数必须以Test开头，一般来说就是Test+被测试的函数名，比如TestAddUpper。
//  3)TestAddUpper(t *tesing.T)的形参类型必须是 *testing.T 【看一下手册】
//  4)一个测试用例文件中，可以有多个测试用例函数，比如TestAddUpper、TestSub4)
//  5)运行测试用例指令
// 		(1)cmd>go test[如果运行正确，无日志，错误时，会输出日志]
//		(2)cmd>go test-v [运行正确或是错误，都输出日志]
//  6)当出现错误时，可以使用t.Fatalf 来格式化输出错误信息，并退出程序
//  7)t.Logf 方法可以输出相应的日志
//  8)测试用例函数，并没有放在main函数中，也执行了，这就是测试用例的方便之处[原理图]
//  9)PASS表示测试用例运行成功，FAIL表示测试用例运行失败
// 10)测试单个文件，一定要带上被测试的原文件go test -v cal test.go cal.go
// 11)测试单个方法go test -v -test.run TestAddUpper

// AddUpper 这里有一个函数 怎样确定它运行的结果是否是正确的
func AddUpper(a int) int {
	res := 0
	for i := 1; i <= a; i++ {
		res += i
	}
	return res
}

func AddUpper2(a int) int {
	res := 0
	for i := 1; i <= a; i++ {
		res += i
	}
	return res
}

func getSub(n1, n2 int) int {
	return n1 - n2
}
