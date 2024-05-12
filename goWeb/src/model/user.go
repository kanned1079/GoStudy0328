package model

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

func (u *User) AddUser(err error) {
	// 1.写sql语句
	sqlStr := "insert into users(username, password, email) values (?, ?, ?)"
	// 2.预编译
	inStmt, err := utils.D
}
