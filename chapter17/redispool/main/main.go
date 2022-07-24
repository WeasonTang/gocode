package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

//程序启动时，初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialPassword("54321"), redis.DialDatabase(6))
		},
	}
}

func main() {
	//先从pool 取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "hero", "taka")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//取出
	res, err := redis.String(conn.Do("get", "hero"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println("hero is", res)

}
