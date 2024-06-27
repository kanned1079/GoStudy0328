package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func main() {
	r := gin.Default()

	// 定义处理 /home 请求并发送 index.html 文件
	r.GET("/home", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.File("./public/index.html")
	})

	r.GET("/data", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.String(http.StatusOK, "hello world")
	})

	r.GET("/jsonp", func(context *gin.Context) {
		context.String(http.StatusOK, "console.log('hello jsonp')")
	})

	r.GET("/jsonp2", func(context *gin.Context) {
		var usr1 = user{"kinggyo", 16}
		str, _ := json.Marshal(usr1)
		context.String(http.StatusOK, fmt.Sprintf("handle(%s)", str))
	})

	r.GET("/jsonp3", func(context *gin.Context) {
		var usr1 = user{"kanna", 21}
		context.JSONP(http.StatusOK, usr1)
	})

	// 启动服务器，监听端口 8080
	r.Run("localhost:8080")
}
