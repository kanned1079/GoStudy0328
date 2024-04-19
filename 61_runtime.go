package main

import (
	"fmt"
	"runtime"
)

// import "runtime"
// runtime包提供和go运行时环境的互操作，如控制go程的函数。
// 它也包括用于reflect包的低层次类型信息；参见reflect报的文档获取运行时类型系统的可编程接口。

// 為了充分利用多核CPU的優勢 在Golang程序中 可以設置運行的CPU數目

func main() {
	cpuCounts := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuCounts) // 設置運行時的最大CPU數
	fmt.Println("CPU(s) = ", cpuCounts)
	// go 1.8版本后程序自动运行在多核上 不需要设置
}
