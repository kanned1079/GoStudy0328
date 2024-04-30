package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nums1 [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums1[i] = rand.Intn(100)
	}
	fmt.Println("nums1 = ", nums1)

	for i := 0; i < len(nums1)-1; i++ {
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}

	fmt.Println("sorted num1 = ", nums1)

}
