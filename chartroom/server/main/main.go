package main

import (
	"fmt"
	"gocode/chartroom/server/model"
	"net"
	"time"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	processor := Processor{
		Conn: conn,
	}
	err := processor.deal()
	if err != nil {
		fmt.Println(err.Error())
	}
}

//初始化UserDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	//服务器启动时 初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	//初始化UserDao
	initUserDao()
}

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen er=", err)
		return
	}
	defer listen.Close()
	//监听成功 等待客户端连接
	fmt.Println("等待客户端来连接服务器")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//连接成功 启动协程和客户端保持通讯
		go process(conn)
	}
}
