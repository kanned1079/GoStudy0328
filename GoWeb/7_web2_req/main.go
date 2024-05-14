package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

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
	//length := r.ContentLength
	//// 创建一个字节切片
	//body := make([]byte, length)
	//// 读取请求体
	//r.Body.Read(body)
	//fmt.Fprintln(w, "请求体中的内容: ", string(body))
	//log.Println(string(body))
	fmt.Fprintln(w, "-------------------------------\n下面是获取表单Form	的内容")
	// Form 字段
	// 1)类型是url.Value类型 Form是解析好的表达数据 包含url字段和query参数和POST或PUT的表单数据
	// 2)字段只有在调用request的ParseForm方法后才有效 在客户端会忽略请求中的本字段而使用Body代替
	//err := r.ParseForm()
	// 如果form表单的action属性的URL地址中也有与form表单参数相同的请求参数
	// 那么参数值都可以得到 并且for吗表单中的参数值在URL的参数值的前面
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Fprintln(w, "请求参数有: ", r.Form)
	//// 下面这个只会获取到表单里面的值
	//fmt.Fprintln(w, "POST请求的from表单中的请求参数有: ", r.PostForm)

	// 可以通过FromValue和PostFromValue快速获得表单的值
	fmt.Fprintln(w, "Url中的user请求参数的值: ", r.FormValue("user"))
	fmt.Fprintln(w, "Url中的user请求参数的值: ", r.PostFormValue("username"))

}

func testJson(w http.ResponseWriter, r *http.Request) {
	// 1.设置相应头的类型
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 2.创建一个结构体实例
	user := User{
		Id:       1,
		UserName: "administrator",
		Password: "123456",
		Email:    "administrator@aaa.com",
	}
	// 3.序列化
	// 4.相应给客户端
	data, _ := json.Marshal(user)
	w.Write(data)
}

func testRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://ikanned.com:55444/")
	// 上面这个一定一定一定要在WriteHeader前面
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJson)
	http.HandleFunc("/testRedirect", testRedirect)
	http.ListenAndServe(":8080", nil)
	// 在网页端执行 http://localhost:8080/hello?user=admin&pwd=password
}
