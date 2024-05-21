package router

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/user"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	r := gin.Default()
	log.Println("数据库总记录数: ", user.RecordsCount())
	log.Println("无效记录数: ", user.InValidRecordsCount())
	log.Println("有效记录数: ", user.ValidRecordsCount())

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
