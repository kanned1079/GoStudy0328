package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// func Copy(dst Writer, src Reader) (written int64, err error)
// 将src的数据拷贝到dst，直到在src上到达EOF或发生错误。
// 返回拷贝的字节数和遇到的第一个错误。

const srcFileName = "./imgs/testpic.JPG"
const destFileName = "./imgs/testpic2.JPG"

// 自己编写一个函数
// 它接受两个文件路径 一个是src一个是dest
func copyFile(destName, srcName string) (written int64, err error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		fmt.Println("打开文件错误 = ", err)
	}
	defer func() {
		if err := srcFile.Close(); err != nil {
			fmt.Println("关闭失败")
		} else {
			fmt.Println("Success")
		}
	}()
	reader := bufio.NewReader(srcFile)

	destFile, err := os.OpenFile(destName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("打开文件2错误 = ", err)
	}
	defer func() {
		if err := destFile.Close(); err != nil {
			fmt.Println("关闭失败")
		} else {
			fmt.Println("Success")
		}
	}()
	writer := bufio.NewWriter(destFile)
	func() {
		if err := writer.Flush(); err != nil {
			fmt.Println("flush失败")
		} else {
			fmt.Println("Success")
		}
	}()
	return io.Copy(writer, reader)
}

func main() {
	start := time.Now()
	//if _, err := copyFile(destFileName, srcFileName); err != nil {
	//	fmt.Println("Error = ", err)
	//}
	fmt.Println(copyFile(destFileName, srcFileName))
	end := time.Now()
	fmt.Println("time = ", end.Sub(start))
}
