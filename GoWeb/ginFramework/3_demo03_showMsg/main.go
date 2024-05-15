package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Id  int    `json:"id"`
	Mes string `json:"mes"`
}

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "值: %v\n", "首页") // 这是返回的一个String数据
	})

	r.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]any{
			"success": true,
			"msg":     "Hello Gin",
		})
	})

	r.GET("/json2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{ // 和上面的mao[string]any 相同
			"success": false,
			"msg":     "Hello Gin2",
		})
	})

	r.GET("/article", func(context *gin.Context) {
		var article = &Article{
			Title:   "This is Tiele",
			Desc:    "This is Desc",
			Content: "This is Content",
		}
		context.JSON(http.StatusOK, article)
	})

	// 相应一个JSOPN 主要用户跨域请求 jsonp可以传入回调函数
	// http://localhost:8080/jsopn?callback=xxxxxxxx
	r.GET("jsopn", func(context *gin.Context) {
		context.JSON(http.StatusOK, &Article{
			Title:   "This is Tiele from jsopn",
			Desc:    "This is Desc",
			Content: "This is Content",
		})
	})

	// 相应xml
	r.GET("xml", func(context *gin.Context) {
		// 使用context.XML
		context.XML(http.StatusOK, &Article{
			Title:   "This is Tiele from xml",
			Desc:    "This is Desc",
			Content: "This is Content",
		})

	})

	r.Run("localhost:8080")
}
