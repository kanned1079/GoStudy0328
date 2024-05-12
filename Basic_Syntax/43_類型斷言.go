package main

import "fmt"

// 类型断言，由于接口是一般类型，不知道具体类型
// 如果要轉換成具體類型 就需要用到類型斷言

type Point_43 struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point1 Point_43 = Point_43{1, 2}
	a = point1
	var b Point_43
	//b = a // 這裏會報錯 需要是用類型斷言
	b = a.(Point_43) // type assertion (類型斷言)
	fmt.Println(b)

	var x interface{}
	var b1 float32 = 1.1
	x = b1           // 空接口可以接收任何類型
	y := x.(float32) // 這裏你不知道又沒有轉換成功 那麼就需要一喔判斷的機制
	fmt.Printf("y的類型=%T, val= %v\n", y, y)

	// 在進行類型斷言的時候 如果結果不匹配 就會報panic錯
	// 因此在進行類型斷言的時候 要確保原類型是空接口之前的類型
	// 帶有檢查的類型斷言

	var x2 interface{}
	var b12 float32 = 1.1
	x2 = b12 // 空接口可以接收任何類型
	//y2 := x2.(float32) // 這裏你不知道又沒有轉換成功 那麼就需要一喔判斷的機制
	y2, ok := x2.(float64)
	if ok { // true
		fmt.Printf("轉換成功\n")
	} else { // false 雖然這裏失敗 但是程序可以繼續執行 不會奔潰
		fmt.Printf("轉換失敗\n")
	}
	fmt.Printf("y2的類型=%T, val2= %v\n", y2, y2)

}
