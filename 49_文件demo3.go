package main

import (
	"fmt"
	"io/ioutil"
)

// 读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中)
// 这种方式适用于文件不大的情况
// 相关方法和函数(ioutil.ReadFile)

func main() {
	file1 := "./Files/a.txt"
	// 这个方法现在已被弃用
	content, err := ioutil.ReadFile(file1) // 这个的返回值是 []byte, error
	if err != nil {
		fmt.Scanln(err)
	}
	//fmt.Printf("%v", content) // 这样输出的是byte切片
	fmt.Printf("%v", string(content)) // 需要转换成string
	// 因为没有显式的open文件 所以也没有close文件
	// 因为Open和Close被封装到ReadFile里了

}
