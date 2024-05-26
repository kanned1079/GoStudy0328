package user

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"log"
)

// FullQuery 请求获取用户所有数据
// 返回值 请求通过 请求失败 用户不存在
func (u *MyUser) FullQuery() (code int) {
	isExist := IsUserExistById(u.UserId)
	if isExist == USER_EXISTED {
		result := dao.Db.First(&u)
		if result.RowsAffected == 1 {
			log.Println("FullQuery PASS")
			var encry Encryptor
			u.Password = encry.Decrypt(u.Password)
			return QUERY_PASS
		} else {
			log.Println("FullQuery FAIL")
			return QUERY_FAILURE
		}
	} else if isExist == USER_NOT_EXIST {
		log.Println("用户已经删除或不存在")
		return USER_NOT_EXIST
	}
	return UNKNOW_ERROR
}

// FullQueryByEmail 请求获取用户所有数据通过邮箱地址
// 返回值 请求通过 请求失败 用户不存在
func (u *MyUser) FullQueryByEmail() (code int) {
	isExist := IsUserExistByEmail(u.Email)
	if isExist == USER_EXISTED {
		result := dao.Db.Model(&u).Where("email = ?", u.Email).First(&u)
		if result.RowsAffected == 1 {
			log.Println("FullQuery PASS")
			var encry Encryptor
			u.Password = encry.Decrypt(u.Password)
			return QUERY_PASS
		} else {
			log.Println("FullQuery FAIL")
			return QUERY_FAILURE
		}
	} else if isExist == USER_NOT_EXIST {
		log.Println("用户已经删除或不存在")
		return USER_NOT_EXIST
	}
	return UNKNOW_ERROR
}

// NewCreate 首次创建新用户注册
// 返回值 用户创建成功 创建失败 用户已经存在而创建失败 未知错误
func (u *MyUser) NewCreate() (code int) {
	isExist := IsUserExistByEmail(u.Email)
	if isExist == USER_NOT_EXIST {
		log.Println("用户不存在 可以注册")
		var encry Encryptor
		u.Password = encry.Encrypt(u.Password) // 密码base64编码
		result := dao.Db.Model(&u).Create(&u)
		if result.RowsAffected == 1 {
			log.Println("写入数据库成功")
			return USER_CREATE_SUCCESS
		} else {
			log.Println("写入错误 err =", result.Error)
			return USER_CREATE_FAILURE
		}
	} else if isExist == USER_EXISTED {
		log.Println("用户已在数据库中存在 创建失败")
		code = USER_EXISTED
		return
	}
	return UNKNOW_ERROR
}

// NewUpdate 更新用户的信息
// 返回信息有 更新成功 更新失败 用户不存在 未知错误
func (u *MyUser) NewUpdate() (code int) {
	isExist := IsUserExistByEmail(u.Email)
	log.Println("NewUpdate.isExist:", isExist)
	if isExist == USER_EXISTED { // 用户存在
		var result = dao.Db.Model(&MyUser{}).Where("email = ?", u.Email).Updates(&u)
		log.Println("NewUpdate: ", result)
		if result.RowsAffected == 1 {
			log.Println("更新成功")
			return USER_UPDATE_SUCCESS
		} else { // 插入数据没有成功
			log.Println("更新失败 err =", result.Error)
			return USER_UPDATE_FAIULRE
		}
	} else if isExist == USER_NOT_EXIST { // 用户不存在 不能更新
		return USER_NOT_EXIST
	}
	return UNKNOW_ERROR // 默认返回未知错误
}

// NewDelete 新方法 用于删除 给定邮箱的用户
// 返回值 删除成功 删除失败 用户不存在 未知错误
func (u *MyUser) NewDelete() (code int) {
	isExist := IsUserExistByEmail(u.Email)
	if isExist == USER_EXISTED { // 如果用户已经存在
		result := dao.Db.Model(&MyUser{}).Where("email = ?", u.Email).Delete(&u)
		if result.RowsAffected == 1 {
			return USER_DELETE_SUCCESS
		} else {
			return USER_DELETE_FAILURE
		}
	} else if isExist == USER_NOT_EXIST {
		return USER_NOT_EXIST
	}
	return UNKNOW_ERROR
}
