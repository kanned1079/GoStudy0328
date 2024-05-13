package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

// 这个是实现了方法
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器处理请求")
}

func main() {
	http.Handle("/my", &MyHandler{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
