package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() // 创建一个默认的路由

	// 对路由进行配置
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "值：%v", "Hello Gin!")
		// 用http.StatusOK需要导入包http
	})
	// 第二个路由
	r.GET("/news", func(context *gin.Context) {
		context.String(200, "值：%v", "第二个页面")
	})
	// 配置一个处理POST 主要用于增加数据
	r.POST("/add", func(context *gin.Context) {
		context.String(200, "这是一个post请求")
	})
	// 配置一个处理put 主要用于编辑数据
	r.PUT("/put", func(context *gin.Context) {
		context.String(200, "这是一个put请求")
	})

	// 配置一个处理delete 主要用于删除
	r.DELETE("/del", func(context *gin.Context) {
		context.String(200, "这是一个del请求")
	})

	// 启动服务器监听
	r.Run(":9090")
}
