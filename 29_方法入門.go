package main

import "fmt"

type Person_29 struct {
	Name string
}

func (p Person_29) test29() {
	p.Name = "kanna"
	fmt.Println("test29 p's Name = ", p.Name)
}

func (p Person_29) jisuan(end int) {
	var sum int = 0
	for i := 1; i <= end; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}

func (p Person_29) getSum(a, b int) (res int) {
	res = a + b
	return
}

func main() {
	p1 := Person_29{}
	p1.Name = "Tom"
	p1.test29()
	p1.jisuan(100)
	fmt.Println("getSum(10, 24) = ", p1.getSum(10, 24))
}
