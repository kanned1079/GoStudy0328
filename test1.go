package main

import "fmt"

func main() {
	// 還有97天放假 那麼還有幾個星期零幾天
	var daysLeft int = 97
	var weekLeft int = daysLeft / 7
	var zDaysLeft int = 97 - (weekLeft * 7)

	fmt.Printf("還有%d週零%d天\n", weekLeft, zDaysLeft)

	var tempF_1 myFloat32 = 169.4
	fmt.Printf("F:%f => C:%f\n", tempF_1, tempF_1.fToC())

	a, b := 10, 20
	a, b = b, a
	fmt.Printf("a=%d, b=%d\n", a, b)

}

type myFloat32 float32

func (tempF myFloat32) fToC() (tempC myFloat32) {
	tempC = (5.0 / 9) * (tempF - 100)
	return
}

func fToC_2(tempF float32) (tempC float32) {
	tempC = (5.0 / 9) * (tempF - 100)
	return
}
