package user

import (
	"GoStudy0328/SimpleTest/1_ExampleApi/dao"
	"errors"
	"gorm.io/gorm"
	"log"
)

// 钩子函数

func (u *MyUser) AfterDelete(tx *gorm.DB) (err error) {
	log.Println("[hook]AfterDelete")
	result := dao.Db.Limit(1).Find(&u)
	log.Println("AfterDelete: ", result.RowsAffected)
	if result.RowsAffected != 0 {
		return nil
	}
	return errors.New("删除未成功")
}

func (u *MyUser) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("[hook]BeforeCreate")
	if u.Name == "admin" {
		log.Println("非法用户名")
		return errors.New("非法用户名")
	}
	return nil
}

func (u *MyUser) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("[hook]AfterCreate")
	return nil
}

func (u *MyUser) AfterUpdate(tx *gorm.DB) (err error) {
	log.Println("[hook]AfterUpdate")
	log.Println("AfterUpdate: ", u.Name)
	return nil
}
