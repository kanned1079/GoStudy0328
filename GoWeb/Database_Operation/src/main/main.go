package main

import "GoStudy0328/GoWeb/Database_Operation/src/model"

func main1() {
	var user1 model.User
	user1.Id = 0
	user1.UserName = "kinggyo"
	user1.Password = "1202"
	user1.Email = "kinggyo@gmail.com"

	user1.AddUser()
}
