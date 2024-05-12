package main

import (
	"fmt"
	"net/http"
)

// 请求报文的格式
// 	1)请求首行（请求行）
//	2)请求信息头（请求头）
//	3)空行
//	4)请求体

// Get请求
// 没有请求体

// GET / HTTP/1.1
// Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
// Accept-Encoding: gzip, deflate, br, zstd
// Accept-Language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6
// Connection: keep-alive
// DNT: 1
// Host: localhost:8080
// Sec-Fetch-Dest: document
// Sec-Fetch-Mode: navigate
// Sec-Fetch-Site: none
// Sec-Fetch-User: ?1
// Upgrade-Insecure-Requests: 1
// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36
// sec-ch-ua: "Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"
// sec-ch-ua-mobile: ?0
// sec-ch-ua-platform: "macOS"

// Post请求

// POST /bin HTTP/1.1
// Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
// Accept-Encoding: gzip, deflate, br, zstd
// Accept-Language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6
// Cache-Control: max-age=0
// Connection: keep-alive
// Content-Length: 22
// Content-Type: application/x-www-form-urlencoded
// Cookie: Goland-85a9f76d=182cfd0f-476e-4434-8053-1ae89a469e70
// DNT: 1
// Host: localhost:8080
// Origin: http://localhost:63342
// Referer: http://localhost:63342/
// Sec-Fetch-Dest: document
// Sec-Fetch-Mode: navigate
// Sec-Fetch-Site: same-site
// Sec-Fetch-User: ?1
// Upgrade-Insecure-Requests: 1
// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36
// sec-ch-ua: "Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"
// sec-ch-ua-mobile: ?0
// sec-ch-ua-platform: "macOS"

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "测试Http协议")
}

func main() {
	// 调用处理器处理请求
	http.HandleFunc("/bin", handler)
	http.ListenAndServe(":8080", nil)

	// 路由
}
