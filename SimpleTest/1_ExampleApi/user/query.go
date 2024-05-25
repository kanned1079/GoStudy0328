package user

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"log"
)

// IsUserExistById 通过ID判断用户是否存在
func IsUserExistById(userId int) (code int) {
	var user MyUser
	user.UserId = userId
	result := dao.Db.Model(&MyUser{}).Where("user_id = ?", user.UserId).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		log.Println("IsUserExistById: 用户不存在")
		code = USER_NOT_EXIST
		return
	} else {
		log.Println("IsUserExistById: 用户存在")
		code = USER_EXISTED
		return
	}
}

// IsUserExistByEmail 通过Email判断用户是否存在
func IsUserExistByEmail(mail string) (code int) {
	var user MyUser
	user.Email = mail
	result := dao.Db.Where("email = ?", user.Email).First(&user)
	if result.RowsAffected != 0 {
		log.Println("IsUserExistByEmail: 用户存在")
		code = USER_EXISTED
		return
	} else {
		log.Println("IsUserExistByEmail: 用户不存在")
		code = USER_NOT_EXIST
		return
	}
}

// AuthLoginQuery 验证用户名和密码
func AuthLoginQuery(mail, pwd string) (code int) {
	var user MyUser
	user.Email = mail
	user.Password = pwd
	if IsUserExistByEmail(user.Email) == USER_EXISTED {
		// 验证密码
		var newUser MyUser
		var encry Encryptor
		dao.Db.Model(&MyUser{}).Where("email = ?", user.Email).Limit(1).Find(&newUser)
		newUser.Password = encry.Decrypt(newUser.Password)
		if newUser.Password == user.Password {
			return AUTH_PASS
		} else {
			return PASSWORD_INCORRECT
		}
	} else {
		return USER_NOT_EXIST
	}
}

// HaveBeenDeleted 查询用户是否被删除
func HaveBeenDeleted(userId int, email string) (code int) {
	var user MyUser
	user.UserId = userId
	user.Email = email
	// 查找被软删除的记录 可以使用Unscoped来查询到被软删除的记录
	result := dao.Db.Unscoped().Where("user_id = ? or email = ?", userId, email).Find(&user)
	if result.RowsAffected == 0 {
		log.Println("不存在已经被删除的用户")
		return USER_NOT_EXIST
	} else {
		log.Println("存在已经被删除的用户: ", user.Email)
		return USER_EXISTED
	}
}

func RecordsCount() (count int64) {
	dao.Db.Model(&MyUser{}).Count(&count)
	return
}

func ValidRecordsCount() (count int64) {
	dao.Db.Model(&MyUser{}).Where("deleted_at != null").Count(&count)
	return
}

func InValidRecordsCount() (count int64) {
	dao.Db.Model(&MyUser{}).Where("deleted_at IS NOT NULL").Count(&count)
	return
}
