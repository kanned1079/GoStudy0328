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

type NxcUser struct {
	ID                int         `json:"id"`
	InviteUserID      interface{} `json:"invite_user_id"`
	TelegramID        interface{} `json:"telegram_id"`
	Email             string      `json:"email"`
	Password          string      `json:"password"`
	PasswordAlgo      interface{} `json:"password_algo"`
	PasswordSalt      interface{} `json:"password_salt"`
	Balance           int         `json:"balance"`
	Discount          interface{} `json:"discount"`
	CommissionType    int         `json:"commission_type"`
	CommissionRate    interface{} `json:"commission_rate"`
	CommissionBalance int         `json:"commission_balance"`
	T                 int         `json:"t"`
	U                 int         `json:"u"`
	D                 int         `json:"d"`
	TransferEnable    int64       `json:"transfer_enable"`
	Banned            int         `json:"banned"`
	IsAdmin           int         `json:"is_admin"`
	LastLoginAt       int         `json:"last_login_at"`
	IsStaff           int         `json:"is_staff"`
	LastLoginIP       interface{} `json:"last_login_ip"`
	UUID              string      `json:"uuid"`
	GroupID           int         `json:"group_id"`
	PlanID            int         `json:"plan_id"`
	SpeedLimit        int         `json:"speed_limit"`
	RemindExpire      int         `json:"remind_expire"`
	RemindTraffic     int         `json:"remind_traffic"`
	Token             string      `json:"token"`
	ExpiredAt         int         `json:"expired_at"`
	Remarks           interface{} `json:"remarks"`
	CreatedAt         int         `json:"created_at"`
	UpdatedAt         int         `json:"updated_at"`
	TotalUsed         int         `json:"total_used"`
	PlanName          string      `json:"plan_name"`
	SubscribeURL      string      `json:"subscribe_url"`
}

func (u *WxUsers) TableName() string {
	return "wx_users"
}
