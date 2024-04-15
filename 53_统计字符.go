package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 统计英文、数字、空格和其他字符数量
// 说明:统计一个文件中含有的英文、数字、空格及其它字符数量

// 定一个结构体拿来保存统计结果
type charCount struct {
	chCount    int
	numCount   int
	spaceCount int
	otherCount int
}

func main() {
	fileName := "./dir/text"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件错误 = ", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("关闭文件错误 = ", err)
		}
	}()
	var cut charCount // 创建实例
	// 创建一个reader
	reader := bufio.NewReader(file)
	// 循环读取
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("结束")
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				cut.chCount++
			case v == ' ' || v == '\t':
				cut.spaceCount++
			case v >= '0' && v <= '9':
				cut.numCount++
			default:
				cut.otherCount++
			}
		}
	}
	fmt.Printf("字符 = %d\n数字 = %d\n空格 = %d\n其他 = %d\n", cut.chCount, cut.numCount, cut.spaceCount, cut.otherCount)
}
