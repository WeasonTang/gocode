package message

import "gocode/chartroom/server/model"

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//定义用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息类型
}

//定义两个消息..
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code    int    `json:"code"`    //返回状态码 500-用户未注册 200-登录成功
	Error   string `json:"error"`   //返回错误信息
	UserIds []int  `json:"userIds"` //增加保存用户id的切片
}

type RegisterMes struct {
	User model.User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//配合服务器端推送用户状态
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//发送消息
type SmsMes struct {
	Content string     `json:"content"` //消息内容
	User    model.User `json:"user"`
}
