package router

import (
	"GoStudy0328/SimpleJSONTest/encrypt"
	"GoStudy0328/SimpleJSONTest/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleRegister(context *gin.Context) {
	var user users.WxUsers
	if err := context.Bind(&user); err != nil {
		log.Println("HandleRegister err: ", err)
		ResponseMsg(context, http.StatusInternalServerError, err.Error())
	}
	isExist, err := users.IsUserExistByEmail_New(user.Email)
	log.Printf("isExist: %v, err: %v\n", isExist, err)
	if isExist == users.User_Not_Exist {
		log.Println("执行注册")
		code, err := user.CreateNew()
		if code == users.Insert_Success {
			ResponseMsg(context, http.StatusOK, user)
		} else {
			ResponseMsg(context, http.StatusInternalServerError, err) //pass
		}
	} else {
		log.Println("用户已经存在")
		ResponseMsg(context, http.StatusInternalServerError, err.Error())
	}

	//if isExist == users.User_Exist {
	//	log.Println("HandleRegister err: ", err)
	//	ResponseMsg(context, http.StatusInternalServerError, err)
	//} else if isExist == users.User_Not_Exist {
	//	code, err := user.CreateNew()
	//	if code == users.Insert_Failure {
	//		log.Println("HandleRegister err: ", err)
	//		ResponseMsg(context, http.StatusInternalServerError, err)
	//	} else if code == users.Insert_Success {
	//		ResponseMsg(context, http.StatusOK, "ok")
	//	} else {
	//		log.Println("HandleRegister err: ", err)
	//		ResponseMsg(context, http.StatusInternalServerError, err)
	//	}
	//}
	//context.JSON(http.StatusOK, user)
	//log.Println(user)

}

func HandleLogin(context *gin.Context) {
	var thisUser users.WxUsers
	var user users.WxUsers
	if err := context.Bind(&thisUser); err != nil {
		log.Println("HandleLogin err: ", err)
		ResponseMsg(context, http.StatusInternalServerError, err.Error())
	}
	// 检查此用户是否存在
	isExist, err := users.IsUserExistByEmail_New(thisUser.Email)
	log.Println("[HandleLogin] err: ", err)
	if isExist != users.User_Exist {
		log.Println("用户不存在")
		ResponseMsg(context, http.StatusInternalServerError, err.Error())
		return
	}
	user.Email = thisUser.Email
	// 开始请求该用户所有信息
	if code, err := user.FullQueryByEmail(); code != users.Query_Pass {
		ResponseMsg(context, http.StatusInternalServerError, err)
		return
	}
	// 处理登录完成后面的消息
	log.Println("user.Password: ", user.Password)
	log.Println("thisUser.Password:", thisUser.Password)
	thisUser.Password = encrypt.Base64Encryptor.Base64Encrypt(thisUser.Password)
	if user.Password == thisUser.Password {
		context.JSON(http.StatusOK, gin.H{
			"code:": users.Auth_Pass,
			"data":  user,
		})
	} else {
		log.Println("密码错误")
		ResponseMsg(context, users.Auth_Failure, "密码错误")
		return
	}

}

func HandleUpdate(context *gin.Context) {
	//var updateUserId = -1
	var updateUserEmail string
	var thisUser users.WxUsers
	//updateUserId, _ = strconv.Atoi(context.PostForm("update_user_id"))
	updateUserEmail = context.PostForm("email")
	if err := context.Bind(&thisUser); err != nil {
		log.Println("HandleUpdate err: ", err)
		ResponseMsg(context, http.StatusInternalServerError, err.Error())
	}
	// 根据Id查询用户是否存在
	isExist, err := users.IsUserExistByEmail_New(updateUserEmail)
	if isExist != users.User_Exist {
		ResponseMsg(context, http.StatusInternalServerError, err.Error()) //pass
		return
	}
	// 后面处理用户更新逻辑
	//thisUser.UserId = uint(updateUserId)
	thisUser.Email = updateUserEmail
	if code, err := thisUser.UpdateUserInfo(); code != users.Update_Success {
		ResponseMsg(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": users.Update_Success,
	})

}

func ResponseMsg(context *gin.Context, status any, msg any) {
	context.JSON(http.StatusOK, gin.H{
		"code": status,
		"msg":  msg,
	})
}
