package main

import "fmt"

func main() {
	//aErrFunc()
	// 默認發生錯誤(panic)後 程序會崩潰停止運行 所以要進行錯誤處理
	// 在Golang中沒有try catch
	// Go中引入的處理方式是 defer panic recover
	// 使用defer和recover來進行處理

	aErrFunc()

	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}()

	fmt.Println("下面的代码......")

}

func aErrFunc() {

	// 使用匿名函數完成異常捕獲
	defer func() {
		//err := recover()
		if err := recover(); err != nil { //上一行的可以寫到這裡
			fmt.Println("錯誤信息:", err)
			// 這裏可以把錯誤信息發送給管理員

		}
	}()

	a := 10
	b := 0
	fmt.Printf("resu = ", a/b)

}
