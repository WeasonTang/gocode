package processes

import (
	"encoding/json"
	"fmt"
	"gocode/chartroom/common/message"
	"gocode/chartroom/common/utils"
	"gocode/chartroom/server/model"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//表示Conn是哪个用户的
	UserId int
}

//通知所有在线用户
func (userProcess *UserProcess) NotifyOthersOnlineUser(userId int) {

	//组装NotifyUserStatus
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) err=", err)
		return
	}
	mes.Data = string(data)
	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}

	//遍历 onlineUsers, 然后一个一个的发送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤掉自己
		if id == userId {
			continue
		}
		//通知up 用户userId 上线
		up.NotifyMeHasUserOnline(userId, data)
	}
}

func (userProcess *UserProcess) NotifyMeHasUserOnline(userId int, data []byte) {

	//创建Transfer实例发送
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeHasUserOnline err=", err)
		return
	}
}

func (userProcess *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从 mes 中取出 mes.Data,并反序列化
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal mes.Data fail err=", err)
		return
	}

	//准备处理后返回消息数据实例
	var resMes message.Message
	resMes.Type = message.RegisterMesType
	var registerResMes message.RegisterResMes

	//2.注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.Err_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.Err_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = err.Error()
		}
	} else {
		registerResMes.Code = 200
	}

	//3.将 registerResMes
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) fail err=", err)
		return
	}

	//4.把登陆信息赋给resMes 并序列化
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail err=", err)
		return
	}

	//5.发送消息给客户端
	// fmt.Println("发送消息给客户端 data=", string(data))
	tf := utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//处理登录
func (userProcess *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从 mes 中取出 mes.Data,并反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal mes.Data fail err=", err)
		return
	}
	//准备处理后返回消息数据实例
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	//2.登录
	login := model.User{
		UserId:  loginMes.UserId,
		UserPwd: loginMes.UserPwd,
	}
	user, err := model.MyUserDao.CheckUserInformation(&login)
	if err != nil {
		loginResMes.Code = 500
		loginResMes.Error = err.Error()
		fmt.Println(loginMes, "登录失败 err=", err)
	} else {
		loginResMes.Code = 200
		userProcess.UserId = user.UserId
		userMgr.AddOnlineUser(userProcess)
		//通知其它用户 用户user上线
		userProcess.NotifyOthersOnlineUser(user.UserId)

		//获取当前在线用户id列表
		for id := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Printf("登录成功..userId=%d userName=%s\n", user.UserId, user.UserName)
	}

	//3.将 loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) fail err=", err)
		return
	}

	//4.把登陆信息赋给resMes 并序列化
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail err=", err)
		return
	}

	//5.发送消息给客户端
	// fmt.Println("发送消息给客户端 data=", string(data))
	tf := utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}
