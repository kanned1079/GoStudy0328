package routers

import (
	"GoStudy0328/Vue2/TestBackEnd/dao"
	"GoStudy0328/Vue2/TestBackEnd/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleGetAllUsers(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	var uniUsers []users.UniUsers
	if result := dao.Db.Find(&uniUsers).Limit(100); result.Error != nil {
		log.Panicf(result.Error.Error())
	}
	context.JSON(http.StatusOK, gin.H{
		"data": uniUsers,
	})
}

func HandleAddUser(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	var uniUser users.UniUsers
	if err := context.Bind(&uniUser); err != nil {
		log.Println(err)
	}
	log.Println(uniUser)
	dao.Db.Model(&uniUser).Create(&uniUser)
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
	//var uniUser users.UniUsers
	//context.PostForm()
}
func HandleDelUser(context *gin.Context) {

}

func HandleEditUser(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
}
