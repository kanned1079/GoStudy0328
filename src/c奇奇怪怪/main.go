package main

// 注意这是个死循环会不停分配内存
func main() {
	sliceA := make([]string, 0)
	for {
		sliceA = append(sliceA, "6")
	}
}
