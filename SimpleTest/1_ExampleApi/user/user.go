package user

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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
	DeletedAt   gorm.DeletedAt
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

func HandleRegister(context *gin.Context) {
	var user MyUser
	user.Name = context.PostForm("name")
	user.Gender = context.PostForm("gender")
	user.Age, _ = strconv.Atoi(context.PostForm("age"))
	user.PhoneNumber = context.PostForm("phone")
	user.Email = context.PostForm("mail")
	user.Password = context.PostForm("password")
	user.Premium = context.PostForm("lv")
	user.CreatedAt = time.Now()
	if IsExist(user.Email) {
		result := dao.Db.Create(&user)
		log.Println(result.RowsAffected)
		if result.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user})
		}
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "用户已存在"})
	}

}

func IsExist(email string) (res bool) {
	log.Println("email:", email)
	var user MyUser
	user.Email = email
	result := dao.Db.Model(&MyUser{}).Where("email = ?", email).First(&user)
	if result.RowsAffected == 0 {
		log.Println("该用户已注册")
		return true
	} else {
		log.Println("用户不存在")
		return false
	}
}

func HandleLogin(context *gin.Context) {
	var user MyUser
	user.Email = context.PostForm("mail")
	user.Password = context.PostForm("password")
	log.Println(user)
	newUser, isExist := AuthUserInfo(user.Email, user.Password)
	if isExist {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "user": newUser})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "用户不存在或密码错误"})
	}
}

func AuthUserInfo(email string, password string) (u *MyUser, res bool) {
	var user MyUser
	result := dao.Db.Where("email = ?", email).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		log.Println("用户不存在")
		return nil, false
	}
	if user.Password == password {
		log.Println("密码正确")
		return &user, true
	} else {
		log.Println("密码错误")
		return nil, false
	}
}

func HandlerDeleteUser(context *gin.Context) {
	var email string
	email = context.PostForm("mail")
	log.Println("HandlerDeleteUser:", email)
	var user MyUser
	result := dao.Db.Where("email = ?", email).Limit(1).Find(&user)
	log.Println(result.RowsAffected, result.Error)
	log.Println(user)
	if result.RowsAffected == 1 {
		result = dao.Db.Model(&MyUser{}).Where("email = ?", email).Delete(&user)
		if result.RowsAffected == 1 {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "deleted": "ok"})

		}
	}
	//if IsExist(email) {
	//	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "deleted": "ok"})
	//} else {
	//	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "deleted": "false"})
	//}

}

func IsIDExist(id int) (res bool) {
	log.Println("id:", id)
	var user MyUser
	result := dao.Db.Model(&MyUser{}).Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		log.Println("该用户已注册")
		return true
	} else {
		log.Println("用户不存在")
		return false
	}
}

//
//func translateJSON() {
//	context.JSON(http.StatusOK, &MyUser{})
//}
