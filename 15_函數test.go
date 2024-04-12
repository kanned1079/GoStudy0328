package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getDays(year, month, days int) (flag bool) {
	var isLYear bool = (year%400 == 0) || (year%4 == 0 && year%100 != 0)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		{
			if days > 0 && days <= 31 {
				flag = true
				return
			} else {
				flag = false
				return
			}
		}
	case 4, 6, 9, 11:
		{
			if days > 0 && days <= 30 {
				flag = true
				return
			} else {
				flag = false
				return
			}
		}
	case 2:
		{
			if isLYear == true {
				if days > 0 && days <= 29 {
					flag = true
					return
				} else {
					flag = flag
					return
				}
			} else {
				if days > 0 && days <= 28 {
					flag = true
					return
				} else {
					flag = flag
					return
				}
			}
		}

	}
	return false
}

func main() {
	//var year, month, day int
	//fmt.Print("输入年份：")
	//fmt.Scan(&year)
	//fmt.Print("输入月份：")
	//fmt.Scan(&month)
	//fmt.Print("输入天数：")
	//fmt.Scan(&day)
	//
	//flag := getDays(year, month, day)
	//if flag == true {
	//	fmt.Printf("输入正确， 你输的是 %d年%d月%d日\n", year, month, day)
	//} else {
	//	fmt.Printf("输入不正确")
	//}
	guessNumber()

}

func guessNumber() {
	rand.Seed(time.Now().UnixNano())
	var targetNumber int = rand.Intn(100) + 1
	fmt.Println("randNum = ", targetNumber)
	var inpNum int = 0
	var nums int
	for nums = 1; nums <= 10; nums++ {
		fmt.Printf("You have %d chances left\t", 11-nums)
		fmt.Scan(&inpNum)
		if startGuess(targetNumber, inpNum) == true {
			switch nums {
			case 1:
				{
					fmt.Println("你猜中了 真是个天才")
				}
			case 2, 3:
				{
					fmt.Println("你真聪明 赶上我了")
				}
			case 4, 5, 6, 7, 8, 9:
				{
					fmt.Println("一般般")
					break
				}
			case 10:
				{
					fmt.Println("可算猜对了")

				}

			}
		} else {
			fmt.Println("没有猜中")
		}
	}
	if nums == 10 && !startGuess(targetNumber, inpNum) {
		fmt.Println("说你啥好呢")
	}

}

func startGuess(targetNum, inpNum int) (flag bool) {
	if targetNum == inpNum {
		flag = true
		return
	} else {
		flag = false
		return
	}
}
