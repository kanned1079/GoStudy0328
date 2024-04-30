package main

import (
	"fmt"
	"os"
)

// os.Args是一个string切片

func main() {
	fmt.Println("命令行的参数个数 = ", len(os.Args))
	for i, val := range os.Args {
		fmt.Printf("args[%d] = %v\n", i, val)
	}
}
