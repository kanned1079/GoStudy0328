package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Mankind struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	r := gin.Default()
	// 先解析网页
	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob("templates/**/*")
	// 后面就需要写模版名称
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// 以下要求 请求json接口的时候返回一个json数据
	r.GET("/json", func(context *gin.Context) {
		data, _ := json.Marshal(Mankind{"kanna", 20, "Shanghai"})
		context.String(http.StatusOK, fmt.Sprintln(string(data)))
	})

	// 根据请求返回指定html
	r.GET("/def/goods", func(context *gin.Context) {
		context.HTML(http.StatusOK, "defaults/goods.html", gin.H{
			"title": "Goods",
		})
	})

	r.GET("def/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "defaults/news.html", gin.H{
			"title":  "News",
			"newsId": 678567,
		})
	})

	r.GET("admin/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title":  "News",
			"newsId": 678567,
		})
	})

	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "NewTitle",
			// 下面的用于条件判断使用
			"score":  91,
			"hobby":  []string{"aaaa", "bbbb", "cccc"},
			"hobby1": []string{},
			"news": &Article{
				Title: "新闻111",
				Body:  "内容111111111",
			},
		})
	})

	// gin中的一些比较函数
	// eq 如果 arg1 == arg2 则返回真如果
	// ne 如果 arg1 != arg2 则返回真如果
	// lt 如果 arg1 < arg2  则返回真如果
	// le 如果 arg1 <= arg2 则返回真如果
	// gt 如果 arg1 > arg2  则返回真如果
	// ge 如果 arg1 >= arg2 则返回真

	r.Run("localhost:8080")
}
