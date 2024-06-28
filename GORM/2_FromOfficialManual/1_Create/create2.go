package main

import (
	"GoStudy0328/SimpleJSONTest/dao"
	"time"
)

type NxcUser struct {
	ID           int    `json:"id"`
	InviteUserID any    `json:"invite_user_id"`
	TelegramID   any    `json:"telegram_id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type OfficialUser struct {
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	testuser := NxcUser{
		ID:           1063421234,
		InviteUserID: nil,
		TelegramID:   "sakey7202",
		Email:        "sakey7202@ikanned.com",
		Password:     "123456",
	}
	dao.Db.Model(&NxcUser{}).Create(&testuser)
}
