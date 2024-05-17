package user

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type MyUser struct {
	UserId      int    `gorm:"primary_key unique"`
	Name        string `gorm:"unique"`
	Gender      string
	Age         int
	PhoneNumber string
	Email       string `gorm:"unique"`
	Password    string
	Premium     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Response struct {
	Code   int    `json:"code"`
	MyUser MyUser `gorm:" - "`
}

func (MyUser) TableName() string {
	return "myusers"
}

func GetDataByEmail(context *gin.Context) {
	nameData := context.PostForm("mail")
	passData := context.PostForm("pass")
	log.Println("data:", nameData)
	user := MyUser{}
	result := dao.Db.Where("email = ?", nameData).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "用户不存在"})
		log.Println("err:", result.Error)
	} else {
		if user.Password == passData {
			var resp Response
			resp.Code = http.StatusOK
			resp.MyUser = user
			context.JSON(http.StatusOK, resp)
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "密码错误"})

		}

	}
}

func GetAllUsers(context *gin.Context) {
	var user1 MyUser
	result := dao.Db.First(&user1)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		log.Println("err")
	} else {
		var response Response
		response.Code = http.StatusOK
		response.MyUser = user1
		context.JSON(response.Code, response)
	}
}

func (u *MyUser) Insert() {
	result := dao.Db.Create(&u)
	if result.Error != nil {
		log.Println("insert err:", result.Error)
	} else {
		log.Println("insert success")
	}
}

func GetID(context *gin.Context) {

}

//
//func translateJSON() {
//	context.JSON(http.StatusOK, &MyUser{})
//}
