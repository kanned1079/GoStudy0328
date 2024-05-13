package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "请求路径", r.URL.Path)
	fmt.Fprintln(w, "请求地址后查询的字符串是", r.URL.RawQuery)
	// 获取请求头所有信息
	fmt.Fprintln(w, "请求头中所有的信息 ", r.Header) // 这个拿到的是一个map
	// 如果需要拿到其中一个信息
	fmt.Fprintln(w, "Accept-Encoding:", r.Header.Get("Accept-Encoding"))
	fmt.Fprintln(w, "User-Agent:", r.Header.Get("User-Agent"))
	fmt.Fprintln(w, "-------------------------------\n下面是获取请求体body的内容")
	//获取内容长度
	length := r.ContentLength
	// 创建一个字节切片
	body := make([]byte, length)
	// 读取请求体
	r.Body.Read(body)
	fmt.Fprintln(w, "请求体中的内容: ", string(body))
	log.Println(string(body))
	fmt.Fprintln(w, "-------------------------------\n下面是获取表单Form	的内容")
	// Form 字段
	// 1)类型是url.Value类型 Form是解析好的表达数据 包含url字段和query参数和POST或PUT的表单数据
	// 2)字段只有在调用request的ParseForm方法后才有效 在客户端会忽略请求中的本字段而使用Body代替

}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
	// 在网页端执行 http://localhost:8080/hello?user=admin&pwd=password
}
