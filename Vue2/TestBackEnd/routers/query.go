package routers

import (
	"GoStudy0328/Vue2/TestBackEnd/dao"
	"GoStudy0328/Vue2/TestBackEnd/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var err error

type pageData struct {
	page  int
	limit int
}

// HandleGetAllUsers 获取所有用户列表
func HandleGetAllUsers(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	log.Println("-------------------------------HandleGetAllUsers------")
	var pgData pageData
	var recordCount int64                                                                   // 从数据库中获取用户的总记录数 数据类型只能使用int64
	if result := dao.Db.Model(&users.UniUsers{}).Count(&recordCount); result.Error != nil { // 使用Count方法获取用户记录数
		log.Panicf(result.Error.Error())
	}
	log.Println("总用户计数: ", recordCount)
	pgData.page, _ = strconv.Atoi(context.DefaultQuery("page", "1"))
	pgData.limit, _ = strconv.Atoi(context.DefaultQuery("limit", "10"))

	offset := (pgData.page - 1) * pgData.limit

	var uniUsers []users.UniUsers
	if result := dao.Db.Limit(pgData.limit).Offset(offset).Find(&uniUsers); result.Error != nil {
		log.Panicf(result.Error.Error())
	}
	context.JSON(http.StatusOK, gin.H{
		"data":  uniUsers,    // 用户列表
		"count": recordCount, // 总用户计数
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
	context.Header("Access-Control-Allow-Origin", "*")
	var u users.UniUsers
	if err := context.Bind(&u); err != nil {
		log.Println(err)
	}
	//log.Println(u)
	dao.Db.Model(&u).Delete(&u)
	//var id int
	//if id, err = strconv.Atoi(context.PostForm("id")); err != nil {
	//	log.Println(err)
	//}
	//log.Println("id: ", id)

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
