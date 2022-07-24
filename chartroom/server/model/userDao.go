package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体
//完成对 User结构体 的各种操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式 创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{pool}
	return
}

//根据id获取用户信息
func (userDao *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定id 去redis查询用户
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		// fmt.Println("err=", err)
		if err == redis.ErrNil {
			err = Err_USER_NOTEXISTS
		}
		return
	}
	// fmt.Println("res=", res)
	user = &User{}
	//把res反序列化成user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}
	return
}

//登录校验
func (userDao *UserDao) CheckUserInformation(login *User) (user *User, err error) {
	//先从userDao 的连接池取出一个连接
	conn := userDao.pool.Get()
	defer conn.Close()
	user, err = userDao.getUserById(conn, login.UserId)
	if err != nil {
		return
	}
	//证明用户存在 再校验密码
	if user.UserPwd != login.UserPwd {
		err = Err_USER_PWD
	}
	return
}

//登录校验
func (userDao *UserDao) Register(user *User) (err error) {
	//先从userDao 的连接池取出一个连接
	conn := userDao.pool.Get()
	defer conn.Close()
	//校验该用户是否已注册
	dataBaseUser, err := userDao.getUserById(conn, user.UserId)
	if err != Err_USER_NOTEXISTS {
		if dataBaseUser != nil {
			err = Err_USER_EXISTS
		}
		return
	}
	//证明该用户未注册 再完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户信息错误 err=", err)
	}
	return
}
