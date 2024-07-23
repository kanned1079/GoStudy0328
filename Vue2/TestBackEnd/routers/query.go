package routers

import (
	"GoStudy0328/Vue2/TestBackEnd/dao"
	"GoStudy0328/Vue2/TestBackEnd/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// HandleGetAllUsers 获取所有用户列表
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

// HandleAddUser 添加一个用户
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
}

// HandleDelUser 处理删除一个用户
func HandleDelUser(context *gin.Context) {

}

// HandleEditUser 编辑一个已经存在的用户
func HandleEditUser(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	log.Println("HandleEditUser")
	var newUniUser users.UniUsers
	if err := context.Bind(&newUniUser); err != nil {
		log.Println(err)
	}
	log.Println(newUniUser)
	dao.Db.Model(&newUniUser).Where("id = ?", newUniUser.Id).Updates(&newUniUser)
	context.String(http.StatusOK, "ok")
}
