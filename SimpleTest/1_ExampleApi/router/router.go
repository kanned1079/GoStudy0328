package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	//r.GET("/users", user.GetAllUsers)
	////r.POST("/post", user.GetData)
	//r.POST("/api/getid", user.GetDataByEmail)
	//r.POST("/api/reg", user.HandleRegister)
	//r.POST("/api/login", user.HandleLogin)
	//r.POST("/api/del", user.HandlerDeleteUser)
	//r.POST("/api/update", user.HandlerUpdateUser)
	r.POST("/api/getById", HandleGetId)
	r.POST("/api/update", HandleUpdate)
	r.POST("/api/register", HandleRegister)
	r.POST("/api/delete", HandleDelete)

	r.Run("localhost:8080")

}
