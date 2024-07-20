package user

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"gorm.io/gorm"
	"log"
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

//const (
//	= "UserIsNotCreate"
//
//)

//const (
//	USER_IS_NOT_EXIST  = 403
//	USER_IS_EXISTED    = 201
//	INCORRECT_PASSWORD = 500
//	USER_AUTH_PASS     = 1200
//	USER_AUTH_FAIL     = 1501
//)

type Response struct {
	Code   int    `json:"code"`
	MyUser MyUser `gorm:" - "`
}

func (MyUser) TableName() string {
	return "myusers"
}

//func GetDataByEmail(context *gin.Context) {
//	nameData := context.PostForm("mail")
//	passData := context.PostForm("pass")
//	log.Println("data:", nameData)
//	user := MyUser{}
//	result := dao.dao.Where("email = ?", nameData).Limit(1).Find(&user)
//	if result.RowsAffected == 0 {
//		context.JSON(http.StatusInternalServerError, gin.H{"error": "用户不存在"})
//		log.Println("err:", result.Error)
//	} else {
//		if user.Password == passData {
//			var resp Response
//			resp.Code = http.StatusOK
//			resp.MyUser = user
//			context.JSON(http.StatusOK, resp)
//		} else {
//			context.JSON(http.StatusInternalServerError, gin.H{"error": "密码错误"})
//
//		}
//
//	}
//}
//
//func GetAllUsers(context *gin.Context) {
//	var user1 MyUser
//	result := dao.dao.First(&user1)
//	if result.Error != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
//		log.Println("err")
//	} else {
//		var response Response
//		response.Code = http.StatusOK
//		response.MyUser = user1
//		context.JSON(response.Code, response)
//	}
//}

func (u *MyUser) Insert() {
	var ency Encryptor
	u.Password = ency.Encrypt(u.Password)
	result := dao.Db.Create(&u)
	if result.Error != nil {
		log.Println("insert err:", result.Error)
	} else {
		log.Println("insert success")
	}
}

// Register 绑定的方法
//func (u *MyUser) Register() (resp int) {
//	if IsIdExist(u.UserId) == USER_IS_NOT_EXIST {
//		log.Println("此用户不存在 ID=", u.UserId)
//		return 0
//	} else {
//		result := dao.dao.Model(&MyUser{}).Create(&u)
//		if result.Error != nil {
//			log.Println("insert err:", result.Error)
//			return 4
//		}
//	}
//	return 0
//}

//func HandleRegister(context *gin.Context) {
//	var user MyUser
//	user.Name = context.PostForm("name")
//	user.Gender = context.PostForm("gender")
//	user.Age, _ = strconv.Atoi(context.PostForm("age"))
//	user.PhoneNumber = context.PostForm("phone")
//	user.Email = context.PostForm("mail")
//	user.Password = context.PostForm("password")
//	user.Premium = context.PostForm("lv")
//	user.CreatedAt = time.Now()
//	if IsExist(user.Email) == USER_IS_NOT_EXIST { // 用户不存在就创建用户
//		var ency Encryptor
//		user.Password = ency.Encrypt(user.Password)
//		result := dao.dao.Create(&user)
//		log.Println(result.RowsAffected)
//		if result.Error != nil {
//			context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		} else {
//			context.JSON(http.StatusOK, gin.H{"user": user})
//		}
//	} else {
//		context.JSON(http.StatusInternalServerError, gin.H{"error": "用户已存在"})
//	}
//
//}
//
//func IsExist(email string) (Code int) {
//	log.Println("email:", email)
//	var user MyUser
//	user.Email = email
//	dao.dao.Where("email = ?", email).Limit(1).Find(&user)
//	log.Println(user)
//	if user.UserId != 0 {
//		log.Println("IsExist:该用户已注册")
//		return USER_IS_EXISTED
//	} else {
//		log.Println("IsExist:用户不存在")
//		return USER_IS_NOT_EXIST
//	}
//}
//
//func HandleLogin(context *gin.Context) {
//	var user MyUser
//	user.Email = context.PostForm("mail")
//	user.Password = context.PostForm("password")
//	log.Println(user)
//	newUser, isExist := AuthUserInfo(user.Email, user.Password)
//	if isExist == USER_IS_EXISTED {
//		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "user": newUser})
//	} else if isExist == USER_IS_NOT_EXIST {
//		context.JSON(http.StatusInternalServerError, gin.H{"error": "用户不存在"})
//	} else if isExist == USER_AUTH_PASS {
//
//		context.JSON(http.StatusInternalServerError, gin.H{"error": "ok"})
//	}
//}
//
//func AuthUserInfo(email string, password string) (u *MyUser, code int) {
//	var user MyUser
//	result := dao.dao.Where("email = ?", email).Limit(1).Find(&user)
//	log.Println(user)
//	var ency Encryptor
//	user.Password = ency.Decrypt(user.Password)
//	if result.RowsAffected == 0 {
//		log.Println("用户不存在")
//		return nil, USER_IS_NOT_EXIST
//	}
//	if user.Password == password {
//		log.Println("密码正确")
//		return &user, USER_AUTH_PASS
//	} else {
//		log.Println("密码错误")
//		return nil, INCORRECT_PASSWORD
//	}
//}
//
//// HandlerDeleteUser 处理删除用户 使用软删除
//func HandlerDeleteUser(context *gin.Context) {
//	var email string
//	email = context.PostForm("mail")
//	log.Println("HandlerDeleteUser:", email)
//	var user MyUser
//	result := dao.dao.Where("email = ?", email).Limit(1).Find(&user)
//	log.Println(result.RowsAffected, result.Error)
//	log.Println(user)
//	if result.RowsAffected == 1 {
//		result = dao.dao.Model(&MyUser{}).Where("email = ?", email).Delete(&user)
//		if result.RowsAffected == 1 {
//			context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "deleted": "ok"})
//		}
//	} else {
//
//	}
//	//if IsExist(email) {
//	//	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "deleted": "ok"})
//	//} else {
//	//	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "deleted": "false"})
//	//}
//
//}
//
//func HandlerUpdateUser(context *gin.Context) {
//	var user MyUser
//	user.Name = context.PostForm("name")
//	user.Gender = context.PostForm("gender")
//	user.Age, _ = strconv.Atoi(context.PostForm("age"))
//	user.PhoneNumber = context.PostForm("phone")
//	user.Email = context.PostForm("mail")
//	user.Password = context.PostForm("password")
//	user.Premium = context.PostForm("lv")
//
//	user.UserId, _ = strconv.Atoi(context.PostForm("id"))
//	log.Println("需要查询的用户Id: ", user.UserId)
//	if IsIdExist(user.UserId) == USER_IS_EXISTED {
//		var ency Encryptor
//		user.Password = ency.Encrypt(user.Password)
//		result := dao.dao.Model(&MyUser{}).Where("user_id = ?", user.UserId).Updates(&user)
//		log.Println(result.RowsAffected, result.Error)
//		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": "update ok"})
//	} else {
//		context.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "data": "not found"})
//	}
//}
//
//func IsIdExist(id int) (code int) {
//	log.Println("UserId: ", id)
//	var usr MyUser
//	result := dao.dao.Model(&MyUser{}).Where("user_id = ?", id).Limit(1).Find(&usr)
//	log.Println("RowsAffected: ", result.RowsAffected)
//	if usr.UserId == id {
//		log.Println("指定Id的用户存在")
//		return USER_IS_EXISTED
//	} else {
//		log.Println("指定Id的用户不存在")
//		return USER_IS_NOT_EXIST
//	}
//}

//func (u *MyUser) AfterSave(tx *gorm.DB) {
//	u.UpdatedAt = time.Now()
//}

//
//func translateJSON() {
//	context.JSON(http.StatusOK, &MyUser{})
//}
