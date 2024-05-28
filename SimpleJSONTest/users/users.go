package users

import (
	"gorm.io/gorm"
	"time"
)

type WxUsers struct {
	UserId    uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique" form:"email"`
	Password  string `gorm:"not null" form:"password"`
	Name      string `form:"name"`
	Age       uint8  `form:"age"`
	Phone     string `form:"phone"`
	Level     uint8  `form:"level"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *WxUsers) TableName() string {
	return "wx_users"
}
