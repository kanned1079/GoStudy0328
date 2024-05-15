package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	// 解析模版

	// template.ParseFiles 会创建一个同名模版 在后面再进行解析
	// 我们在解析模板时都没有对错误进行处理，Go 提供了一个 Must 函数专门用来处理这个错误。
	// Must 函数可以包裹起一个函数，被包裹的函数会返回一个指向模板的指针和一个错误，如果错误不是 nil，那么 Must 函数将产生一个 panic。
	t := template.Must(template.ParseFiles("GoWeb/8_template/index1.html"))
	// Must会自动对异常进行处理
	//template.ParseGlob() 可以一次性读入多个模版文件

	// 执行
	t.Execute(w, "Hello World") // 如果只有一个模版就可以使用这个
	// 如果有多个 那就使用 t.ExecuteTemplate()
	//

}

func test2Template(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("GoWeb/8_template/index1.html", "GoWeb/8_template/index2.html"))
	t.ExecuteTemplate(w, "index2.html", "val1")
}

func main() {
	http.HandleFunc("/t", testTemplate)
	http.HandleFunc("/t2", test2Template)
	http.ListenAndServe(":8080", nil)
}
