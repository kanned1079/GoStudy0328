package main

import (
	"fmt"
	"os"
)

// func Stat(name string) (fi FileInfo, err error)
// Stat返回一个描述name指定的文件对象的FileInfo。
// 如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接。
// 如果出错，返回的错误值为*PathError类型

func main() {
	fmt.Println(PathExist("./Files/a.tx"))
}

func PathExist(path string) (bool, error) {
	// 如果返回的错误为nil,说明文件或文件夹存在
	// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	// 如果返回的错误为其它类型,则不确定是否在存在
	var err error
	if _, err = os.Stat(path); err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
