package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 这种适用于大文件
// bufio的默认缓冲区为4096
// 读取文件的内容并显示在终端(带缓冲区的方式)
// 使用 os.Open,file.Close,bufo.NewReader0) reader.ReadString函数和方法

func main() {
	// 打開一個文件 os.Open(name string) (*File, error)
	file1, err := os.Open("./Files/a.txt")
	if err != nil {
		fmt.Println("打开文件错误：", err)
	}

	fmt.Println("file1的地址 = ", file1)
	// 输出文件地址 &{0x1400011a120}

	// 创建一个 *Reader
	reader1 := bufio.NewReader(file1)
	// 循环读取文件的内容
	for {
		str, err := reader1.ReadString('\n') // 读到换行符号
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")

	// defer 当函数退出时 要及时关闭file
	defer file1.Close()

}
