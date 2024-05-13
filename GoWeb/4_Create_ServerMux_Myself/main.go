package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "自行创建一个多路复用器", r.URL.Path)
}

func main() {
	// 自己创建
	mux := http.NewServeMux()
	mux.HandleFunc("/bin", handler) // HandleFunc可以自动转换为一个处理器
	_ = http.ListenAndServe(":8080", mux)
}
