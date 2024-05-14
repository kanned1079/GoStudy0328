package model

import (
	"GoStudy0328/GoWeb/6_Database_Operation/src/utils"
	"log"
)

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

// AddUser 方法1 带有Prepare
func (u *User) AddUser() (err error) {
	// 1.写sql语句
	sqlStr := "insert into users(username, password, email) values (?, ?, ?)"
	// 2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		log.Println("预编译出现异常 err: ", err)
		return
	}
	// 3.执行
	_, err = inStmt.Exec(u.UserName, u.Password, u.Email)
	if err != nil {
		log.Println("inStmt.Exec 错误 err: ", err)
		return
	}
	// 4.测试
	return nil
}

// AddUser2 方法2 不带有Prepare
func (u *User) AddUser2() (err error) {
	// 1.写sql语句
	sqlStr := "insert into users(username, password, email) values (?, ?, ?)"
	// 2.预编译

	// 3.执行
	_, err = utils.Db.Exec(sqlStr, "testuser2", "pass", "user@xxx.com")
	if err != nil {
		log.Println("inStmt.Exec 错误 err: ", err)
		return
	}
	// 4.测试
	return nil
}

// GetUserById 根据用户id从数据库中查询一条记录
// QueryRow 用于从数据库中取出一条数据
func GetUserById(fromid int) (user *User, err error) {
	sqlStr := "select id, username, password, email from users where id = ?"
	row := utils.Db.QueryRow(sqlStr, fromid)
	var id int
	var username string
	var password string
	var email string
	err = row.Scan(&id, &username, &password, &email)
	if err != nil {
		log.Println("row.Scan error: ", err)
		return
	}
	user = &User{
		Id:       id,
		UserName: username,
		Password: password,
		Email:    email,
	}
	return user, nil
}

// GetUsersById 查询所有
func GetUsersById() (usersSlice []*User, err error) {
	sqlStr := "select id, username, password, email from users"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		log.Println("row.Scan error: ", err)
		return
	}
	usersSlice = make([]*User, 0)
	for rows.Next() {
		var id int
		var username string
		var password string
		var email string
		err = rows.Scan(&id, &username, &password, &email)
		if err != nil {
			log.Println("row.Scan error: ", err)
			return
		}
		user := &User{
			Id:       id,
			UserName: username,
			Password: password,
			Email:    email,
		}
		usersSlice = append(usersSlice, user)
	}

	return usersSlice, nil
}
