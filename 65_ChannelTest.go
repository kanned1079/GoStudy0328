package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

// 作业2： 使用goroutine + channel 配合完成排序并写入文件
// 开一个协程writeDataToFile 随机写入1000个数据 存放到文件中
// 当writeDataToFile完成后 让sort协程从文件中读取1000个数据
// 并完成排序 重新写入到另一个文件

func writeDataToFile(i int, saved chan bool) {
	startTime := time.Now()
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
			//res = true
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
	//return
	endTime := time.Now()
	fmt.Printf("插入文件[%d]用时%v\n", i, endTime.Sub(startTime))
}

func readAndSort(i int, sortedChan chan bool) {
	startTime := time.Now()
	var err error
	file, err := os.OpenFile(fmt.Sprintf("num%d.txt", i), os.O_RDONLY, 0444)
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
	var nums []int
	//var term int
	for {
		var str string
		str, err = reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("到达文件末尾")
			break
		}
		str = str[:len(str)-1]

		num, err = strconv.Atoi(str)
		//fmt.Println(num)
		if err != nil {
			fmt.Println("转换失败 err = ", err)
		} else {
			//nums[term] = num
			//term++
			nums = append(nums, num)
		}
	}
	// 这里开始排序
	sort.Ints(nums)
	//fmt.Println(nums)
	//sortedChan <- true
	// 这里开始写入文件
	saveFile, err := os.OpenFile(fmt.Sprintf("num%d_sorted.txt", i), os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("文件2打开失败 err =", err)
	} else {
		fmt.Println("文件2打开成功")
	}
	defer func() {
		if err = saveFile.Close(); err != nil {
			fmt.Println("文件2关闭失败 err =", err)
		} else {
			fmt.Println("文件2关闭成功")
			sortedChan <- true
		}
	}()
	writer := bufio.NewWriter(saveFile)
	for i := 0; i < len(nums); i++ {
		if _, err = writer.WriteString(fmt.Sprintf("%d\n", nums[i])); err != nil {
			fmt.Println("文件2写入错误 err = ", err)
		}
	}
	func() {
		if err = writer.Flush(); err != nil {
			fmt.Println("文件2刷新失败 err =", err)
		} else {
			fmt.Println("文件2刷新成功")
		}
	}()
	endTime := time.Now()
	fmt.Printf("保存文件[%d]用时%v\n", i, endTime.Sub(startTime))
}

func main() {
	//var numChan chan int = make(chan int, 1000)
	var savedChan = make(chan bool, 10)
	var sortedChan = make(chan bool, 10)
	//writeDataToFile(1, savedChan)
	for i := 1; i <= 10; i++ {
		go writeDataToFile(i, savedChan)
	}
	var done [10]bool
	var a int
	for i := 1; i <= 10; i++ {
		done[a] = <-savedChan // 等待接收完成信号
		a++
	}
	fmt.Println(done)
	fmt.Println("--------------------------")

	for i := 1; i <= 10; i++ {
		go readAndSort(i, sortedChan)
	}

	for i := 1; i <= 10; i++ {
		<-sortedChan
	}

}
