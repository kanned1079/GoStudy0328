package main

import (
	"errors"
	"fmt"
)

func main() {
	// 自定義錯誤
	// 使用 errors.New和panic的內置函數
	// errors.New("錯誤說明") 會返回一個error類型的值 表示一個錯誤

	// 測試
	test002()
	fmt.Println("main()下面的代碼...")
}

/*
函數去讀取init.conf的值
如果傳入不正確 那就返回一個自定義錯誤
*/
func readConfig(name string) (err error) {
	if name == "init.conf" {
		return nil
	} else {
		return errors.New("讀取文件錯誤")
	}
}

func test002() {
	err := readConfig("init.conf")
	if err != nil {
		panic(err)
		// 這裏執行後 後面的代碼都不會再執行了
	}
	fmt.Println("test002後面的代碼繼續執行...")
}
