package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Exist int    `json:"exist"`
	Msg   string `json:"msg"`
}

func main() {
	r := gin.Default()
	r.Any("/checkname", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, gin.H{
			"exist": 1,
			"msg":   "用户已经存在",
		})
	})
	r.Any("/check", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		testUser1 := User{1, "用户已经存在"}
		infos, _ := json.Marshal(testUser1)
		context.String(http.StatusOK, string(fmt.Sprintf("handle(%s)", infos)))

	})
	r.Run("localhost:8080")
}
