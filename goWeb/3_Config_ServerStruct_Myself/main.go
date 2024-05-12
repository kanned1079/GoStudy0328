package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
}

// 这个是实现了方法
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自定Server配置的")
}

func main() {
	// 创建Server结构体 进行详细配置
	server := http.Server{
		Addr:        ":8080",
		Handler:     &MyHandler{},
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}
