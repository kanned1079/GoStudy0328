package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nums [10]int
	var targerNum int = 89
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println("nums = ", nums)

	var index int = -1
	for i := 0; i < len(nums); i++ {
		if targerNum == nums[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("found, index = %d\n", index)
	} else {
		fmt.Println("not found")
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	fmt.Println("Sorted nums = ", nums)
	var findVal int = 34
	binarySearch(&nums, 0, len(nums)-1, findVal)

	//fmt.Printf("位置 = ", seqSearch(nums, 89))
}

func seqSearch(arr []int, target int) (position int) {
	for i := 0; i < len(arr); i++ {
		if target == arr[i] {
			position = i
			return
		}
	}
	return -1
}

func binarySearch(arr *[10]int, leftIndex, rightIndex, findValue int) {
	// 判斷leftIndex是否大於rightIndex
	if leftIndex > rightIndex {
		fmt.Println("Not Found")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findValue {
		// 說明要查找的數 應該在 leftIndex --- middle - 1
		binarySearch(arr, leftIndex, middle-1, findValue)
	} else if (*arr)[middle] < findValue {
		// 說明要查找的數 應該在 middle+1 --- rightIndex
		binarySearch(arr, middle+1, rightIndex, findValue)
	} else {
		fmt.Println("Found val = ", middle)
	}
}
