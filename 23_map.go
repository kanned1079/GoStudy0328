package main

import (
	"fmt"
)

func main() {
	// 這是第一種方式創建map
	var a map[string]string
	// 在使用map前 需要先make分配空間
	a = make(map[string]string, 10)
	a["N1"] = "kanna"
	a["N2"] = "kinggyo"
	a["N3"] = "izumi"
	a["N1"] = "kanna2" // 這裏會給前面覆蓋掉
	// key不可以重複 但是value可以
	// Golang中的map是無序的
	fmt.Println(a)

	// 第二種方式 聲明的時候make
	var citits = make(map[int]string, 10)
	citits[1] = "Shanghai"
	citits[2] = "HongKong"
	fmt.Println(citits)

	// 在聲明的時候 直接賦值
	var provinces map[int]string = map[int]string{1: "JiangSu", 2: "ChongQing", 3: "YunNan"}
	fmt.Println(provinces)

	var stuInfos = make(map[string]map[string]string, 10)
	stuInfos["stu01"] = make(map[string]string, 2) // 有map就要make
	stuInfos["stu01"]["name"] = "kanna"
	stuInfos["stu01"]["sex"] = "M"
	stuInfos["stu01"]["addr"] = "Shanghai"

	stuInfos["stu02"] = make(map[string]string, 2) // 有map就要make
	stuInfos["stu02"]["name"] = "kinggyo"
	stuInfos["stu02"]["sex"] = "F"
	stuInfos["stu02"]["addr"] = "HongKong"

	//fmt.Println(stuInfos)
	// 分開輸出
	fmt.Println("stu01 = ", stuInfos["stu01"])
	fmt.Println("stu02 = ", stuInfos["stu02"])

	fmt.Println("stu01's addr = ", stuInfos["stu01"]["addr"])
	fmt.Println("stu01's sex = ", stuInfos["stu02"]["sex"])

	// map增刪改查
	// 直接增加
	citits[4] = "Beijing"
	citits[5] = "Changzhou"
	citits[6] = "Xuzhou"
	fmt.Println(citits)
	// 將Changzhou改為常州
	// 因為cities[6]已經存在 那麼就相當於修改了
	citits[5] = "常州"
	fmt.Println(citits)
	// 刪除語法 func delete(m map[Type]Type1, key Type)
	// delete是一个内置函数，如果key存在，就删除该key-value如果key不存在，不操作，但是也不会报错
	delete(citits, 6)  // 這裏刪除了6 Xuzhou
	delete(citits, 10) // 如果找不到的key那麼也能運行而且不會報錯
	fmt.Println(citits)

	// 如果刪除一整個key
	fmt.Println("provinces = ", provinces)
	fmt.Println("刪除map後 ")
	// 1. 遍歷map 一個個刪除
	// 2. map = make(...) make一個新的 舊的會被GC回收成為垃圾
	//for i := 0; i < len(citits); i++ {
	//	delete(citits, i)
	//}
	provinces = make(map[int]string)
	//fmt.Println(citits)
	fmt.Println("provinces = ", provinces)

	// map的查詢 查詢一個map中有沒有指定的key
	val, ok := citits[1]
	if ok {
		fmt.Printf("有1 key， 值 = %v\n", val)
	} else {
		fmt.Println("沒有")
	}

	// 遍歷map 注意只能使用for range
	for key, value := range citits {
		fmt.Printf("key[%v] = %v\n", key, value)
	}
	fmt.Println()
	// 遍歷一個結構比較複雜的map
	for key1, val1 := range stuInfos {
		fmt.Println("key1 = ", key1)
		for key2, val2 := range val1 {
			fmt.Printf("key2[%v]  = %v\n", key2, val2)
		}
	}

	// 返回map的长度
	fmt.Printf("len(stuInfo2) = %v\nlen(cities) = %v\nlen(province) = %v\n", len(stuInfos), len(citits), len(provinces))

	// 使用clear()方法删除map
	clear(stuInfos)
	fmt.Println("stuInfo = ", stuInfos)
}
