package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//test22_1()
	//test22_2()
	//test22_3()
	//test22_4()
	//test22_5()
	//test22_06()
	//test22_07()

	//var a [5]int = [5]int{19, 6, 43, 634, 9}
	//maxIndex, minIndex := test22_8(a)
	//fmt.Printf("最大值[%d]， 最小值[%d]\n", maxIndex, minIndex)

	test22_9()

}

// 随机生成 10 个整数(1 100 的范围)保存到数组，并倒序打印以及求平均值、求最大值和最大值的下标、并查找里面是否有 55
func test22_1() {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(101)
	}
	fmt.Println("生成的数组num = ", nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("逆序后的数组num = ", nums)
	max, min, maxIndex, minIndex, all := nums[0], nums[0], 0, 0, 0
	isExist := false
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
		if nums[i] < min {
			min = nums[i]
			minIndex = 1
		}
		all += nums[i]
		if nums[i] == 55 {
			isExist = true
		}
	}
	var avgVal float64
	avgVal = float64(all) / float64(len(nums))
	fmt.Printf("max[%d] = %d, min[%d] = %d\n", maxIndex, max, minIndex, min)
	fmt.Println("avgVal = ", avgVal)
	sort.Ints(nums)
	if isExist {
		fmt.Println("存在55")
	} else {
		fmt.Println("不存在")
	}
}

// 已知有个排序好(升序)的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
func test22_2() {
	sortedArray := []int{10, 20, 30, 40, 50}
	newElement := 35
	sortedArray = append(sortedArray, newElement)
	fmt.Println(sortedArray)
	sort.Ints(sortedArray)
	fmt.Println(sortedArray)
}

// 定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数据清0
func test22_3() {
	var arr1 [3][4]int = [3][4]int{}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("val[%d][%d]: ", i, j)
			fmt.Scan(&arr1[i][j])
		}
	}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			if i == 0 || j == 0 {
				arr1[i][j] = 0
			}
			if i == len(arr1)-1 || j == len(arr1)-1 {
				arr1[i][j] = 0
			}
		}
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("%d ", arr1[i][j])
		}
		fmt.Println()
	}
}

// 定义一个4行4列的二维数组,逐个从键盘输入值,
// 然后将第1行和第4行的数据进行交换将第2行和第3行的数据进行交换
func test22_4() {
	var a int = 1
	var numsArr [4][4]int
	for i := 0; i < len(numsArr); i++ {
		for j := 0; j < len(numsArr[i]); j++ {
			numsArr[i][j] = a
			a++
		}
	}
	numsArr[0], numsArr[3] = numsArr[3], numsArr[0]
	numsArr[2], numsArr[3] = numsArr[3], numsArr[2]
	fmt.Println(numsArr)

	for _, v := range numsArr {
		for _, v2 := range v {
			fmt.Printf("%02d ", v2)
		}
		fmt.Println()
	}
}

// 试保存13579 五个奇数到数组，并倒序打印
func test22_5() {
	var arr [5]int
	a := 1
	for i := 0; i < 5; i++ {
		arr[i] = a
		a += 2
	}
	fmt.Println("arr = ", arr)
	fmt.Print("逆序输出arr = ")
	for i := 4; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}

// 试写出实现查找的核心代码，比如已知数组 arr [10]string;
// 里面保存了十个元素，现要查找"AA"在其中是否存在，打印提示
// 如果有多个”AA”，也要找到对应的下标。
func test22_06() {
	var arr [10]string = [10]string{"HH", "LI", "AA", "UI", "PY", "AA", "NA"}
	var isExist bool = false
	for i, val := range arr {
		if val == "AA" {
			fmt.Printf("Found %s in index = %d\n", val, i)
			isExist = true
		} else if i == len(arr)-1 && !isExist {
			fmt.Println("Not found")
		}
	}
}

// 随机生成 10 个整数(1-100 之间)，使用冒泡排序法进行排序，
// 然后使用二分查找法，查找是否有 90 这个数，并显示下标，如果没有则提示"找不到该数"
func test22_07() {
	rand.Seed(time.Now().UnixNano())
	var nums [10]int
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println("原始数据nums = ", nums)
	whatFind := 90
	binSrarch_Via07(&nums, 0, len(nums)-1, whatFind)
	//binarySearch02(&nums, 0, len(nums)-1, whatFind)
}

func binSrarch_Via07(arr *[10]int, leftIndex, rightIndex, findVal int) {
	if leftIndex > rightIndex {
		fmt.Printf("找不到")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if arr[middle] > findVal {
		binSrarch_Via07(arr, leftIndex, middle-1, findVal)
	} else if arr[middle] < findVal {
		binSrarch_Via07(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了%d在Index = %d\n", arr[middle], middle)
	}
}

// 编写一个函数，可以接收一个数组，该数组有5个数，
// 请找出最大的数和最小的数和对应的数组下标是多
func test22_8(arr [5]int) (maxIndex, minIndex int) {
	max, min := arr[0], arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
	}
	return
}

// 定义一个数组，并给出8个整数
// 求该数组中大于平均值的数的个数，和小于平均值的数的个数。
func test22_9() {
	var nums [8]int = [8]int{1, 78, 34, 90, 243, 65, 23, 6}
	var (
		all float64
		avg float64
	)
	for _, val := range nums {
		all += float64(val)
	}
	avg = all / float64(len(nums))
	low, high := 0, 0
	for _, val := range nums {
		if float64(val) < avg {
			low++
		} else {
			high++
		}
	}
	fmt.Printf("大於avg的有%d個， 小於avg的有%d個\n", high, low)
}
