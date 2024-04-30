package main

import (
	"fmt"
	"os"
)

// 從文件到程序中叫做 輸入流
// 從程序到文件中叫做 輸出流
// os.File 封裝了所有的相關操作 File是一個結構體

func main() {
	// 打開一個文件 os.Open(name string) (*File, error)
	file1, err := os.Open("./Files/a.tx1t")
	// file1 叫file对象
	// file1 叫file指针
	// file1 叫file句柄
	if err != nil {
		fmt.Println("打开文件错误：", err)
	}

	fmt.Println(file1)
	// 输出文件地址 &{0x1400011a120}

	file1.Close()
	os.Open("./Files/a.txt")

}
