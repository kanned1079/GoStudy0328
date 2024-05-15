package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 解析模版
	t := template.Must(template.ParseFiles("GoWeb/9_BookStore/views/index.html"))
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)

}
