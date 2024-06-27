package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 定义处理 /home 请求并发送 index.html 文件
	r.GET("/home", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.File("./public/index.html")
	})

	// 启动服务器，监听端口 8080
	r.Run("localhost:8080")
}
