package model

import (
	"GoStudy0328/src/79_Chat/common/message"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
)

// 在服务器启动后就初始化一个UserDao实例
// 在需要对Redis操作时候直接调用
var (
	MyUserDao *UserDao // 现在只是一个空的
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
	res, err := redis.String(conn.Do("HGET", "users", id))
	log.Println("数据库读取的用户数据 =", res)
	if err != nil {
		if err == redis.ErrNil { // 这个说明在users中找不到对应的Id
			log.Println("在数据库中找不到对应Id的用户")
			err = ERROR_USER_NOT_EXIST // 用户不存在
		}
		return
	} // 到这里就是成功
	// 这里需要把res反序列化成User实例
	user = &User{}                           // 先创建要
	err = json.Unmarshal([]byte(res), &user) // pass
	if err != nil {
		log.Println("res反序列化失败 err = ", err)
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
			log.Println("recover err ", err)
		}
	}()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		log.Println("getUserById err ", err)
		return
	} // 这时证明这个用户是获取到了
	if user.UserPassword != userPwd {
		err = ERROR_USER_PWD
		log.Println("密码不正确")
		return
	}
	// pass
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	// 先从连接池中取出一个连接
	conn := this.pool.Get()
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover err ", err)
		}
	}()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil { // 这里查询到了反而是用户已经存在
		log.Println("用户已经存在 getUserById err ", err)
		err = ERROR_USER_EXISTS // 返回用户已经存在
		return
	} // 这时证明这个用户还没有注册过
	// 在Redis中查询不到 因此可以入库完成注册
	data, err := json.Marshal(user)
	if err != nil {
		log.Println("新建用户序列化错误 json err ", err)
		return
	}
	// 开始入库
	_, err = conn.Do("HSET", "users", user.UserId, string(data))
	if err != nil {
		log.Println("新用户入库错误 err ", err)
		return
	}
	// pass
	return
}
