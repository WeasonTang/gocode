package processes

import (
	"encoding/json"
	"fmt"
	"gocode/chartroom/common/message"
	"gocode/chartroom/common/utils"
)

type SmsProcess struct{}

func (smsProcess *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	smsMes, err := smsProcess.dealSendMes(mes)
	if err != nil {
		fmt.Println("smsProcess.dealSendMes(mes) err=", err)
		return
	}
	//将smsMes序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		return
	}
	mes.Data = string(data)
	if err != nil {
		return
	}
	//遍历服务器端onlineUsers
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.User.UserId {
			continue
		}
		//创建Transfer实例 发送mes
		tf := &utils.Transfer{
			Conn: up.Conn,
		}
		err := tf.WriteMesPkg(mes)
		if err != nil {
			fmt.Printf("给%d转发消息出错 err=", err)
		}
	}
	return
}

func (smsProcess *SmsProcess) dealSendMes(mes *message.Message) (smsMes message.SmsMes, err error) {
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) err=", err)
		return
	}
	//对content进行处理的代码
	content := smsMes.Content
	fmt.Println("content=", content)
	smsMes.Content = content
	return
}
