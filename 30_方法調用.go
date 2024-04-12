package main

import (
	"fmt"
	"math"
)

// 方法的调用和传参机制和函数基本一样
// 不一样的地方是方法调用时会将调用方法的变量，当做实参也传递给方法

type Circle_30 struct {
	radius int
}

func (c Circle_30) area() (a float64) {
	a = 3.1415 * math.Pow(float64(c.radius), 2)
	return
}

// 為提高效率 這裏使用指針
func (c *Circle_30) area2() (a float64) {
	// 因為是指針 因此使用標準寫法
	a = 3.1415 * math.Pow(float64((*c).radius), 2)
	return
}

func main() {
	cir1 := Circle_30{}
	cir1.radius = 10
	fmt.Println("area = ", cir1.area())

	cir2 := Circle_30{}
	cir2.radius = 5
	fmt.Println("area = ", (&cir2).area())
}
