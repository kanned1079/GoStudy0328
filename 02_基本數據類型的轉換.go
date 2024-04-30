package main

import "fmt"

func main() {
	// Golang中不支持自動轉換 需要顯式轉換 也就是強制類型轉換
	var i int = 100
	var j float32 = float32(i) // 得這麼個寫
	var k int8 = int8(j)
	// 把int8鑽換為int64或者int64轉換為int32都不行
	fmt.Printf("i = %d, j = %f, k = %d\n", i, j, k)

	var a int64 = 99999999
	var b int8
	b = int8(a) // 這樣寫不會報錯
	fmt.Println("int8 b = ", b)

}
