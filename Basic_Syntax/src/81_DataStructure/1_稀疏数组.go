package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type valNode1 struct {
	row int
	col int
	val int
}

func main() {
	file, err := os.OpenFile("./src/81_DataStructure/saves/chessMap.data", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	// 1.创建一个原始数组
	var chess [11][11]int
	chess[1][2] = 1 // 黑子
	chess[2][3] = 2 // 蓝子
	chess[6][5] = 1 // 蓝子
	chess[8][0] = 2 // 蓝子
	chess[7][8] = 2 // 蓝子

	// 2.遍历
	for _, v := range chess {
		for _, v2 := range v {
			fmt.Printf("%d  ", v2)
		}
		fmt.Println()
	}

	// 3.转换为一个稀疏数组
	// 遍历chess 如果发现有一个元素的值不为0 就创建一个node结构体放置到切片中
	var sparseArr []valNode1
	// 标准的稀疏数字在第一行有记录二维数组的规模和默认值
	sparseArr = append(sparseArr, valNode1{11, 11, 0})
	for i, v := range chess {
		for j, v2 := range v {
			if v2 != 0 {
				sparseArr = append(sparseArr, valNode1{i, j, v2})
			}
		}
	}

	// 输出稀疏数组
	var buf string
	for i, v := range sparseArr {
		log.Printf("%d > %d %d %v\n", i, v.row, v.col, v.val)
		buf = strconv.Itoa(v.row) + " " + strconv.Itoa(v.col) + " " + strconv.Itoa(v.val)
		save1(buf, writer)
	}

	log.Println("Starting Read func.")
	read1()

}

// save1 将稀疏数组存入文件
func save1(str string, writer *bufio.Writer) {
	nums, err := writer.WriteString(str + "\n")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Written: ", nums)
	defer writer.Flush()
}

// read1 从文件读取稀疏数组并恢复
func read1() {
	file, _ := os.OpenFile("./src/81_DataStructure/saves/chessMap.data", os.O_RDONLY, 0400)
	// 先创建
	var newChess [11][11]int

	// 读取文件
	reader := bufio.NewReader(file)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Println(err)
			break
		}
		line = line[:len(line)-1]
		strSlice := strings.Split(line, " ")
		if i > 0 { // 掠过第一行
			cow, _ := strconv.Atoi(strSlice[0])
			col, _ := strconv.Atoi(strSlice[1])
			val, _ := strconv.Atoi(strSlice[2])
			newChess[cow][col] = val
		}
		i++
	}

	// 2.遍历
	fmt.Println("恢复的数组: ")
	for _, v := range newChess {
		for _, v2 := range v {
			fmt.Printf("%d  ", v2)
		}
		fmt.Println()
	}

}
