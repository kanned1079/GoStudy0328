package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	Id       int
	Name     string
	Age      int
	Birthday time.Time
}

type User2 struct {
	ID        string         `gorm:"primarykey;size:16"`
	Name      string         `gorm:"size:24"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var (
	Db  *gorm.DB
	err error
)

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "go:PassCode987!@tcp(api.ikanned.com:4407)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// parseTime=True&loc=Local 用于处理时间
	//dao, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 第一种写法
	// MySQL 驱动程序提供了 一些高级配置 可以在初始化过程中使用，例如：
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{})
	if err != nil {
		log.Println("Database connection failed. err: ", err)
	} else {
		log.Println("Database connection success.")
	}
}

// 检索单个对象
func SelectOneRecord() {
	user := User{}
	// 获取第一条记录（主键升序）
	result := Db.First(&user)
	// 相当于 SELECT * FROM users ORDER BY id LIMIT 1;
	log.Println(user)

	// 获取一条记录，没有指定排序字段
	Db.Take(&user)
	// 相当于 SELECT * FROM users LIMIT 1;
	log.Println(user)

	// 获取最后一条记录（主键降序）
	Db.Last(&user)
	// 相当于 SELECT * FROM users ORDER BY id DESC LIMIT 1;
	log.Println(user)

	if result.Error != nil {
		log.Println(result.Error)
	}
	//log.Println(user)

	// 如果你想避免ErrRecordNotFound错误，
	// 你可以使用Find，比如db.Limit(1).Find(&user)，Find方法可以接受struct和slice的数据。
	result = Db.Limit(1).Find(&user)
	log.Println(result.Error)
	log.Println(user)

	// First and Last 方法会按主键排序找到第一条记录和最后一条记录 (分别)。
	// 只有在目标 struct 是指针或者通过 db.Model() 指定 model 时，该方法才有效。
	// 此外，如果相关 model 没有定义主键，那么将按 model 的第一个字段进行排序。
	var user1 User
	var users1 []User
	Db.Find(&users1) // 正确 因为目标结构体被传入
	log.Println(user1)

	res := map[string]any{}
	Db.Model(&User{}).Find(&res) // 正确 因为目标结构体被传入

	// 错误
	//dao.Table("users").First(&result)
	//log.Println(users1)
	// 但是可以使用Take， 如下
	Db.Model("users").Take(&user1) // 取出所有记录
	log.Println(users1)
}

// TestSelectByID 根据主键检索
func TestSelectByID() {
	// 如果主键是数字类型，您可以使用 内联条件 来检索对象。
	user1 := User{}
	user2 := User{}
	users := []User{}
	result := Db.First(&user1, 1)
	log.Println("err: ", result.Error, "| num: ", result.RowsAffected)
	log.Println(user1, "\n-------------------")

	// 当使用字符串时，需要额外的注意来避免SQL注入
	result = Db.First(&user2, "2")
	log.Println("err: ", result.Error, "| num: ", result.RowsAffected)
	log.Println(user2, "\n-------------------")

	// 注意下面用的是Find 而不是First
	result = Db.Find(&users, []int{1, 2, 3, 4, 5, 6})
	log.Println("err: ", result.Error, "| num: ", result.RowsAffected)
	for _, v := range users {
		log.Println(v)
	}

	// 如果用户ID为uuid，要这么写
	user4 := User{}
	Db.First(&user4, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	fmt.Println("----------------")

	// 当目标对象有一个主键值时，将使用主键构建查询条件
	user5 := User{
		Id: 10,
	}
	Db.Find(&user5)
	log.Println(user5)

	fmt.Println("-----------------")

	// 也可以这么写
	var res2 User
	Db.Model(User{Id: 3}).First(&res2)
	log.Println(res2)

	// DeletedAt
}

// SelectAll 检索所有对象
func SelectAll() {
	users := []User{}
	res := Db.Find(&users)
	log.Println(res)
	for _, v := range users {
		log.Println(v)
	}
}

// StringMatch String 条件
func StringMatch() {
	user1 := User{}

	// 取到一条记录
	Db.Where("name = ?", "涂佳园").First(&user1)
	log.Println(user1, "\n-------------------")

	// 取到匹配的所有记录
	// 查询所有age=17的记录
	user_list1 := []User{}
	Db.Where("age = ?", 16).Find(&user_list1)
	rangeList(user_list1)
	fmt.Println("-----------------")

	// 查询记录是否在提供的切片中的
	user_list2 := []User{}
	Db.Where("name IN ?", []string{"涂佳园", "李雨桐"}).Find(&user_list2)
	rangeList(user_list2)
	fmt.Println("-----------------")

	// 同时满足条件的
	user_list3 := []User{}
	Db.Where("name = ? AND age = ?", "孙伟博", 14).Find(&user_list3)
	rangeList(user_list3)
	fmt.Println("-----------------")

	// 模糊查找姓名
	user_list4 := []User{}
	Db.Where("name LIKE ?", "%博%").Find(&user_list4)
	rangeList(user_list4)
	fmt.Println("-----------------")

	user_list5 := []User{}
	Db.Where("updated_at > ?", "2019-02-08 18:08:27").Find(&user_list5)
	rangeList(user_list5)
	fmt.Println("-----------------")

	// 如果对象设置了主键，条件查询将不会覆盖主键的值，而是用 And 连接条件。 例如：
	var user7 = User{Id: 10}
	Db.Where("id = ?", 5).First(&user7)
	// SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1
	// 这个查询将会给出record not found错误
	// 所以，在你想要使用例如 user 这样的变量从数据库中获取新值前，需要将例如 id 这样的主键设置为nil。

}

func rangeList(user []User) {
	for _, i := range user {
		log.Println(i)
	}
}

// SelectViaStructOrMap Struct & Map 条件
func SelectViaStructOrMap() {
	// Struct
	list1 := []User{}
	Db.Where(&User{Name: "涂佳园", Age: 17}).Find(&list1)
	rangeList(list1)
	fmt.Println("-----------------")

	// Map
	list2 := []User{}
	Db.Where(map[string]any{"Name": "丁浩晨", "Age": 14}).Find(&list2)
	rangeList(list2)
	fmt.Println("-----------------")

	// Slice
	list3 := []User{}
	Db.Where([]int64{6, 11}).Find(&list3) // 这里将会默认使用PK Id
	rangeList(list3)
	fmt.Println("-----------------")

}

// SelectAndSort 排序
func SelectAndSort() {
	list1 := []User{}
	Db.Order("age desc").Find(&list1)
	rangeList(list1)
	fmt.Println("-----------------")
}

func main() {
	//SelectOneRecord()
	//TestSelectByID()
	//SelectAll()
	//StringMatch()
	//SelectViaStructOrMap()
	SelectAndSort()
}
