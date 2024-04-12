package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 数组是相同类型的组合

func main() {

	//printRand()
	//
	//arr := [3]int{11, 22, 33}
	//change0(arr)
	//fmt.Println("arr = ", arr)
	//change1(&arr)
	//fmt.Println("arr = ", arr)
	qa1()

	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(101)
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
	fmt.Println("Max num = ", getMaxNum(arr))
	var arr2 [10]float32
	for i := 0; i < len(arr); i++ {
		arr2[i] = float32(arr[i])
	}
	fmt.Println("Avg num = ", getAvgNum(arr2))

	reverseNumsGroup(&arr)
	fmt.Println("逆序后的数组 = ", arr)

}

func change0(arr [3]int) {
	arr[1] = 777
}

func change1(arr *[3]int) {
	arr[0] = 777
}

func printRand() {
	rand.Seed(time.Now().UnixNano())
	//var nums0 [...]int // 这是个切片
	//fmt.Println("nums0 = ", nums0)
	var nums1 []int
	nums1 = make([]int, 20)
	for i := 0; i < 20; i++ {
		nums1[i] = rand.Intn(101)
		fmt.Printf("%d ", nums1[i])
	}
	fmt.Println()
}

func qa1() {
	var words []byte
	words = make([]byte, 26)
	for i := 0; i < 26; i++ {
		words[i] = byte('A' + i)
	}
	for _, i := range words {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
}

func getMaxNum(nums [10]int) (maxNum int) {
	maxNum = nums[0]
	for _, i := range nums {
		if i > maxNum {
			maxNum = i
		}
	}
	return maxNum
}

func getAvgNum(nums [10]float32) (avg float32) {
	for i := 0; i < len(nums); i++ {
		avg += nums[i]
	}
	avg /= float32(len(nums))
	return
}

func reverseNumsGroup(nums *[10]int) {
	//for i := 0; i < len(nums)/2; i++ {
	//	nums[i], nums[10-1-i] = nums[10-1-i], nums[i]
	//}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[len(nums)-1-i], _ = strconv.Atoi(strconv.Itoa(nums[i]))
	}

}
