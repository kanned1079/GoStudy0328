package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	startServer()
}

func startServer() {
	r := gin.Default()
	r.Any("/0s", func(context *gin.Context) {
		context.File("./vue.js")
	})
	r.Any("/1s", func(context *gin.Context) {
		time.Sleep(time.Second)
		context.File("./vue.js")
	})
	r.Any("/2s", func(context *gin.Context) {
		time.Sleep(time.Second * 2)
		context.File("./vue.js")
	})
	r.Any("/3s", func(context *gin.Context) {
		time.Sleep(time.Second * 3)
		context.File("./vue.js")
	})
	r.Run("localhost:8080")
}
