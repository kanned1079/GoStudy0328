package users

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"errors"
	"log"
)

// IsUserExist 查找用户是否存在
// 提取了公共逻辑到 IsUserExist 函数中，IsUserExistById 和 IsUserExistByEmail 只需调用这个函数。
func IsUserExist(query, args any) (Code int, err error) {
	var user WxUsers
	result := dao.Db.Model(&WxUsers{}).Where(query, args).First(&user)
	log.Println("RowAffected:", result.RowsAffected)
	switch result.RowsAffected {
	case 0:
		return User_Not_Exist, errors.New("指定Id/Email的用户不存在")
	case 1:
		return User_Exist, errors.New("指定Id/Email的用户存在")
	default:
		return Unknow_Error, errors.New("未知的错误")
	}
}

// IsUserExistById_New 用于根据Id查询用户是否在数据库中的存在
// 返回值有 Unknow_Error User_Not_Exist User_Exist
func IsUserExistById_New(id int) (Code int, err error) {
	return IsUserExist("id = ?", id)
}

// IsUserExistByEmail_New 用于根据Email查询用户是否在数据库中的存在
// 返回值有 Unknow_Error User_Not_Exist User_Exist
func IsUserExistByEmail_New(email string) (Code int, err error) {
	return IsUserExist("email = ?", email)

}

// --------------改进前的-----------------------------------------------

func IsUserExistById(id int) (Code int, err error) {
	var user WxUsers
	result := dao.Db.Model(&WxUsers{}).First(&user)
	if result.Error != nil {
		log.Println("IsUserExistById err:", result.Error)
		Code = Unknow_Error
		err = result.Error
		return
	}
	switch result.RowsAffected {
	case 0:
		return User_Not_Exist, errors.New("指定Id的用户不存在")
	case 1:
		return User_Exist, nil
	default:
		return Unknow_Error, errors.New("未知的错误")
	}
}

func IsUserExistByEmail(email string) (Code int, err error) {
	var user WxUsers
	result := dao.Db.Model(&WxUsers{}).Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Println("IsUserExistByEmail err:", result.Error)
		Code = Unknow_Error
		err = result.Error
	}
	switch result.RowsAffected {
	case 0:
		return User_Not_Exist, errors.New("用户不存在")
	case 1:
		return User_Exist, nil
	default:
		return Unknow_Error, errors.New("未知的错误")

	}
}

func AuthUserCred(email, password string) (Code int) {
	//var user WxUsers
	return -1
}
