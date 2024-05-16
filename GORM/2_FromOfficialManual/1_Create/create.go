package main

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"strconv"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "go:PassCode987!@tcp(api.ikanned.com:4407)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// parseTime=True&loc=Local 用于处理时间
	//Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 第一种写法
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

type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

// CreateRecord 创建记录
func CreateRecord() {
	users := []*User{ // NOTE 你无法向 ‘create’ 传递结构体，所以你应该传入数据的指针.
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}
	result := Db.Create(users)
	if result.Error != nil {
		log.Println("error: ", result.Error)
	}
	log.Println("影响的行数 =", result.RowsAffected)
}

// CreateMultiRecord 批量插入
func CreateMultiRecord() {
	// 要高效地插入大量记录，请将切片传递给Create方法。
	// GORM 将生成一条 SQL 来插入所有数据，以返回所有主键值，并触发 Hook 方法。
	// 当这些记录可以被分割成多个批次时，GORM会开启一个事务</0>来处理它们。
	var users []*User
	for i := 1; i <= 10; i++ {
		users = append(users, &User{"testUser" + strconv.Itoa(i), 18, time.Now()})
	}
	// 方法1
	// Db.Create(&users) // 将切片传递给Create方法

	// 方法2
	// 你可以通过db.CreateInBatches方法来指定批量插入的批次大小
	result := Db.CreateInBatches(users, len(users))
	if result.Error != nil {
		log.Println("error: ", result.Error)
	}
	log.Println("影响的行数 =", result.RowsAffected)
}

// GORM 中的钩子（Hooks）是 GORM 在执行某些操作时自动调用的一些方法。
// 这些方法可以在执行操作前或操作后执行一些自定义逻辑。
// 常用的钩子包括 `BeforeSave`、`BeforeCreate`、`AfterSave` 和 `AfterCreate` 等。

// BeforeCreate 创建钩子
// 根据规则 那么这个会在调用Create用户的时候被调用
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Name == "admin" {
		return errors.New("非法用户名")
	}
	return err
}

// AfterCreate 创建一个Create成功后会被执行的钩子
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	// 这个实现的接口将会在Create用户成功后被调用 如果失败则不会调用
	log.Println("创建用户成功 by AfterCreate")
	return
}

// TestCreateInvalidUser 模拟创建一个非法用户来调用上面实现的钩子
func TestCreateInvalidUser() {
	user := &User{
		Name:     "admin", // 这里的字段值admin是不合法的
		Age:      18,
		Birthday: time.Now(),
	}
	// 如果你想跳过Hooks方法，可以使用SkipHooks会话模式，例子如下
	Db.Session(&gorm.Session{SkipHooks: true}).Create(user)
	// 执行插入
	result := Db.Create(user)
	if result.Error != nil {
		log.Println("error: ", result.Error)
	}
	log.Println("影响的行数 =", result.RowsAffected)
}

// TestMap 根据 Map 创建
// GORM支持通过 map[string]interface{} 与 []map[string]interface{}{}来创建记录。
func TestMap() {
	// 注意
	// 当使用map来创建时，钩子方法不会执行，关联不会被保存且不会回写主键。
	result := Db.Model(&User{}).Create(map[string]interface{}{ // 这里测试使用map
		"Name": "jinzhu", "Age": 18,
	})
	if result.Error != nil {
		log.Println("error: ", result.Error)
	}
}

// TestSqlSyntaxInsert 使用 SQL 表达式、Context Valuer 创建记录
func TestSqlSyntaxInsert() {
	//  GORM允许使用SQL表达式来插入数据，有两种方法可以达成该目的，使用map[string]interface{} 或者 Customized Data Types
	// 从map来创建
	Db.Model(User{}).Create(map[string]interface{}{
		"Name":     "user1",
		"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	})
	// 等同于sql语句 INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));
}

// TestSliceMap 使用[]map[string]interface{}{}来创建记录
func TestSliceMap() {
	// batch insert from `[]map[string]interface{}{}`
	result := Db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})
	if result.Error != nil {
		log.Println("error: ", result.Error)
	}
}

func main() {
	log.Println("本地时间: ", time.Now())
	//CreateRecord()
	//CreateMultiRecord()
	//TestCreateInvalidUser()
	//TestMap()
	//TestSliceMap()
	TestSqlSyntaxInsert()

}
