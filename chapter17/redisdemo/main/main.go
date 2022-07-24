package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//1.通过密码连接redis对应库
	pwd := redis.DialPassword("54321")
	db := redis.DialDatabase(6)
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", pwd, db)
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	fmt.Println("connect success...")

	//2.String操作 set写入
	_, err = conn.Do("set", "name", "weason")
	if err != nil {
		fmt.Println("set error :", err)
		return
	}
	fmt.Println("set操作成功")

	//获取字符串
	res, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println("get操作成功 name =", res)

	//3.hash操作
	conn.Do("hmset", "emp", "name", "weason", "age", 28)
	if err != nil {
		fmt.Println("hmset err=", err)
		return
	}
	result, err := redis.Strings(conn.Do("hmget", "emp", "name", "age"))
	if err != nil {
		fmt.Println("hmget err=", err)
		return
	}
	fmt.Printf("hmget操作成功 emp =%v\n", result)

	//4.设置过期时间
	_, err = conn.Do("expire", "name", 10) //10秒过期
	if err != nil {
		fmt.Println("set expire error: ", err)
		return
	}
}
