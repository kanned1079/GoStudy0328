package router

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/users", user.GetAllUsers)
	//r.POST("/post", user.GetData)
	r.POST("/api/getid", user.GetDataByEmail)

	r.Run("localhost:8080")

}
