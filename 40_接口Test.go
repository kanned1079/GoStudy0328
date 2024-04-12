package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//需要實現以下方法
// type Interface interface {
//    // Len方法返回集合中的元素个数
//    Len() int
//    // Less方法报告索引i的元素是否比索引j的元素小
//    Less(i, j int) bool
//    // Swap方法交换索引i和j的两个元素
//    Swap(i, j int)
//}

// 实现对Hero结构体切片的排序: sort.Sort(data Interface)
// 1.定義Hero結構體

type Hero struct {
	Name string
	Age  int
}

type Student_40 struct {
	Name  string
	Age   int
	Score float64
}

// 2.聲明一個Hero結構體的切片類型
type HerosSlice []Hero

type Students_40 []Student_40

// 3.實現Interface接口 就是三個方法
func (hs HerosSlice) Len() int {
	return len(hs)
}

// 按照Heros的年齡從小到大排序
func (hs HerosSlice) Less(i, j int) bool {
	// Less()方法決定按照使用升序還是降序排序
	//return hs[i].Age < hs[j].Age
	// 这里可以修改为按照Name进行排序
	return hs[i].Name < hs[j].Name
}

func (hs HerosSlice) Swap(i, j int) {
	// Less()方法決定按照使用升序還是降序排序
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	// 先定定義一個數組/切片
	//var intSlice = []int{0, -1, 6, 12, 48, -17}
	//// 1. 冒泡排序
	//// 2. 使用系統自帶的包
	//sort.Ints(intSlice)
	//fmt.Println(intSlice)
	// 3. 對結構體切片進行排序
	//var heros HerosSlice
	//for i := 0; i < 10; i++ {
	//	hero := Hero{
	//		Name: fmt.Sprintf("Hero-%d", rand.Intn(100)),
	//		Age:  rand.Intn(100),
	//	}
	//	heros = append(heros, hero)
	//}
	//// 排序前
	//fmt.Println("排序前")
	//for _, v := range heros {
	//	fmt.Println(v)
	//}
	//// 调用sort.Sort()
	//// 上面要是有一个方法没有实现 下面就会报错
	//sort.Sort(heros)
	//
	//// 排序后
	//fmt.Println("排序后")
	//for _, v := range heros {
	//	fmt.Println(v)
	//}

	var stus Students_40
	for i := 0; i < 40; i++ {
		st := Student_40{
			Name:  fmt.Sprintf("xxx-%d", rand.Intn(1000)),
			Age:   18,
			Score: float64(rand.Intn(100)),
		}
		stus = append(stus, st)
	}
	fmt.Println("排序前")
	for _, v := range stus {
		fmt.Println(v)
	}
	sort.Stable(stus)
	fmt.Println("排序后")

	for _, v := range stus {
		fmt.Println(v)
	}
	fmt.Println("isSorted() = ", sort.IsSorted(stus))

}

func (st Students_40) Len() int {
	return len(st)
}

func (st Students_40) Less(i, j int) bool {
	return st[i].Score < st[j].Score
}

func (st Students_40) Swap(i, j int) {
	st[i], st[j] = st[j], st[i]
}
