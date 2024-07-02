package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
	"time"
)

type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	Addr string `json:"addr"`
}

func main() {
	StartAll()
}

func StartAll() {
	go showNumbers()
	r := gin.Default()
	us1 := User{Name: "kanna", Age: 20, Addr: "江苏省常州市"}

	r.GET("/server", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   "GET: hello from gin",
		})
	})

	r.GET("/ie", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.String(http.StatusOK, "hello IE")
	})

	r.POST("/json-server", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		//context.Header("Access-Control-Allow-Headers", "*")
		context.JSON(http.StatusOK, us1)
		data := context.PostForm("a")
		log.Println(data)
	})

	r.GET("/axios1", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		//context.String(http.StatusOK, "GET: hello from gin via axios")
		context.JSON(http.StatusOK, us1)
	})

	// 中间件来处理 CORS 预检请求
	r.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Methods", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	})

	r.POST("/axios2", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Content-Type", "application/json")
		context.Header("Access-Control-Allow-Methods", "*")
		username := context.PostForm("name")
		pwd := context.PostForm("pwd")
		log.Println(username, pwd)
		context.String(http.StatusOK, "POST: Hello from Gin")
	})

	r.GET("/fetch", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.String(http.StatusOK, "GET: hello from gin")
		//context.JSON(http.StatusOK, us1)
	})

	r.POST("/fetch", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		//context.String(http.StatusOK, "POST: hello from gin")
		context.JSON(http.StatusOK, us1)
	})

	r.Run("localhost:9000")
}

func showNumbers() {
	for {
		log.Println("当前GO协程数量: ", runtime.NumGoroutine())
		time.Sleep(time.Second * 30)
	}

}
