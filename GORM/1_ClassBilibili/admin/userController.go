package admin

import (
	"GoStudy0328/Basic_Syntax/src/79_Chat/server/model"
	"GoStudy0328/GORM/1_ClassBilibili/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (con UserController) Index(context *gin.Context) {
	context.String(http.StatusOK, "这是查询用户列表")
	userList := []model.User{}
	models.Db.Find(&userList)
	context.JSON(http.StatusOK, gin.H{
		"userList": userList,
	})
}
