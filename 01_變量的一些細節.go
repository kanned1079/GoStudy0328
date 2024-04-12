package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n2 int64 = 10
	// unsafe.Sizeof() 獲取變量的佔用空間
	fmt.Printf("n2 = %d, type = %T, size = %d\n", n2, n2, unsafe.Sizeof(n2))

	// 注意在Go中的變量 遵守保小不保大的原則
	// 在保證程序正常運行的情況下 盡量使用較小的變量
	//var age int64 // 這樣就不建議
	//var age byte = 90

	var flag bool = true
	// bool在內存中佔用一個字節
	fmt.Println("flag = ", flag)

	var str1 string = "Build simple, secure, scalable systems with Go"
	fmt.Println("str1 = ", str1)
	// Go支持顯示出字符串源代碼輸出

	// 使用反引號 ` [文本] `
	str1 = `
		fmt.Printf("n2 = %d, type = %T, size = %d\n", n2, n2, unsafe.Sizeof(n2))

		// 注意在Go中的變量 遵守保小不保大的原則
		/	/ 在保證程序正常運行的情況下 盡量使用較小的變量
		//var age int64 // 這樣就不建議
		//var age byte = 90`
	fmt.Println("str2 = ", str1)

}
