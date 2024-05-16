package main

// 注意这是个无限循环
// 可以导致程序占用大量内存并最终耗尽系统资源
func main() {
	sliceA := make([]string, 0)
	for {
		sliceA = append(sliceA, "6")
	}
}
