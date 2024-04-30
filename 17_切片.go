package main

import (
	"fmt"
)

// 創建slice有三種方式
//  1. 通過一個已經創建的數組
//  2. 使用make([切片類型], [len], [cap]) 默認值是0 可通過下標來訪問
//  3. 直接定義並指定值

// 方式1和方式2的区别(面试)
//  方式1是直接引用数组，这个数组是事先存在的，程序员是可见的。
//  方式2是通过make来创建切片，make也会创建一个数组，是由切片在底层进行维护程序员是看不见的。

func main() {
	// 1. Way1 通過一個已經創建的數組
	// slice是引用傳遞
	var intArr1 [5]int = [...]int{11, 22, 33, 44, 55}
	fmt.Println("intArr1 = ", intArr1)

	slice1 := intArr1[1:4]
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice1的元素个数 = ", len(slice1))
	fmt.Println("slice1的容量 = ", cap(slice1)) // 切片的容量是動態可變的

	fmt.Printf("slice[0] = %v, Addr = %p\n", slice1[0], &slice1[0])
	fmt.Printf("intArr[1] = %v, Addr = %p\n", intArr1[1], &intArr1[1])
	fmt.Printf("slice1自身的地址 = %p\n\n", &slice1)

	// 2. Way2 使用make([切片類型], [len], [cap]) 默認值是0 可通過下標來訪問
	var slice2 []float64
	// slice必須make後才能使用
	// 使用make創建的切片 只能通過slice訪問 所以對外不可見
	slice2 = make([]float64, 5, 10)
	fmt.Println("slice2 = ", slice2)
	fmt.Printf("slice2's len = %d, cap = %d\n\n", len(slice2), cap(slice2))

	// 3. Way3 直接定義並指定值
	var slice3 []string = []string{"Tom", "Jerry", "Mary"}
	fmt.Println("slice3 = ", slice3)
	fmt.Printf("slice3's len = %d, cap = %d\n\n", len(slice3), cap(slice3))

	// 切片的遍歷
	// 1.正常for循環
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice[%d] = %d\t ", i, slice1[i])
	}
	fmt.Println()

	// 2. range遍歷
	for _, i := range slice3 {
		fmt.Printf(i + " ")
	}
	fmt.Println("\n")

	// 切片初始化时，仍然不能越界。范围在 [0-len(arr)]之间，但是可以动态增长
	// var slice =arr[0:end] 可以简写: var slice = arr[:end]
	// var slice =arr[start:len(arr)] 可以简写: var slice = arr[start:]
	// var slice = arr[0:len(arr)] 可以简写: var slice = arr[:]

	// 後面可以寫多個
	slice2 = append(slice2, 200, 300, 400)
	fmt.Println("slice2 = ", slice2)
	fmt.Printf("slice2's len = %d, cap = %d\n\n", len(slice2), cap(slice2))

	slice4 := []int{400, 500, 600}
	slice1 = append(slice1, slice4...) // 追加一倍
	// 這裡append在底層創建了新的newArr并操作
	// 程序員看不見
	fmt.Println("slice1 = ", slice1)
	fmt.Printf("slice1's len = %d, cap = %d\n\n", len(slice1), cap(slice1))

	for i := 0; i < 100; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Println("slice1的元素个数 = ", len(slice1))
	fmt.Println("slice1的容量 = ", cap(slice1)) // 切片的容量是動態可變的

	var slice11 []int = []int{1, 2, 3, 4, 5}
	fmt.Println("slice11 = ", slice11) // [1 2 3 4 5]
	var slice12 []int = []int{20, 40, 60}
	// 拷貝slice12到slice11並賦改
	// 要求拷貝的數據必須都是切片
	copy(slice11, slice12)
	fmt.Println("slice11 = ", slice11) // [20 40 60 4 5]
	// 數據是獨立的空間 修改slice12並不會修改到slice11

	test17_1()

	var slice6 []int = []int{1, 2, 3, 4, 5}
	fmt.Println("slice6 = ", slice6)
	change17_2(slice6) // 這裏因為是引用傳遞 所以值會被修改
	fmt.Println("slice6 = ", slice6)

	fmt.Println("fbn的切片 = ", fbn(20))

}

func test17_1() {
	var a []int = []int{1, 2, 3, 4, 5}
	var slice = make([]int, 1)
	fmt.Println(slice) // []0
	copy(slice, a)     // 將a拷貝到slice 雖然長度不夠 但是可以拷貝
	fmt.Println(slice) // [1]
}

func change17_2(slice []int) {
	slice[0] = 100
}

func fbn(n int) (res []uint64) {
	fbnSlice := make([]uint64, n)
	// 第一個和第二個都為1
	fbnSlice[0], fbnSlice[1] = 1, 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	res = fbnSlice
	return
}
