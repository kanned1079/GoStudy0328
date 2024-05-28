package router

import (
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
			ResponseMsg(context, http.StatusInternalServerError, err.Error())
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

func ResponseMsg(context *gin.Context, status any, msg any) {
	context.JSON(http.StatusOK, gin.H{
		"code": status,
		"msg":  msg,
	})
}
