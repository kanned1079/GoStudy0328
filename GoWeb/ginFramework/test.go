package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Any("/api", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")             // 允许跨域
		context.JSON(http.StatusOK, gin.H{"name": "kanna", "age": 18}) // 相应json
		context.String(http.StatusOK, "这是返回的字符串")                      // 相应字符串
	})
}
