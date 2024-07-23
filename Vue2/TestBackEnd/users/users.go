package users

import (
	"gorm.io/gorm"
	"time"
)

type UniUsers struct {
	Id        int       `gorm:"column:id;type:INT;" json:"id" primaryKey;AUTO_INCREMENT""`
	Name      string    `gorm:"column:name;type:VARCHAR(255);" json:"name" form:"name"`
	Age       int32     `gorm:"column:age;type:INT;" json:"age" form:"age"`
	Sex       int32     `gorm:"column:sex;type:INT;" json:"sex" form:"sex"`
	Birth     time.Time `gorm:"column:birth;type:DATETIME;" json:"birth" form:"birth"`
	Addr      string    `gorm:"column:addr;type:VARCHAR(255);" json:"addr" form:"addr"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

/*
type UniUsers struct {
	Name string `gorm:"column:name;type:VARCHAR(255);" json:"name"`
	Age int32 `gorm:"column:age;type:INT;" json:"age"`
	Sex int32 `gorm:"column:sex;type:INT;" json:"sex"`
	Birth time.Time `gorm:"column:birth;type:DATETIME;" json:"birth"`
	Addr string `gorm:"column:addr;type:VARCHAR(255);" json:"addr"`
}


*/

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

func (u *UniUsers) TableName() string {
	return "uni_users"
}
