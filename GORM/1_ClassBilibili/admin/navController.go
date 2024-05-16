package admin

import (
	models2 "GoStudy0328/GORM/1_ClassBilibili/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NavController struct {
}

func (con NavController) Index(context *gin.Context) {
	// 查询全部数据
	navList := []models2.Nav{}
	models2.Db.Find(&navList)
	context.JSON(http.StatusOK, gin.H{
		"navList": navList,
	})

}
