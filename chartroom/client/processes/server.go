package processes

import (
	"encoding/json"
	"fmt"
	"gocode/chartroom/common/message"
	"gocode/chartroom/common/utils"
	"net"
	"os"
)

//接收用户选择
var key int

//发送消息内容
var content string

//循环显示登录成功后的界面
func showMenu() {
	for {
		fmt.Println("-----------恭喜登录成功-----------")
		fmt.Println("\t 1 显示在线用户列表")
		fmt.Println("\t 2 发送消息")
		fmt.Println("\t 3 信息列表")
		fmt.Println("\t 4 退出系统")
		fmt.Println("\t 请选择(1-4):")

		//创建smsProcess实例
		//因为发送消息操作有多次 所有在switch外定义
		smsProcess := &SmsProcess{}

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			// fmt.Println("显示在线用户列表")
			outPutOnlineUser()
		case 2:
			// fmt.Println("发送消息")
			fmt.Println("请输入消息内容")
			fmt.Scanf("%s\n", &content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("查看信息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例， 不停的读取服务器发送的消息
	tf := utils.Transfer{
		Conn: conn,
	}
	for {
		// fmt.Println("客户端正在等待读取消息中")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err=", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType: //其它用户状态有变化
			//1.取出 NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2.把这个用户信息，状态保存到客户端onlineUsers
			updateUserStatus(&notifyUserStatusMes)
			outPutOnlineUser()
		case message.SmsMesType:
			outPutGroupMes(&mes)
		default:
			fmt.Println("服务器返回未知消息类型")
		}
	}
}
