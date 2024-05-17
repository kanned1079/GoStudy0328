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
	Id        int
	Name      string
	Age       int
	Birthday  time.Time
	Address   string
	DeletedAt gorm.DeletedAt // 软删除
	// 当调用Delete时，GORM并不会从数据库中删除该记录，而是将该记录的DeleteAt设置为当前时间，而后的一般查询方法将无法查找到此条记录。
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

// 删除一条记录
func testDeleteOneRecord(t *testing.T) {
	// 删除一条记录时，删除对象需要指定主键，否则会触发**批量删除**
	var user1 User
	user1.Id = 14
	//DB.Delete(&user1)

	// 带有额外条件的删除
	user1.Address = "TW"
	DB.Where("address = ?", "TW").Delete(&user1)
	// 相当于 DELETE from users where id = 14 AND address = "TW";
}

func testDeleteRecordByPK(t *testing.T) {
	DB.Delete(&User{}, 14)
	//
	//DB.Delete(&User{}, "14") // 相同

	//DB.Delete(&User{}, []int{15, 16, 17})
}

// 对于删除操作，GORM 支持 BeforeDelete、AfterDelete Hook，在删除记录时会调用这些方法

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	log.Println(u)
	if u.Name == "admin" {
		log.Println("管理员不能被删除")
		return errors.New("admin cannot be deleted.")
	} else {
		log.Println("该用户允许删除")
	}
	return nil
}

// 批量删除
func testBatchDelete(t *testing.T) {
	// 可以将一个主键切片传递给Delete 方法，以便更高效的删除数据量大的记录
	var users = []User{{Id: 14}, {Id: 15}, {Id: 16}, {Id: 17}}
	result := DB.Delete(&users)
	log.Println("RowsAffected =", result.RowsAffected)
}

// 软删除
func testSoftDelete(t *testing.T) {
	var user1 User
	user1.Id = 13
	result := DB.Delete(&[]User{{Id: 14}, {Id: 15}, {Id: 16}, {Id: 17}})
	log.Println("RowsAffected =", result.RowsAffected)
}

func testSelectAll(t *testing.T) {
	var users []User
	result := DB.Find(&users)
	log.Println("RowsAffected =", result.RowsAffected)
	for _, users := range users {
		log.Println(users)
	}
}

func TestDelete(t *testing.T) {
	//t.Run("测试删除一条记录", testDeleteOneRecord)
	//t.Run("根据主键删除", testDeleteRecordByPK)
	//t.Run("批量删除", testBatchDelete)
	t.Run("软删除", testSoftDelete)

	t.Run("查询所有用户", testSelectAll)
}
