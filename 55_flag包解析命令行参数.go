package main

import (
	"flag"
	"fmt"
)

// func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
// IntVar用指定的名称、默认值、使用信息注册一个int类型flag，并将flag的值保存到p指向的变量。

func main() {
	// 定义几个变量 用于接收命令行参数
	// -u root -h 192.168.12.244 -p 3306
	var user, host string
	var port int
	flag.StringVar(&user, "u", "", "默认为空")
	flag.StringVar(&host, "h", "localhost", "默认为本机地址")
	flag.IntVar(&port, "p", 3306, "默认为3306")

	// 这里又一个很重要的方法用来解析flag 必须要调用
	// func (f *FlagSet) Parse(arguments []string) error
	// 从arguments中解析注册的flag。
	// 必须在所有flag都注册好而未访问其值时执行。
	// 未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	fmt.Printf("用户名 = %v， 主机 = %v， 端口 = %v\n", user, host, port)

}
