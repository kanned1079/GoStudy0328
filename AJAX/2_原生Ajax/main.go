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
	// us1 := User{Name: "kanna", Age: 20, Addr: "Changzhou"}

	r.GET("/server", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   "GET: hello from gin",
		})
	})

	r.GET("/ie", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.String(http.StatusOK, "hello ie1")
	})

	r.GET("/query1", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		time.Sleep(time.Second * 3) // 延迟3s
		context.String(http.StatusOK, "hello from gin after 3sec")
	})

	r.POST("/server", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		//context.Header("Access-Control-Allow-Headers", "*")
		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   "POST: hello from gin",
		})
		data := context.PostForm("a")
		log.Println(data)
	})

	r.Run("localhost:9000")
}

func showNumbers() {
	for {
		log.Println("当前GO协程数量: ", runtime.NumGoroutine())
		time.Sleep(time.Minute)
	}

}
