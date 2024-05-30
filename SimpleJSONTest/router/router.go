package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func InitAll() {

	r := gin.Default()

	r.POST("/api/register", HandleRegister)
	r.POST("/api/login", HandleLogin)
	r.POST("/api/update", HandleUpdate)

	if err := r.Run("localhost:8080"); err != nil {
		log.Println("服务器监听端口异常 可能是端口已被占用")
		os.Exit(0)
	}
}
