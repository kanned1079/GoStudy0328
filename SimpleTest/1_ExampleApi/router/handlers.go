package router

import (
	user2 "GoStudy0328/SimpleTest/1_ExampleApi/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取所有信息
func HandleGetId(ctx *gin.Context) {
	var user user2.MyUser
	user.UserId, _ = strconv.Atoi(ctx.PostForm("id"))
	log.Println("UserId:", user.UserId)
	if code := user.FullQuery(); code != user2.QUERY_PASS {
		switch code {
		case user2.USER_NOT_EXIST:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"userExisted": "用户不存在",
			})
		}
	}
	log.Println(user)
	ctx.JSON(http.StatusOK, user)
}

// HandleUpdate 处理更新操作
func HandleUpdate(context *gin.Context) {
	var user user2.MyUser
	user.UserId, _ = strconv.Atoi(context.PostForm("id"))
	user.Name = context.PostForm("name")
	user.Gender = context.PostForm("gender")
	user.Age, _ = strconv.Atoi(context.PostForm("age"))
	user.PhoneNumber = context.PostForm("phone")
	user.Email = context.PostForm("mail")
	user.Password = context.PostForm("password")
	user.Premium = context.PostForm("lv")
	log.Println("user =", user)
	code := user.NewUpdate()
	switch code {
	case user2.USER_UPDATE_SUCCESS:
		{
			log.Println("更新成功")
			context.JSON(http.StatusOK, gin.H{"code": "ok"})

		}
	case user2.USER_UPDATE_FAIULRE:
		{
			log.Println("更新失败")
			context.JSON(http.StatusInternalServerError, gin.H{"code": "failure"})
		}

	}
	// 返回信息有 更新成功 更新失败 用户不存在 未知错误
}

func HandleRegister(ctx *gin.Context) {
	var user user2.MyUser
	GetFormMsg(ctx, &user) // 从表单中获取用户注册信息
	log.Println("HandleRegister.Get =", user)
	code := user.NewCreate()
	switch code {
	case user2.USER_CREATE_SUCCESS:
		{
			log.Println("用户创建成功")
			ctx.JSON(http.StatusOK, gin.H{"code": "ok"})
		}
	case user2.USER_CREATE_FAILURE:
		{
			log.Println("用户创建失败")
			ctx.JSON(http.StatusOK, gin.H{"code": "failure"})
		}
	}
}

func HandleDelete(ctx *gin.Context) {
	var user user2.MyUser
	GetFormMsg(ctx, &user)
	log.Println("HandleDelete.user =", user)
	code := user.NewDelete()
	switch code {
	case user2.USER_DELETE_SUCCESS:
		{
			log.Println("用户删除成功")
			ctx.JSON(http.StatusOK, gin.H{"code": "delete ok"})
		}
	case user2.USER_DELETE_FAILURE:
		{
			log.Println("用户删除失败")
			ctx.JSON(http.StatusOK, gin.H{"code": "delete failure"})
		}
	case user2.USER_NOT_EXIST:
		{
			log.Println("用户不存在")
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": "not exist"})
		}

	}
}

func GetFormMsg(context *gin.Context, user *user2.MyUser) {
	user.UserId, _ = strconv.Atoi(context.PostForm("id"))
	user.Name = context.PostForm("name")
	user.Gender = context.PostForm("gender")
	user.Age, _ = strconv.Atoi(context.PostForm("age"))
	user.PhoneNumber = context.PostForm("phone")
	user.Email = context.PostForm("mail")
	user.Password = context.PostForm("password")
	user.Premium = context.PostForm("lv")
}
