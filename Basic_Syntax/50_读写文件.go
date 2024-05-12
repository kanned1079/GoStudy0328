package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
)

// func OpenFile(name string, flag int perm FileMode) (file *File, err error)
// os.0penFile是一个更一般性的文件打开函数
// 它会使用指定的选项(如O RDONLY等)
// 指定的模式(如0666等)打开指定名称的文件
// 如果操作成功，返回的文件对象可用于I/0
// 如果出错，错误底层类型是*PathError
// 文件权限 r -> 4  w -> 2  x -> 1

const (
	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)

type FileMode uint32 // 这个是对于类UNIX类 在Windows下无效
// FileMode代表文件的模式和权限位。
// 这些字位在所有的操作系统都有相同的含义，因此文件的信息可以在不同的操作系统之间安全的移植。
// 不是所有的位都能用于所有的系统，唯一共有的是用于表示目录的ModeDir位。
const (
	// 单字符是被String方法用于格式化的属性缩写。
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
	ModeAppend                                     // a: 只能写入，且只能写入到末尾
	ModeExclusive                                  // l: 用于执行
	ModeTemporary                                  // T: 临时文件（非备份文件）
	ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
	ModeDevice                                     // D: 设备
	ModeNamedPipe                                  // p: 命名管道（FIFO）
	ModeSocket                                     // S: Unix域socket
	ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
	ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
	ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
	ModeSticky                                     // t: 只有root/创建者能删除/移动文件
	// 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
	ModeType          = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
	ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
)

const (
	filePath2 string = "./Files/a.txt"
)

func main() {
	start := time.Now()
	//test1()
	//test2()
	//test3()
	//test4()
	test5()

	end := time.Now()
	spendTime := end.Sub(start)
	fmt.Println("时间 = ", spendTime)
}

// test1 创建一个新文件写入内容 5句"hello, Gardon"
func test1() {
	// 参数为 文件路径 打开的方式 权限控制
	//file1, err := os.OpenFile(filePath, syscall.O_RDWR, 0777)

	// 创建一个新文件写入内容 5句"hello, Gardon"
	filePath2 := "./Files/b.txt"
	file1, err := os.OpenFile(filePath2, O_WRONLY|O_CREATE, 0777)
	if err != nil {
		fmt.Println("错误 = ", err)
		return
	}
	str1 := "Hello, Gardon"
	// 写入时 使用带缓存的 *Writer
	writer1 := bufio.NewWriter(file1)
	for i := 0; i < 5; i++ {
		writer1.WriteString(str1 + "\n") // 加一个换行
	}
	// 因为Writer是带有缓存的 因此在调用Write时候 内容是先写入缓存(buff)还没有被写入磁盘
	// 还需要调用一下Flush() 才会真正写入到文件中
	writer1.Flush()

	defer file1.Close()
}

// test2 打开一个存在的文件中，将原来的内容覆盖成新的内容10句"你好，尚硅谷!"
func test2() {
	filePath2 := "./Files/b.txt"
	file1, err := os.OpenFile(filePath2, O_WRONLY|O_TRUNC, 0777) // O_TRUNC 在每次打开文件时清空文件
	if err != nil {
		fmt.Println("错误 = ", err)
	}
	defer file1.Close()
	str1 := "你好，尚硅谷"
	writer1 := bufio.NewWriter(file1)
	var writtenNum int
	var err2 error
	for i := 0; i < 10; i++ {
		writtenNum, err2 = writer1.WriteString(str1 + "\n")
		if err2 != nil {
			fmt.Println("错误2 = ", err2)
			break
		}
	}
	fmt.Println("写入的字节数 = ", writtenNum)
	if err3 := writer1.Flush(); err3 != nil {
		fmt.Println("保存失败， 错误3 = ", err3)
	}
	fmt.Println("写入完成")
}

// test3 打开一个存在的文件，在原来的内容追加内容'ABC! ENGLISH!'
func test3() {
	filePath2 := "./Files/b.txt"
	file1, err := os.OpenFile(filePath2, O_WRONLY|O_APPEND, 0777) // O_TRUNC 在每次打开文件时清空文件
	if err != nil {
		fmt.Println("错误 = ", err)
	}
	defer func() {
		if err := file1.Close(); err != nil {
			fmt.Println("关闭文件出错，错误 = ", err)
		} else {
			fmt.Println("关闭文件成功")
		}
	}()
	str1 := "ABC! ENGLISH!"
	writer1 := bufio.NewWriter(file1)
	var writtenNum int
	var err2 error
	for i := 0; i < 10; i++ {
		writtenNum, err2 = writer1.WriteString(str1 + "\n")
		if err2 != nil {
			fmt.Println("错误2 = ", err2)
			break
		}
	}
	fmt.Println("写入的字节数 = ", writtenNum)
	if err3 := writer1.Flush(); err3 != nil {
		fmt.Println("保存失败， 错误3 = ", err3)
	}
	fmt.Println("写入完成")
}

// test4 打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,北京!"
func test4() {
	filePath2 := "./Files/b.txt"
	file1, err := os.OpenFile(filePath2, O_RDWR|O_APPEND, 0777) // O_TRUNC 在每次打开文件时清空文件
	if err != nil {
		fmt.Println("错误 = ", err)
	}
	defer func() {
		if err := file1.Close(); err != nil {
			fmt.Println("关闭文件出错，错误 = ", err)
		} else {
			fmt.Println("关闭文件成功")
		}
	}()
	reader1 := bufio.NewReader(file1)
	for {
		str, err := reader1.ReadString('\n')
		if err == io.EOF {
			fmt.Println("到了文件末尾")
			break
		}
		fmt.Print(str)
	}

	str1 := "ABC! ENGLISH!"
	writer1 := bufio.NewWriter(file1)
	var writtenNum int
	var err2 error
	for i := 0; i < 10; i++ {
		writtenNum, err2 = writer1.WriteString(str1 + "\n")
		if err2 != nil {
			fmt.Println("错误2 = ", err2)
			break
		}
	}
	fmt.Println("写入的字节数 = ", writtenNum)
	if err3 := writer1.Flush(); err3 != nil {
		fmt.Println("保存失败， 错误3 = ", err3)
	}
	fmt.Println("写入完成")
}

// 编写个程序，将一个文件的内容，写入到另外一个文件
// 注:这两个文件已经存
func test5() {
	sourceFile, targetFile := "./Files/a.txt", "./Files/a2.txt"
	var err error
	file, err := os.OpenFile(sourceFile, O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件错误")
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件关闭失败 = ", err)

		} else {
			fmt.Println("Success")
		}
	}()
	var str string
	reader := bufio.NewReader(file)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件结束")
			break
		}
		str += content
	}
	//fmt.Println(str)
	file2, err := os.OpenFile(targetFile, O_RDWR|O_CREATE, 0777)
	if err != nil {
		fmt.Println("文件2打开失败 = ", err)
	}
	defer func() {
		if err = file2.Close(); err != nil {
			fmt.Println("文件file2关闭失败 = ", err)

		} else {
			fmt.Println("Success")
		}
	}()
	writer := bufio.NewWriter(file2)
	writtenBytes, err := writer.WriteString(str)
	fmt.Println("写入字节数 = ", writtenBytes)
	func() {
		if err = writer.Flush(); err != nil {
			fmt.Println("写入出错 = ", err)
		} else {
			fmt.Println("Success")
		}
	}()

}
