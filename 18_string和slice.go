package main

import (
	"fmt"
)

// string的底層是byte數組 因此string可以進行切片處理操作

func main() {
	str1 := "hello@kanned"
	// 使用切片獲取到kanned
	slice1 := str1[6:]
	fmt.Println("獲取到kanned = ", slice1)

	// string是不可變的 因此沒法通過 slice1[0] = 'z'修改
	//str1[0] = 'z'

	// 將"hello@kanned" 修改為 "zello@kanned"
	slice2 := []rune(str1)
	slice2[0] = '北' // 要修改漢字的話這裏就用 []rune 因為rune是按照字符來切片的
	str1 = string(slice2)
	fmt.Println("修改後的str1 = ", str1)

	// 轉換成[]byte後可以修改英文和數字 但是不能對中文進行處理
	// []byte是按照字節來處理的 而一個漢字需佔用3個字節

}
