package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"reflect"
	"strings"
)

const path = "./config.cnf"

type MyConnection struct {
	Protocl   string `json:"protocl"`
	IpAddress string `json:"ipAddress"`
	Port      string `json:"port"`
}

func main() {
	var connection1 MyConnection = MyConnection{}
	//testReflect(&connection1)
	connection1.setConnectionInfo()
	fmt.Println(connection1)
	connection1.StartAServer()
}

// setConnectionInfo() 用于读取文件中的服务器配置信息
func (this *MyConnection) setConnectionInfo() (res bool) {
	var err error
	var file *os.File
	if file, err = os.OpenFile(path, os.O_RDONLY, 0666); err != nil {
		fmt.Println("open file err:", err)
		res = false
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("close file err:", err)
			res = false
		}
	}()
	var reader *bufio.Reader = bufio.NewReader(file)
	var strSlice = make([]string, 3)
	var i int = 0
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("READ FILE FINISHED.")
			break
		}
		//str = strings.TrimSpace(str)
		//strSlice = append(strSlice, strings.TrimSpace(str))
		strSlice[i] = strings.TrimSpace(str)
		i++
	}
	fieldsNums := reflect.TypeOf(this).Elem().NumField()
	for i := 0; i < fieldsNums; i++ {
		field := reflect.ValueOf(this).Elem().Field(i)
		field.SetString(strSlice[i])
	}
	res = true
	return
}

// StartAServer() 用于启动服务器
func (this *MyConnection) StartAServer() {
	var err error
	var listener net.Listener
	// 1. 先定义Listener
	if listener, err = net.Listen("tcp", this.IpAddress+":"+this.Port); err != nil {
		fmt.Println("listen err:", err)
	} else {
		fmt.Println("服务器在", this.IpAddress+":"+this.Port, "上启动成功")
	}
	// 2. 记得关闭Listener
	defer func() {
		if err := listener.Close(); err != nil {
			fmt.Println("Listen Accept 失败")
			fmt.Println("close listener err:", err)
			return
		}
	}()
	// 3. 循环等待客户端来连接
	for {
		// 等待客户端连接
		fmt.Println("等待客户端连接")
		// 4. 创建conn
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
		} else {
			fmt.Println("客户端连接成功 IP地址为", conn.RemoteAddr().String())
		}
		// 4. 每个连接开启一个协程
		go Worker(&conn)
	}

}

// 5. 创建Worker协程
func Worker(conn *net.Conn) {
	defer func() {
		if err := (*conn).Close(); err != nil { // 记得也要关闭conn
			fmt.Println("Close err:", err)
		}
	}()
	// 6. 循环接收数据
	for {
		buffer := make([]byte, 1024) // 定义切片存储 而且需要写在里面每次循环都清空
		//fmt.Printf("服务器正在等待客户端%s发送的消息......\n", (*conn).RemoteAddr().String())
		if _, err := (*conn).Read(buffer); err == io.EOF || err != nil {
			fmt.Println("worker err:", err)
			fmt.Println("客户端退出")
			return
		} else {
			buffer = bytes.TrimSpace(buffer)
			//fmt.Println("消息长度 = ", length)
			fmt.Println("发送的数据 = ", string(buffer))
			//fmt.Println()
		}

	}
}

func testReflect(b interface{}) {
	rTpy := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	//funcCount := rTpy.Elem().NumMethod()
	fieldCount := rVal.Elem().NumField()
	var fieldMap = make(map[string]string, 3)
	fieldMap["protocol"] = "tcp"
	fieldMap["ipaddr"] = "localhost"
	fieldMap["port"] = "8080"
	for i := 0; i < fieldCount; i++ {
		tagName := rTpy.Elem().Field(i).Tag.Get("json")
		if tagName != "" {
			fieldMap[rTpy.Elem().Field(i).Name] = tagName
		}
		fmt.Println(rTpy.Elem().Field(i).Name)
		//fieldMap[]
	}

	fmt.Println(fieldMap)

}
