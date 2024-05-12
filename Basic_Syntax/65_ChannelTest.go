package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

// 作业2： 使用goroutine + channel 配合完成排序并写入文件
// 开一个协程writeDataToFile 随机写入1000个数据 存放到文件中
// 当writeDataToFile完成后 让sort协程从文件中读取1000个数据
// 并完成排序 重新写入到另一个文件
const path = "./num.txt"

func writeDataToFile(i int, saved chan bool) {
	var err error
	//var file *os.File
	file, err := os.OpenFile(fmt.Sprintf("num%d.txt", i), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println("文件打开失败 err = ", err)
	} else {
		fmt.Println("文件打开成功")
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件关闭失败 err = ", err)
		} else {
			fmt.Println("文件关闭成功")
			saved <- true
		}
	}()
	writer := bufio.NewWriter(file)
	for i := 0; i < 1000; i++ {
		_, err = writer.WriteString(fmt.Sprintf("%d\n", rand.Intn(1000)))
		if err != nil {
			fmt.Println("写入错误 err = ", err)
		}
	}
	if err = writer.Flush(); err != nil {
		fmt.Println("文件刷新失败")
	} else {
		fmt.Println("文件刷新成功")
	}
}

func read(numChan chan int) {
	var err error
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("文件打开失败 err = ", err)
	} else {
		fmt.Println("文件打开成功")
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件关闭失败 err = ", err)
		} else {
			fmt.Println("文件关闭成功")
		}
	}()
	reader := bufio.NewReader(file)
	var num int
	for {
		var str string
		str, err = reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("到达文件末尾")
			break
		}
		str = str[:len(str)-1]

		num, err = strconv.Atoi(str)
		fmt.Println(num)
		if err != nil {
			fmt.Println("转换失败 err = ", err)
		}
		numChan <- num
	}
	close(numChan)
}

func sortNumbers(numChan chan int, sortedNums chan int) {

	var nums []int
	for num := range numChan {
		nums = append(nums, num)
	}
	sort.Ints(nums)

	file, err := os.OpenFile("./numSorted.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println("文件打开失败 err = ", err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件关闭失败 err = ", err)
		}
	}()

	writer := bufio.NewWriter(file)
	for _, num := range nums {
		_, err = writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			fmt.Println("写入错误 err = ", err)
		}
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("文件刷新失败")
	} else {
		fmt.Println("文件刷新成功")
	}
}

func allFilesSaved(savedChan chan bool) {
	for i := 1; i <= 10; i++ {
		<-savedChan
	}
	close(savedChan)
}

func main() {
	var savedChan chan bool = make(chan bool, 10)
	//var numChan chan int = make(chan int, 1000)
	for i := 1; i <= 10; i++ {
		go writeDataToFile(i, savedChan)
	}
	close(savedChan)

	//fmt.Println(len(savedChan))

	// 创建一个通道来告知主协程所有的文件都已经写入完毕
	allFilesSavedChan := make(chan bool)
	go allFilesSaved(allFilesSavedChan)

	// 等待所有文件写入完成
	<-allFilesSavedChan

	// 启动 10 个排序协程
	for i := 1; i <= 10; i++ {
		go func(filename string) {
			numChan := make(chan int)
			sortedNums := make(chan int)
			defer close(sortedNums)

			// 从文件中读取数据
			go read(numChan)

			// 完成排序
			go sortNumbers(numChan, sortedNums)

			// 可以在这里读取 sortedNums 通道中的数据，或者直接在 sortNumbers 函数中写入文件
		}(fmt.Sprintf("num%d.txt", i))
	}

	// 功能扩展: 开 10个协程 writeDataToFile，每个协程随机生成 1000 个数据，存放到10个文件中
	// 当10个文件都生成了，让 10个sort协程从10个文件中读取 1000 个数据，并完成排序,重新写入到 10 个结果文件

}
