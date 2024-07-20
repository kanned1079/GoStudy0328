package main

import (
	"GoStudy0328/Vue2/TestBackEnd/dao"
	"GoStudy0328/Vue2/TestBackEnd/users"
	"time"
)

func main() {
	testUser1 := users.UniUsers{
		Name:  "kanna",
		Age:   18,
		Sex:   1,
		Birth: time.Now(),
		Addr:  "Changzhou",
	}
	dao.Db.Model(&users.UniUsers{}).Create(&testUser1)
}
