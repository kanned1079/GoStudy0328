package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 獲取字符串長度 len(v) 內建函數 按字節返回的
	str1 := "hello北"
	// Golang的中文編碼統一為utf-8 字母和數字佔用1字節 漢字佔用3字節
	fmt.Println("len(str) = ", len(str1)) // 5+3

	// 字符串遍歷 同時有中文的情況
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i]) // 這樣輸出會有亂碼
	}
	fmt.Println()

	// 解決方法是 轉換為切片
	str2 := []rune(str1)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i]) // 這樣輸出會有亂碼
	}
	fmt.Println()

	// 字符串轉換為整數
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("轉換錯誤", err)
	} else {
		fmt.Println("轉換成功 n = ", n)
	}

	// 整數轉字符串
	str3 := strconv.Itoa(12324)
	fmt.Println("str3 = ", str3)

	// 字符串轉 []byte
	var byte1 = []byte("hello go")
	fmt.Printf("byte = %v ", byte1)
	fmt.Println()

	// []byte 轉換為字符串
	str4 := string([]byte{97, 98, 99, 100})
	fmt.Println("str4 = ", str4)

	// 進制轉換
	// 传入参数为([原始值], [進制])，返回一個string
	//str5 := strconv.FormatInt(123, 2)
	fmt.Printf("128轉2進制 = %s\n", strconv.FormatInt(123, 2))
	fmt.Printf("123轉16進制 = %s\n", strconv.FormatInt(123, 16))
	fmt.Printf("123轉8進制 = %s\n", strconv.FormatInt(123, 8))

	// 查找字串是否在目标串中 返回bool
	var isExist1 bool
	isExist1 = strings.Contains("seafood", "ea")
	fmt.Println("isExist = ", isExist1)

	// 统计一个目标串中有多少个不重复的子串
	fmt.Println("seq strings num = ", strings.Count("Sesssasssas", "s"))
	fmt.Println("seq strings num = ", strings.Count("Sesssasssas", "o"))

	// 字符串的比较 返回bool類型
	// 1. 使用 == 但是只有大小寫都相同時才返回 true
	fmt.Println("is equal? via == = ", "aaAAA, tg" == "aaAAA, tg")
	fmt.Println("is equal? via == = ", "aaAAA, tg" == "aaaa, tg")
	// 2. 使用strings.EqualFold() 不區分大小寫
	fmt.Println("is equal? via strings.EqualFold() = ", strings.EqualFold("aaAAA, tg", "aaAAA, tg"))
	fmt.Println("is equal? via strings.EqualFold() = ", strings.EqualFold("aaAAA, tg", "AAaaa, TG"))

	// 返回字串第一次出現在目標串中的下標 沒有就返回-1 strings.Index()
	index := strings.Index("NLT_abc", "ab")
	fmt.Println("index = ", index)
	fmt.Println("index2 = ", strings.Index("NLT_abc", "hello"))

	// 返回字串最後一次出現在目標串中的下標 沒有就返回-1 strings.LastIndex()
	index = strings.LastIndex("NLT_abcabcabcabcwdcabc", "abc")
	fmt.Println("LastIndex = ", index)

	// 子串替換 strings.Replace()
	str5 := "go go hello"
	// -1 全部替換
	var n2 int = -1
	str5 = strings.Replace(str5, "go", "golang", n2)
	fmt.Println("strings.Replace() = ", str5)

	// 按照指定的某個字符為分割標示符 將一個字符串分割成字符串數組
	str6 := strings.Split("hello,world,ok", ",")
	for i, data := range str6 {
		fmt.Printf("str6[%d] = %s\t\t", i, data)
	}
	fmt.Println()

	// 大小寫轉換
	fmt.Println("轉換為小寫 = ", strings.ToLower("Go"))
	fmt.Println("轉換為大寫 = ", strings.ToUpper("Go"))

	// 去掉字符串兩邊的空格
	fmt.Println("去掉了空格 = ", strings.TrimSpace(" dergrhds gre "))

	// 去掉字符串左右兩邊的特定字符 指定的可以寫多個
	fmt.Println("去掉指定字符 = ", strings.Trim("! Hello World !", "! "))

	// 去掉左邊的
	fmt.Println("去掉左邊字符 = ", strings.TrimLeft("! Hello World !", "! "))

	// 去掉右邊的
	fmt.Println("去掉右邊字符 = ", strings.TrimRight("! Hello World !", "! "))

	// 判斷字符串是否以指定的字符串開頭
	fmt.Println("是否以指定開頭 = ", strings.HasPrefix("aaa.aa.as", "aa"))

	// 判斷字符串是否以指定的字符串結尾
	fmt.Println("是否以指定結尾 = ", strings.HasSuffix("aaa.aa.as", "aa"))

}
