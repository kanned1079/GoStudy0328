package main

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"GoStudy0328/SimpleTest/1_ExampleApi/router"
)

func main() {
	dao.InitDB()
	//user.TuncatDatabase()
	router.InitRouter()
}
