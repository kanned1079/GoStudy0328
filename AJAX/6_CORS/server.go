package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Any("/cors", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "*")
		context.Header("Access-Control-Allow-Method", "*")
		//context.Header("Access-Control-Allow-Origin", "http://localhost:63442")
		context.String(http.StatusOK, "Hello CORS.")
	})
	r.Run("localhost:8080")
}
