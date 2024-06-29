package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Any("/test", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.String(http.StatusOK, "hello world")
	})

	r.Run("localhost:8080")
}
