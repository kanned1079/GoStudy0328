package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 定义一个UserDao结构体
// 完成对User结构体的各种操作
type UserDao struct {
	// 需要连接池
	pool *redis.Pool
}

// 使用工厂模式 创建一个新的UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// getUserById 根据用户id 返回一个User实例和err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 通过给定的id去redis里查询这个用户
	res, err := redis.String(conn.Do("HGET", id, "user"))
	if err != nil {
		if err == redis.ErrNil { // 这个说明在users中找不到对应的Id
			err = ERROR_USER_NOT_EXIST
		}
		return
	} // 到这里就是成功
	// 这里需要把res反序列化成User实例
	user = &User{}                           // 先创建要
	err = json.Unmarshal([]byte(res), &user) // pass
	if err != nil {
		fmt.Println("res反序列化失败 err = ", err)
		return
	} // 到这里只是拿到了用户的数据 但是密码是否正确不知道
	return
}

// Login 完成登录校验 这里需要首字母大写
// 1)如果用户Id和密码都争取那就返回一个user实例
// 2)如果用户的Id或密码有误 返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 先从连接池中取出一个连接
	conn := this.pool.Get()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover err ", err)
		}
	}()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		fmt.Println("getUserById err ", err)
		return
	} // 这时证明这个用户是获取到了
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
