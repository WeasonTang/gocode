package main

import (
	"fmt"
	"gocode/chartroom/client/processes"
	"os"
)

//用户id
var userId int

//用户密码
var userPwd string

//用户密码
var userName string

func main() {
	//接收用户选择
	var key int
	//判断是否还继续显示菜单
	// var loop = true

	for {
		fmt.Println("------------欢迎登录多人聊天系统------------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录
			up := processes.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登录失败 err=", err)
			}
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名")
			fmt.Scanf("%s\n", &userName)
			up := processes.UserProcess{}
			err := up.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("注册失败 err=", err)
			}
		case 3:
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
