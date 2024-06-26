package users

import (
	"GoStudy0328/SimpleJSONTest/dao"
	"GoStudy0328/SimpleJSONTest/encrypt"
	"errors"
	"log"
)

// CreateNew 创建用户
func (u *WxUsers) CreateNew() (code int, err error) {
	u.Password = encrypt.Base64Encryptor.Base64Encrypt(u.Password)
	log.Println(u)
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
	result := dao.Db.Model(&u).Where("user_id = ?", u.UserId).First(&u)
	if result.RowsAffected == 0 {
		return User_Not_Exist, errors.New("用户不存在")
	} else if result.RowsAffected >= 1 {
		return Query_Pass, nil
	} else {
		return Unknow_Error, errors.New("出现未知错误")
	}
}

// FullQueryByEmail 通过传递的用户邮箱来获取用户所有的个人信息
// 用户时候存在需要调用query中封装的函数实现
// 返回值有 Query_Pass Query_Failure
func (u *WxUsers) FullQueryByEmail() (Code int, err error) {
	log.Println("[FullQueryByEmail] u.Email:", u.Email)
	result := dao.Db.Model(u).Where("email = ?", u.Email).First(u)
	log.Println("[FullQueryByEmail] error:", result.Error)
	if result.RowsAffected == 1 {
		return Query_Pass, result.Error
	} else {
		return Query_Failure, result.Error
	}
}

func (u *WxUsers) UpdateUserInfo() (Code int, err error) {
	result := dao.Db.Model(u).Where("user_id = ? or email = ?", u.UserId, u.Email).Updates(u)
	log.Println("[UpdateUserInfo] error:", result.Error)
	if result.RowsAffected == 1 {
		return Update_Success, result.Error
	} else {
		return Update_Failure, result.Error
	}
}

// DeleteUserByID 用于删除对应ID的用户
// 需要额外判断用户是否存在
func (u *WxUsers) DeleteUserByID() (Code int, err error) {
	result := dao.Db.Model(u).Where("id = ?", u.UserId).Delete(&u)
	log.Println("[DeleteUser] error:", result.Error)
	if result.RowsAffected == 1 {
		return Delete_Success, result.Error
	} else {
		return Delete_Failure, result.Error
	}
}

// Test new
func (u *WxUsers) CreateNewRec(code int) {
	var newUser WxUsers
	result := dao.Db.Model(&WxUsers{}).Limit(1).Find(&newUser)
	if result.RowsAffected >= 1 {
		code = Insert_Success
	} else if result.Error != nil {
		log.Println(result.Error)
		code = Insert_Failure
	}
	return
}
