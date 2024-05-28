package users

import (
	"GoStudy0328/SimpleJSONTest/dao"
	"errors"
)

// CreateNew 创建用户
func (u *WxUsers) CreateNew() (code int, err error) {
	result := dao.Db.Model(WxUsers{}).Create(&u)
	if result.Error != nil {
		return Unknow_Error, errors.New("未知错误")
	}
	if result.RowsAffected == 0 {
		return Insert_Failure, errors.New("创建新用户失败")
	} else if result.RowsAffected >= 1 {
		return Insert_Success, nil
	}
	return Unknow_Error, errors.New("未知错误")
}

// FullQueryById 通过Id获取用户所有信息
func (u *WxUsers) FullQueryById() (Code int, err error) {

	return -1, nil
}

// FullQueryByEmail
func (u *WxUsers) FullQueryByEmail(email string) (Code int, err error) {
	return -1, nil

}
