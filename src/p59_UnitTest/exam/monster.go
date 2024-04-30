package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// 1)编写一个Monster结构体,字段 Name,Age, Skill)
// 2)给Monster绑定方法Store，可以将一个Monster变量(对象),序列化后保存到文件中
// 3)给Monster绑定方法ReStore，可以将一个序列化的Monster,从文件中读取，并反序列化为Monster对象
// 4)编程测试用例文件 store test.go,编写测试用例函数 Teststore和 TestRestore进行测试。

//func main() {
//	var monster12 Monster
//	//monster12.Name = "aaaa"
//	//monster12.Age = 600
//	//monster12.Skill = "unkonw"
//	//monster12.Store()
//	if res := monster12.Restore(); res == true {
//		fmt.Println(res)
//	}
//	fmt.Println(monster12)
//
//}

type Monster struct {
	Name  string
	Age   int
	Skill string
}

const savedPath = "./json.txt"

func (this *Monster) Store() (res bool) {
	var err error
	file, err := os.OpenFile(savedPath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("读取文件错误 = ", err)
		return false
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件关闭错误 = ", err)
			res = false
		} else {
			fmt.Println("关闭文件成功")
		}
	}()
	data, err := json.Marshal(this)
	//fmt.Println(string(data))
	writer := bufio.NewWriter(file)
	if numsBytes, err := writer.WriteString(string(data)); err != nil {
		fmt.Println("写入错误 = ", err)
		return false
	} else {
		fmt.Println("写入字节数 = ", numsBytes)
		fmt.Println("写入成功")
	}
	func() {
		if err = writer.Flush(); err != nil {
			fmt.Println("Flush()错误 = ", err)
			res = false
		} else {
			fmt.Println("Frush() Success")
			res = true
		}
	}()
	return
}

// Restore 给Monster绑定方法ReStore，可以将一个序列化的Monster,从文件中读取，并反序列化为Monster对象
func (this *Monster) Restore() (res bool) {
	file, err := os.OpenFile(savedPath, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("打开文件错误 err = ", err)
		res = false
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件关闭错误 = ", err)
			res = false
			return
		} else {
			fmt.Println("关闭文件成功")
		}
	}()
	reader := bufio.NewReader(file)
	var str string = ""
	for {
		var str1 string
		str1, err = reader.ReadString('\n')
		str += str1
		fmt.Println("str1 = ", str1)
		if err == io.EOF {
			fmt.Println("到达文件结尾")
			break
		}
	}
	fmt.Println("str = ", str)
	var tempMonster Monster
	if err = json.Unmarshal([]byte(str), &tempMonster); err != nil {
		fmt.Println("反序列化失败, err = ", err)
		res = false
		return
	} else {
		res = true
	}
	this.Name = tempMonster.Name
	this.Age = tempMonster.Age
	this.Skill = tempMonster.Skill
	return
}
