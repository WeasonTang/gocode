package processes

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"gocode/chartroom/common/message"
	"gocode/chartroom/common/utils"
	"gocode/chartroom/server/model"
	"net"
)

type UserProcess struct {
}

func (userProcess *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	//延迟关闭
	if err != nil {
		fmt.Println("net.Dail err=", err)
		return
	}
	defer conn.Close()

	//2.1准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	//2.2创建RegisterMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	//2.registerMes
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes) err=", err)
		return
	}
	//2.4把data赋给
	mes.Data = string(data)
	//2.5将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err=", err)
		return
	}

	//3.发送数据data给服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送注册信息报错 err=", err)
		return
	}

	//4.处理服务器端消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) err=", err)
		return
	}
	//解析mes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &registerResMes) err=", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("register success")
		// go serverProcessMes(conn)
		// //显示登录成功后的界面
		// showMenu()
	} else {
		err = errors.New(registerResMes.Error)
	}
	return

}

//登录函数
func (userProcess *UserProcess) Login(userId int, userPwd string) (err error) {
	//下一个就要开始定协议
	// fmt.Printf("userId = %d userPwd = %s\n", userId, userPwd)

	// return nil

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	//延迟registerMes关闭
	if err != nil {
		fmt.Println("net.Dail err=", err)
		return
	}
	defer conn.Close()

	//2.1准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//2.2创建LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//2.3将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err=", err)
		return
	}
	//2.4把data赋给
	mes.Data = string(data)
	//2.5将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err=", err)
		return
	}

	//3.发送数据data给服务器
	//3.1 先发送data的长度
	//先获取到 data的长度->转成一个表示长度的byte切片
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	// fmt.Printf("客户端，发送消息长度=%d\n 内容是%s\n", len(data), string(data))

	//3.2发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//4.处理服务器端消息
	tf := utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) err=", err)
		return
	}
	//解析mes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginResMes) err=", err)
		return
	}
	if loginResMes.Code == 200 {
		// fmt.Println("login success")
		//初始化curUser
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline

		//显示当前在线用户列表
		fmt.Println("当前用户列表如下：")
		for _, v := range loginResMes.UserIds {
			fmt.Println("用户id:\t", v)
			//完成 客户端 onlineUsers初始化
			user := &model.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		go serverProcessMes(conn)
		//显示登录成功后的界面
		showMenu()
	} else {
		err = errors.New(loginResMes.Error)
	}
	return
}
