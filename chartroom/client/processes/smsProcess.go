package processes

import (
	"encoding/json"
	"fmt"
	"gocode/chartroom/common/message"
	"gocode/chartroom/common/utils"
	"gocode/chartroom/server/model"
)

type SmsProcess struct {
}

//发送群聊消息
func (smsProcess SmsProcess) SendGroupMes(content string) (err error) {
	//1 创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2 创建一个SmsMes实例
	smsMes := message.SmsMes{
		Content: content,
		User: model.User{
			UserId:     curUser.UserId,
			UserStatus: curUser.UserStatus,
		},
	}
	//3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(smsMes) err=", err)
		return
	}
	mes.Data = string(data)
	//4.序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(mes) err=", err)
		return
	}

	//5.将序列化后的mes发送给服务器
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes  tf.WritePkg(data) err=", err)
	}
	return
}
