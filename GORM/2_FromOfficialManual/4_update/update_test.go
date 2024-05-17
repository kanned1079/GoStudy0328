package main

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

type User struct {
	Id       int
	Name     string
	Age      int
	Birthday time.Time
	Address  string
}

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "go:PassCode987!@tcp(api.ikanned.com:4407)/db1?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接错误 err:", err)
		return
	} else {
		log.Println("数据库连接成功")
	}
}

// UseSave Save 会保存所有的字段，即使字段是零值
func UseSave() {
	var user User
	DB.Where("id = ?", 1).First(&user)
	user.Age = 91  // 这里修改了结构体的值
	DB.Save(&user) // 这里调用Save进行保存

	// 保存 是一个组合函数。 如果保存值不包含主键，它将执行 Create，否则它将执行 Update (包含所有字段)。
	DB.Save(&User{Name: "于贤达2", Age: 66, Birthday: time.Now()}) // 这里是不包含PK的 所以它将会新建一条记录 相当于Create
	// 相当于 INSERT INTO `users` (`name`,`age`,`birthday`,`update_at`) VALUES ("jinzhu",100,"0000-00-00 00:00:00","0000-00-00 00:00:00")

	DB.Save(&User{Id: 1, Name: "于贤达", Age: 33, Birthday: time.Now()}) // 由于这里有指定PK 所以就相当于进行了update更新操作
	// 相当于 UPDATE `users` SET `name`="jinzhu",`age`=100,`birthday`="0000-00-00 00:00:00",`update_at`="0000-00-00 00:00:00" WHERE `id` = 1
}

// UpdateOneCow 更新单个列
// 当使用 Update 更新单列时，需要有一些条件，否则将会引起ErrMissingWhereClause 错误

func UpdateOneCow() {
	var user User
	// 当使用 Model 方法，并且它有主键值时，主键将会被用于构建条件，例如：

	// 根据条件进行更新
	DB.Model(&User{}).Where("name = ?", "于贤达2").Update("name", "yxd2222")

	// 根据ID进行更新
	user.Id = 2
	DB.Model(&user).Update("age", 99)

	// 根据条件和 model 的值进行更新
	var user2 User
	user2.Id = 12
	DB.Model(&user2).Update("birthday", time.Now())
}

// UpdateMultiCows 更新多列
// Updates 方法支持 struct 和 map[string]interface{} 参数。
// 当使用 struct 更新时，默认情况下GORM 只会更新非零值的字段
func UpdateMultiCows() {
	var user1 User
	user1.Id = 12
	// 根据 `struct` 更新属性，只会更新非零值的字段
	DB.Model(&user1).Updates(&User{Name: "陈丽2", Age: 66})

	// 根据 `map` 更新属性
	DB.Model(&user1).Updates(map[string]any{"Name": "陈丽3", "Age": 33})

}

// UpdateSpecFields 更新选定字段
func UpdateSpecFields() {
	var user1 User
	user1.Id = 14
	// 在更新时选择、忽略某些字段，您可以使用 Select、Omit
	// 使用Map对象
	DB.Model(&user1).Select("address").Updates(map[string]any{"Name": "ydddd", "Address": "ShangHai"})
	// 以上语句Select作用是 只更新用户的Address数据

	DB.Model(&user1).Omit("age").Updates(map[string]any{"Name": "ydd001", "Age": 16, "Birthday": time.Now(), "Address": "BeiJing"})
	// 以上语句Omit作用是 掠过age 更新其他的

	// 选择 Struct 的字段（会选中零值的字段）
	DB.Model(&user1).Select("name", "age").Updates(&User{Name: "ydd002", Age: 22})

	// 选择所有字段（选择包括零值字段的所有字段）
	// 下面这句不会跳过默认为0的字段
	// 由于后面没有写入Id 因为为0值 会将此用户的Id改为0
	DB.Model(&user1).Select("*").Updates(&User{Name: "ydd03", Age: 24, Address: "ChangzhouCity", Birthday: time.Now()})

	// 选择除 name 外的所有字段（包括零值字段的所有字段）
	//
	DB.Model(&user1).Where("id = ?", 0).Select("*").Omit("name").Updates(&User{Id: 14, Name: "admin", Age: 99, Birthday: time.Now(), Address: "Nanjing"})
}

// UpdateHook 更新 Hook
// GORM 支持的 hook 包括：BeforeSave, BeforeUpdate, AfterSave, AfterUpdate. 更新记录时将调用这些方法

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Name == "admin" {
		log.Println("非法用户名")
		return errors.New("非法用户名")
	} else {
		log.Println("正常更新流程开始")
	}
	return
}

// BatchUpdate 批量更新
func BatchUpdate() {
	// 如果没有指定具体的struct 那么GORM将会进行全局的更新
	// 使用Struct
	DB.Model(&User{}).Where("age = ?", 17).Updates(&User{Birthday: time.Now()})
	// 以上语句将更新所有年龄为17岁的用户的日期为当前日期

	// 使用map
	DB.Model(&User{}).Where("age = ?", 16).Updates(map[string]any{"Address": "China"})
	// 以上语句将所有年龄为16对的用户的地址改为China
}

func TestUpdate(t *testing.T) {
	//UseSave()
	//UpdateOneCow()
	//UpdateMultiCows()
	//UpdateSpecFields()
	BatchUpdate()
}
