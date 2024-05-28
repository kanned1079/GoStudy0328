package router

import "github.com/gin-gonic/gin"

func InitAll() {

	r := gin.Default()

	r.POST("/api/register", HandleRegister)

	r.Run("localhost:8080")
}
