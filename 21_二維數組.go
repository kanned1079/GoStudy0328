package main

import "fmt"

func main() {
	var arr1 [4][6]int // 先定义
	arr1[1][2] = 1     // 再賦值

	fmt.Println("arr1 = ", arr1)
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("%v ", arr1[i][j])
		}
		fmt.Println()
	}

	// 直接聲明並賦值
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr2 = ", arr2)

	fmt.Println()
	// 两种二维数组遍历方法
	// 1. for循环
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Printf("%v ", arr2[i][j])
		}
		fmt.Println()
	}

	// 2. for...range
	for i, v := range arr2 {
		//fmt.Printf( v) // 這裏的v都是一個個一位數組
		for j, v2 := range v {
			fmt.Printf("arr[%d][%d] = %d\t", i, j, v2)
		}
		fmt.Println()
	}
}
