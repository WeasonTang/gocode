package main

import (
	"errors"
	"fmt"
	"gocode/chartroom/common/message"
	"gocode/chartroom/common/utils"
	"gocode/chartroom/server/processes"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//ServerProcessMes 根据客户端发送消息不同,来决定调用哪个函数来处理
func (processor *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		up := processes.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := processes.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//处理用户发送的消息
		smsProcess := &processes.SmsProcess{}
		err = smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("invalid message type")
	}
	return
}

func (processor *Processor) deal() (err error) {
	//循环读取客户端信息
	for {
		tf := utils.Transfer{
			Conn: processor.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				err = errors.New("客户端退出，服务器重新等待客户端连接")
				// fmt.Println("客户端退出，服务器重新等待客户端连接..")
			}
			return err
		}
		// fmt.Println("mes=", mes)
		err = processor.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes(conn, mes) err=", err)
			return err
		}
	}

}
