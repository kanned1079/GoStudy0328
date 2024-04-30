package main

import (
	"fmt"
	"time"
)

// 引入的包是 time

func main() {
	fmt.Printf("Time Now = %v, Type = %T\n", time.Now(), time.Now())
	// Time返回的是一個結構體 包含字段和方法

	now := time.Now()
	fmt.Printf("Year = %v\n", now.Year())
	fmt.Printf("Month = %v\n", int(now.Month()))
	fmt.Printf("Day = %v\n", now.Day())
	fmt.Printf("Hour = %v\n", now.Hour())
	fmt.Printf("Min = %v\n", now.Minute())
	fmt.Printf("Second = %v\n", now.Second())

	// 格式化日期事件
	// 第一種方法
	fmt.Printf("當前年月日時間 = %02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	dataStr := fmt.Sprintf("當前年月日時間 = %02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dataStr = ", dataStr)

	// 第二種方法
	fmt.Println("time = ", now.Format("2006/01/02 15:04:05")) // 這裏只能這樣寫
	fmt.Println("time = ", now.Format("2006/01/02"))          // 這裏只能這樣寫
	fmt.Println("time = ", now.Format("15:04:05"))            // 這裏只能這樣寫
	fmt.Println("month = ", now.Format("01"))                 // 只要月份
	fmt.Println("year = ", now.Format("2006"))                // 只要年份

	// 時間的常量
	// 得到100ms  100 * time.Millisecond

	// 結合sleep使用時間常量

	var a int = 0
	for {
		a++
		fmt.Printf("%02d ", a)
		if a%50 == 0 {
			fmt.Println()
		}
		if a == 100 {
			break
		}
		time.Sleep(10 * time.Millisecond) // 前面的必須傳入一個整數
	}

	// 獲取當前Unix時間戳 作用可以是生成隨機數
	fmt.Println("Unix時間戳 = ", now.Unix())
	fmt.Println("UnixNano時間戳 = ", now.UnixNano())

}
