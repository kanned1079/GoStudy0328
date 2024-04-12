package main

import "fmt"

func main() {
	var y myInt = 2000
	var y2 int = 2019
	fmt.Printf("year %d = %+v\n", y, y.isLyear())
	fmt.Printf("year %d = %+v\n", y2, isLyear2(y2))
}

type myInt int
type myBool bool

func (year myInt) isLyear() (result myBool) {
	result = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	return
}

func isLyear2(year int) (result bool) {
	result = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	return
}
